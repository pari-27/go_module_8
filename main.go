package main

import (
	"fmt"
	"os"
)

func main() {
	cmdArgs := os.Args
	operation := cmdArgs[1]
	switch operation {
	case "fileops":
		FileOperationDemo()
	case "envosops":
		EnvVariablesDemo()
	case "json":
		JsonDemo()
	case "jsonencoding":
		JsonEncodingDemo()
	default:
		fmt.Println("invalid argument")
	}
	//FileOperationDemo()
	//EnvVariablesDemo()
	//JsonDemo()
	//JsonEncodingDemo()

}
