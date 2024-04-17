package main

import (
	"fmt"
	"os"
)

// plugin name
func main() {

	command := os.Args

	if len(command) == 1 {
		fmt.Printf("Please set microserivice name")
		return
	}

	msname := command[1]

	CreateMicroservice(msname)

}
