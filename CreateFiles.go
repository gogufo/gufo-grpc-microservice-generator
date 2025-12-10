package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CreateFiles(name string, filepath string) {

	endfiles := []string{
		"main.go",
		"config/settings.example.toml",
		"config/env.example",
		"entrypoint/entrypoint.go",
		"entrypoint/update.go",
		"entrypoint/init.go",
		"version/version.go",
		"CHANGELOG.md",
		"Dockerfile",
		"docker-compose.yaml",
		"init.go",
		"README.md",
		"admin/init.go",
		"go.mod",
		"go.sum",
		"docs/README.md", //
		"cron/init.go",
		"global/constants.go",
		"global/variables.go",
		"admin/get/CheckCron.go",
		"admin/post/UpdateCron.go",
		"metrics/metrics.go",
		"middleware/logging.go",
		"middleware/metrics.go",
		"middleware/recovery.go",
		"middleware/reuestid.go",
		"middleware/security.go",
		".gitignore",
		"heratbeat.go",
		"Makefile",
		"metrics_local.go",
		"processUploadedFile.go",
		"tests/health_tests.go",
	}

	datafiles := []string{
		"main",
		"config",
		"env",
		"entrypoint",
		"update",
		"entrypointinit",
		"version",
		"changelog",
		"dockerfile",
		"docker-compose",
		"init",
		"readme",
		"admininit",
		"mod",
		"sum",
		"docs", //
		"croninit",
		"consts",
		"vars",
		"admingetcheck",
		"adminpostupdate",
		"metrics",
		"logging",
		"m-metrics",
		"recovery",
		"requestid",
		"security",
		"gitignore",
		"heratbeat",
		"makefile",
		"metrics_local",
		"processUploadedFile",
		"health_tests",
	}

	for i := 0; i < len(endfiles); i++ {

		//1. Create file
		maingofilepath := fmt.Sprintf("%s/%s", filepath, endfiles[i])
		file, err := os.OpenFile(maingofilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		defer file.Close()

		// read the whole file at once
		readfile := fmt.Sprintf("/data/%s", datafiles[i])
		maintxt, err := ioutil.ReadFile(readfile)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

		updatedmaintxt := strings.Replace(string(maintxt), "{{name}}", name, -1)
		updatedmaintxt = strings.Replace(string(updatedmaintxt), "{{name_up}}", strings.ToUpper(name), -1)
		updatedmaintxt = strings.Replace(string(updatedmaintxt), "{{version}}", VERSIONPLUGIN, -1)

		err = ioutil.WriteFile(maingofilepath, []byte(updatedmaintxt), 0644)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

	}

}
