package main

import (
	"fmt"
	"log"
	"os"
)

func FileOperationDemo() {

	// create a file
	out, err := os.Create("novel.txt") //give any name
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}

	fmt.Printf("File created ---------> %v \n", out.Name())
	defer out.Close()
	// open a file
	txtFile, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(txtFile)

	fmt.Printf("File name opened ---------> %v %v \n", txtFile.Name())

	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	//O_RDONLY  // open the file read-only.
	//O_WRONLY  // open the file write-only.
	//O_RDWR    // open the file read-write.
	//// The remaining values may be or'ed in to control behavior.
	//O_APPEND  // append data to the file when writing.
	//O_CREATE  // create a new file if none exists.
	//O_EXCL    // used with O_CREATE, file must not exist.
	//O_SYNC    // open for synchronous I/O.
	//O_TRUNC   // truncate regular writable file when opened.

	// read from a file

	var strOutput []byte
	strOutput, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("os.ReadFile() failed with '%s'\n", err)
	}
	fmt.Printf("Output from input.txt file ------------> \n %v \n", string(strOutput))

	// write to a file

	str := "New text to be added to file output.txt"
	err = os.WriteFile("output.txt", []byte(str), 0777)
	if err != nil {
		log.Fatalf("os.Write() failed with '%s'\n", err)
	}

	// if the file does not exist then writeFile creates one with the permission specified
	strNew := "New text to be added to new file"
	err = os.WriteFile("newFileCreated.txt", []byte(strNew), 0777)
	if err != nil {
		log.Fatalf("os.Write() failed with '%s'\n", err)
	}

	// fileInfo

	fileinfo, err := os.Stat("input.txt")
	if err != nil {
		log.Fatalf("os.Write() failed with '%s'\n", err)
	}
	fmt.Printf("File info : \n Name : %v \n Size : %v \n Mode : %v \n ModTime : %v \n IsDir : %v \n Sys : %v \n", fileinfo.Name(), fileinfo.Size(),
		fileinfo.Mode(), fileinfo.ModTime(), fileinfo.IsDir(), fileinfo.Sys())
	// exit

	os.Exit(0) // range upto 0-125
	fmt.Println("will this get printed ?")
}
