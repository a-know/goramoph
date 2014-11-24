package parser

import (
	"../model"
	"bufio"
	"encoding/xml"
	"io"
	"os"
	r "reflect"
	"strings"
)

func ItunesXmlParse(fp *os.File) []model.Playdata {

	reader := bufio.NewReaderSize(fp, 4096)

	var m map[string]string = map[string]string{}
	var playDataList []model.Playdata
	var should_set bool
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
				should_set = false
			} else {
				should_set = true
			}
		case xml.EndElement:
			//do nothing
		case xml.CharData:
			if string(token.(xml.CharData)) == "Track ID" {
				// CharData が Track ID が来る＝今まで処理対象だったレコードの終了処理をする ＆ 次のレコードに処理を移す
				d := model.Playdata{}
				mapToStruct(m, &d)

				if d.TrackNumber != "" {
					playDataList = append(playDataList, d)
				}

				m = map[string]string{}
			}
			if should_set && key != "" {
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

	return playDataList
}

func mapToStruct(mapVal map[string]string, val interface{}) (ok bool) {
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
