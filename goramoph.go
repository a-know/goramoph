package main

import (
	"./exporter"
	"./external"
	"./model"
	"./parser"
	"./util"
	"fmt"
	"os"
	"os/exec"
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
		external.MakeBucket(project_name)
		fmt.Println("バケット作成完了")
	}
	//作成したcsvファイルをアップロード
	external.FileUpload(project_name, mod_date)
	fmt.Println("gcsへのアップロードを完了")
	//BigQueryにデータセットが作成済みかどうかを調べて、未作成なら作成する
	if external.IsDatasetExists(project_name) {
		fmt.Println("データセット作成済み")
	} else {
		fmt.Println("データセット未作成")
		external.MakeDataset(project_name)
		fmt.Println("データセット作成完了")
	}
	//データセット内に既にテーブルがあるかどうかを調べ、なかったときだけロードを実施する（追記されてしまうため）
	if external.IsTableExists(project_name, mod_date) {
		fmt.Println("当該データロード済み")
	} else {
		fmt.Println("当該データ未ロード")
		//作成したデータセットにデータをロードする
		external.LoadToTable(project_name, mod_date)
		fmt.Println("bqへのロードを完了")
	}
	//ロードに使用したcsvファイルをgcsから削除する
	cmd := exec.Command("gsutil", "rm", "gs://"+project_name+"-csv/"+mod_date+".csv")
	util.FailOnError(cmd.Run())
	fmt.Println("gcs上のファイルの削除を完了")
}
