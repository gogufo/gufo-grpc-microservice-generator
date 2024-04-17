package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CreateFiles(name string, filepath string) {

	endfiles := []string{"main.go", "config/config.toml", "entrypoint/entrypoint.go", "entrypoint/init.go", "version/version.go", "CHANGELOG.md", "Dockerfile", "init.go", "README.md", "admin/init.go", "go.mod", "go.sum", "docs/README.md"}
	datafiles := []string{"main", "config", "entrypoint", "entrypointinit", "version", "changelog", "dockerfile", "init", "readme", "admininit", "mod", "sum", "docs"}

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

		err = ioutil.WriteFile(maingofilepath, []byte(updatedmaintxt), 0644)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

	}

}