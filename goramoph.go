package main

import (
	"./exporter"
	"./external"
	"./model"
	"./parser"
	"./util"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting goramoph.")
	var fp *os.File
	var err error
	// 起動時引数の検証
	if len(os.Args) != 2 {
		fmt.Println("Usage: goramoph.go [path to iTunes Music Library.xml]")
		os.Exit(1)
	} else if _, err := os.Stat(os.Args[1]); err != nil {
		msg := fmt.Sprintf("file %s is not exists.", os.Args[1])
		fmt.Println(msg)
		os.Exit(1)
	}

	// パース対象の xml ファイル名を引数に受ける
	fmt.Printf("xml parsing...")
	fp, err = os.Open(os.Args[1])
	util.FailOnError(err)
	defer fp.Close()

	var playDataList []model.Playdata
	playDataList = parser.ItunesXmlParse(fp)

	// 出力する csv ファイル名として用いるために、xml ファイルの最終更新日時を取得
	mod_date := util.GetModDate(fp)
	fmt.Println("done")

	fmt.Printf("exporting csv...")
	exporter.ExportCsv(mod_date, playDataList)
	fmt.Println("done")

	//外部コマンドを実行し、プロジェクト名を取得する
	fmt.Printf("getting Google Cloud Project name...")
	project_name := external.GetProjectName()
	fmt.Println("project_name:", string(project_name))
	//バケットが既に作成済みかどうかを調べて、未作成なら作成する
	fmt.Printf("checking bucket...")
	if external.IsBucketExists(project_name) {
		fmt.Println("bucket has alredy created")
	} else {
		fmt.Printf("bucket has not created yet, now creating...")
		external.MakeBucket(project_name)
		fmt.Println("done")
	}
	//作成したcsvファイルをアップロード
	fmt.Printf("uploading csv...")
	external.FileUpload(project_name, mod_date)
	fmt.Println("done")
	//BigQueryにデータセットが作成済みかどうかを調べて、未作成なら作成する
	fmt.Printf("checking dataset...")
	if external.IsDatasetExists(project_name) {
		fmt.Println("dataset has alredy created")
	} else {
		fmt.Printf("dataset has not created yet, now creating...")
		external.MakeDataset(project_name)
		fmt.Println("done")
	}
	//データセット内に既にテーブルがあるかどうかを調べ、なかったときだけロードを実施する（追記されてしまうため）
	fmt.Printf("checking table...")
	if external.IsTableExists(project_name, mod_date) {
		fmt.Println("this playdata has alredy loaded")
	} else {
		fmt.Printf("this playdata has not load yet, now loading...")
		//作成したデータセットにデータをロードする
		external.LoadToTable(project_name, mod_date)
		fmt.Println("load to table success")
	}
	//ロードに使用したcsvファイルをgcsから削除する
	fmt.Printf("remove csv file on Cloud Storage...")
	external.RemoveUploadFile(project_name, mod_date)
	fmt.Println("done")
	fmt.Println("completed.")
}
