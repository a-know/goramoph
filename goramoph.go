package main

import (
	"./model"
	"./parser"
	"./util"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	r "reflect"
	"regexp"
	"strings"
)

func main() {
	var fp *os.File
	var err error

	// パース対象の xml ファイル名を引数に受ける
	fp, err = os.Open(os.Args[1])
	util.FailOnError(err)
	defer fp.Close()

	var playDataList []model.Playdata
	playDataList = parser.ItunesXmlParse(fp)

	// 出力する csv ファイル名として用いるために、xml ファイルの最終更新日時を取得
	mod_date := util.GetModDate(fp)

	export_csv(mod_date, playDataList)

	//外部コマンドを実行し、プロジェクト名を取得する
	cmd := exec.Command("gcloud", "config", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	util.FailOnError(cmd.Run())
	fmt.Printf("in all caps: %q\n", out.String())
	//外部コマンドの実行結果からプロジェクト名を抽出する
	re, _ := regexp.Compile("\nproject = (.+)\n")
	one := re.Find([]byte(out.String()))
	fmt.Println("Find:", string(one))
	replace_re, _ := regexp.Compile("project|\\s|\n|=")
	project_name := replace_re.ReplaceAllString(string(one), "")
	fmt.Println("project_name:", string(project_name))
	//バケットが既に作成済みかどうかを調べて、未作成なら作成する
	cmd = exec.Command("gsutil", "ls")
	var ls_out bytes.Buffer
	cmd.Stdout = &ls_out
	util.FailOnError(cmd.Run())
	fmt.Printf("in all caps: %q\n", ls_out.String())
	if m, _ := regexp.MatchString("gs://"+project_name+"-csv/\n", ls_out.String()); m {
		fmt.Println("バケット作成済み")
	} else {
		fmt.Println("バケット未作成")
		cmd = exec.Command("gsutil", "mb", "gs://"+project_name+"-csv")
		util.FailOnError(cmd.Run())
		fmt.Println("バケット作成完了")
	}
	//作成したcsvファイルをアップロード
	cmd = exec.Command("gsutil", "cp", "csv/"+mod_date+".csv", "gs://"+project_name+"-csv")
	util.FailOnError(cmd.Run())
	fmt.Println("gcsへのアップロードを完了")
	//BigQueryにデータセットが作成済みかどうかを調べて、未作成なら作成する
	ds_name := strings.Replace(project_name, "-", "_", -1) + "_ds"
	cmd = exec.Command("bq", "ls")
	var ds_out bytes.Buffer
	cmd.Stdout = &ds_out
	util.FailOnError(cmd.Run())
	fmt.Printf("in all caps: %q\n", ds_out.String())
	fmt.Printf("ds_name: %q\n", ds_name)

	if m2, _ := regexp.MatchString("\n  "+ds_name+"\\s+\n", ds_out.String()); m2 {
		fmt.Println("データセット作成済み")
	} else {
		fmt.Println("データセット未作成")
		cmd = exec.Command("bq", "mk", ds_name)
		util.FailOnError(cmd.Run())
		fmt.Println("データセット作成完了")
	}
	//データセット内に既にテーブルがあるかどうかを調べ、なかったときだけロードを実施する（追記されてしまうため）
	cmd = exec.Command("bq", "ls", ds_name)
	var ds_ls_out bytes.Buffer
	cmd.Stdout = &ds_ls_out
	util.FailOnError(cmd.Run())
	fmt.Printf("in all caps: %q\n", ds_ls_out.String())

	if m3, _ := regexp.MatchString("\n\\s+"+mod_date+"\\s+TABLE\\s+\n", ds_ls_out.String()); m3 {
		fmt.Println("当該データロード済み")
	} else {
		fmt.Println("当該データ未ロード")
		//作成したデータセットにデータをロードする
		cmd = exec.Command("bq", "load", "--schema=playdata_schema.json", ds_name+"."+mod_date, "gs://"+project_name+"-csv/"+mod_date+".csv")
		util.FailOnError(cmd.Run())
		fmt.Println("bqへのロードを完了")
	}
	//ロードに使用したcsvファイルをgcsから削除する
	cmd = exec.Command("gsutil", "rm", "gs://"+project_name+"-csv/"+mod_date+".csv")
	util.FailOnError(cmd.Run())
	fmt.Println("gcs上のファイルの削除を完了")
}

func export_csv(mod_date string, playDataList []model.Playdata) {
	// csv ディレクトリがなかったら作る
	util.FailOnError(os.MkdirAll("./csv", 0744))
	filepath := fmt.Sprintf("./csv/%s.csv", mod_date)
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0600)
	util.FailOnError(err)
	defer file.Close()

	err = file.Truncate(0) // ファイルを空にする(同一ファイルに対して2回目以降の実施の場合)
	util.FailOnError(err)

	writer := csv.NewWriter(file)

	for _, data := range playDataList {
		structVal := r.Indirect(r.ValueOf(data))
		typ := structVal.Type()
		var raw []string

		for i := 0; i < typ.NumField(); i++ {
			field := structVal.Field(i)
			value := fmt.Sprintf("%v", field.Interface())
			raw = append(raw, strings.Replace(value, "\n", " ", -1))
		}
		writer.Write(raw)
	}
	writer.Flush()
}
