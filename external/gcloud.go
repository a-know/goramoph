package external

import (
	"../util"
	"bytes"
	"os/exec"
	"regexp"
)

func GetProjectName() (project_name string) {
	cmd := exec.Command("gcloud", "config", "list")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	util.FailOnError(cmd.Run())

	//外部コマンドの実行結果からプロジェクト名を抽出する
	reg, _ := regexp.Compile("\nproject = (.+)\n")
	result := reg.Find([]byte(stdout.String()))
	replace_reg, _ := regexp.Compile("project|\\s|\n|=")
	project_name = replace_reg.ReplaceAllString(string(result), "")
	return
}
