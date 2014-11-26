package external

import (
	"../exporter"
	"../util"
	"bytes"
	"os/exec"
	"regexp"
)

func IsBucketExists(project_name string) bool {
	cmd := exec.Command("gsutil", "ls")
	var ls_out bytes.Buffer
	cmd.Stdout = &ls_out
	util.FailOnError(cmd.Run())

	m, _ := regexp.MatchString(util.GenerateBucketName(project_name), ls_out.String())
	return m
}

func MakeBucket(project_name string) {
	cmd := exec.Command("gsutil", "mb", util.GenerateBucketName(project_name))
	util.FailOnError(cmd.Run())
}

func FileUpload(project_name, mod_date string) {
	cmd := exec.Command("gsutil", "cp", exporter.GetCsvFilepath(mod_date), util.GenerateBucketName(project_name))
	util.FailOnError(cmd.Run())
}

func RemoveUploadFile(project_name, mod_date string) {
	cmd := exec.Command("gsutil", "rm", util.GenerateUploadFilepath(project_name, mod_date))
	util.FailOnError(cmd.Run())
}
