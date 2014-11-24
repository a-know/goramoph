package util

import (
	"fmt"
	"log"
	"os"
)

func FailOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
		panic(err)
	}
}

func GetModDate(fp *os.File) (mod_date string) {
	finfo, err_finfo := fp.Stat()
	FailOnError(err_finfo)
	ts := finfo.ModTime()
	mod_date = fmt.Sprintf("%d%02d%02d%02d%02d%02d", ts.Year(), ts.Month(), ts.Day(), ts.Hour(), ts.Minute(), ts.Second())
	return
}

func GenerateBucketName(project_name string) string {
	return "gs://" + project_name + "-csv/"
}
