package external

import (
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
