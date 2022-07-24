package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name        string `json:"Name"`
	Age         int    `json:"Age"`
	Active      bool   `json:"Active"`
	lastLoginAt string `json:"lastLoginAt"` // unexported fields are not considered while marshaling/unmarshalling
	//	Age         int    `json:"Age,omitempty"`
	// LastLoginAt string `json:"lastLoginAt"`
}

var jsonString = []byte(`{"Name":"Bob","Age":10,"Active":true,"lastLoginAt":"12345"}`)
var invalidJsonString = []byte(`{"Name":"Bob,,"Age":10,"Active":true,"lastLoginAt":"12345"}`)

func JsonDemo() {
	// Marshall Json
	u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))

	// Unmarshall Json
	var user User
	err = json.Unmarshal(jsonString, &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", user)

	// // fail marshall
	//var ch = make(chan int)
	//_, err := json.Marshal(ch)
	//if err != nil {
	//	panic(err)
	//}

	//// fail unmarshall
	//
	//err = json.Unmarshal(invalidJsonString, &user)
	//if err != nil {
	//	panic(err)
	//}

}

func JsonEncodingDemo() {
	var user []User
	f, err := os.Open("user_list.json")
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	mrgObj := json.NewDecoder(f)
	err = mrgObj.Decode(&user)
	if err != nil {
		log.Fatalf("dec.Decode() failed with '%s'\n", err)
	}
	fmt.Printf("%+v \n", user)

	// encoding it to write back it a txt file

	txtFile, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(txtFile)

	en := json.NewEncoder(txtFile)
	err = en.Encode(user)
	if err != nil {
		log.Fatalf("encode failed with '%s'\n", err)
	}
}
