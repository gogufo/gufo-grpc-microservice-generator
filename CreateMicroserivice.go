package main

import "fmt"

func CreateMicroservice(name string) {

	//1. Create folders
	filepath := CreateFolders(name)

	//2 Create Files
	CreateFiles(name, filepath)

	fmt.Printf("Done \n")

}
