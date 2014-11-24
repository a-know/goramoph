package main

import (
	"./exporter"
	"./external"
	"./model"
	"./parser"
	"./util"
	"bytes"
	"fmt"
	"os"
	"os/exec"
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

	exporter.ExportCsv(mod_date, playDataList)

	//外部コマンドを実行し、プロジェクト名を取得する
	project_name := external.GetProjectName()
	fmt.Println("project_name:", string(project_name))
	//バケットが既に作成済みかどうかを調べて、未作成なら作成する
	if external.IsBucketExists(project_name) {
		fmt.Println("バケット作成済み")
	} else {
		fmt.Println("バケット未作成")
		cmd := exec.Command("gsutil", "mb", "gs://"+project_name+"-csv")
		util.FailOnError(cmd.Run())
		fmt.Println("バケット作成完了")
	}
	//作成したcsvファイルをアップロード
	cmd := exec.Command("gsutil", "cp", "csv/"+mod_date+".csv", "gs://"+project_name+"-csv")
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
