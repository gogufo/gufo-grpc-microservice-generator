package main

import (
	"os"
	"path/filepath"
)

func CreateFolders(name string) string {
	//1. Create folders
	// Create folders in /src folder.
	foldersarray := []string{"config", "entrypoint", "version", "admin", "docs", "global", "cron", "admin/post", "admin/get", "tests", "metrics", "middleware"}
	initFilePath := "/src"
	m9s := filepath.Join(initFilePath, name)

	//Create Microservice Folder
	if _, err := os.Stat(m9s); os.IsNotExist(err) {
		os.Mkdir(m9s, 0755)
	}

	for i := 0; i < len(foldersarray); i++ {
		//Create Config Folder
		newfolder := filepath.Join(m9s, foldersarray[i])
		//Create dir output using above code
		if _, err := os.Stat(newfolder); os.IsNotExist(err) {
			os.Mkdir(newfolder, 0755)
		}

	}

	return m9s
}
