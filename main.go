package main

import (
	"bufio"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	r "reflect"
	"strings"
)

func main() {
	var fp *os.File
	var err error

	// パース対象の xml ファイル名を引数に受ける
	fp, err = os.Open(os.Args[1])
	failOnError(err)
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)

	var m map[string]string = map[string]string{}
	var playDataList []playData
	var shouldSet bool
	var key string
	d := xml.NewDecoder(reader)
	for {
		token, err := d.Token()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			panic(err)
		}
		switch token.(type) {
		case xml.StartElement:
			if token.(xml.StartElement).Name.Local == "key" {
				shouldSet = false
			} else {
				shouldSet = true
			}
		case xml.EndElement:
			//do nothing
		case xml.CharData:
			if string(token.(xml.CharData)) == "Track ID" {
				// CharData が Track ID が来る＝今まで処理対象だったレコードの終了処理をする ＆ 次のレコードに処理を移す
				d := playData{}
				MapToStruct(m, &d)

				if d.TrackNumber != "" {
					playDataList = append(playDataList, d)
					// fmt.Println(d)
				}

				m = map[string]string{}
			}
			if shouldSet && key != "" {
				m[key] = string(token.(xml.CharData))
				key = ""
			} else {
				key = string(token.(xml.CharData))
			}
		case xml.Comment:
			//do nothing
		case xml.ProcInst:
			//do nothing
		case xml.Directive:
			//do nothing
		default:
			panic("unknown xml token.")
		}
	}

	// 出力する csv ファイル名として用いるために、xml ファイルの最終更新日時を取得
	finfo, err_finfo := fp.Stat()
	failOnError(err_finfo)
	ts := finfo.ModTime()
	mod_date := fmt.Sprintf("%d%02d%02d%02d%02d%02d", ts.Year(), ts.Month(), ts.Day(), ts.Hour(), ts.Minute(), ts.Second())

	export_csv(mod_date, playDataList)
}

func export_csv(mod_date string, playDataList []playData) {
	// csv ディレクトリがなかったら作る
	failOnError(os.MkdirAll("./csv", 0744))
	filepath := fmt.Sprintf("./csv/%s.csv", mod_date)
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)
	defer file.Close()

	err = file.Truncate(0) // ファイルを空にする(同一ファイルに対して2回目以降の実施の場合)
	failOnError(err)

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

func MapToStruct(mapVal map[string]string, val interface{}) (ok bool) {
	structVal := r.Indirect(r.ValueOf(val))
	for name, elem := range mapVal {
		// ここで来る name は　plist の key エレメントの値なのでスペースを含んでいる。
		// 一方 struct のフィールド名からはスペースを除去している。
		f := structVal.FieldByName(strings.Replace(name, " ", "", -1))
		if f.IsValid() {
			f.Set(r.ValueOf(elem))
		}
	}
	return
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
		panic(err)
	}
}

type playData struct {
	Artist               string
	Album                string
	Genre                string
	Kind                 string
	Size                 string
	TotalTime            string
	TrackNumber          string
	DateModified         string
	DateAdded            string
	BitRate              string
	SampleRate           string
	PersistentID         string
	TrackType            string
	Location             string
	FileFolderCount      string
	LibraryFolderCount   string
	PlayCount            string
	PlayDate             string
	PlayDateUTC          string
	Rating               string
	AlbumRating          string
	AlbumRatingComputed  string
	ArtworkCount         string
	SkipCount            string
	SkipDate             string
	Disabled             string
	SortAlbum            string
	Year                 string
	Comments             string
	SortName             string
	SortArtist           string
	VolumeAdjustment     string
	AlbumArtist          string
	DiscNumber           string
	DiscCount            string
	TrackCount           string
	ReleaseDate          string
	Protected            string
	Purchased            string
	Compilation          string
	Composer             string
	HasVideo             string
	VideoWidth           string
	VideoHeight          string
	Movie                string
	Master               string
	PlaylistID           string
	PlaylistPersistentID string
	Visible              string
	AllItems             string
	PlaylistItems        string
	SmartCriteria        string
}
