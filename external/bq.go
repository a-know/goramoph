package external

import (
	"../util"
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

func IsDatasetExists(project_name string) bool {
	ds_name := GenerateDatasetName(project_name)
	cmd := exec.Command("bq", "ls")
	var ds_out bytes.Buffer
	cmd.Stdout = &ds_out
	util.FailOnError(cmd.Run())

	m, _ := regexp.MatchString("\n  "+ds_name+"\\s+\n", ds_out.String())
	return m
}

func GenerateDatasetName(project_name string) string {
	return strings.Replace(project_name, "-", "_", -1) + "_ds"
}

func MakeDataset(project_name string) {
	cmd := exec.Command("bq", "mk", GenerateDatasetName(project_name))
	util.FailOnError(cmd.Run())
}

func IsTableExists(project_name, mod_date string) bool {
	cmd := exec.Command("bq", "ls", GenerateDatasetName(project_name))
	var ds_ls_out bytes.Buffer
	cmd.Stdout = &ds_ls_out
	util.FailOnError(cmd.Run())

	m, _ := regexp.MatchString("\n\\s+"+mod_date+"\\s+TABLE\\s+\n", ds_ls_out.String())
	return m
}

func LoadToTable(project_name, mod_date string) {
	cmd := exec.Command("bq", "load", "--schema=playdata_schema.json", GenerateDatasetName(project_name)+"."+mod_date, util.GenerateBucketName(project_name)+mod_date+".csv")
	util.FailOnError(cmd.Run())
}
