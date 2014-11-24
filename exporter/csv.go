package exporter

import (
	"../model"
	"../util"
	"encoding/csv"
	"fmt"
	"os"
	r "reflect"
	"strings"
)

func ExportCsv(mod_date string, playDataList []model.Playdata) {
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
