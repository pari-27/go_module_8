package main

import (
	"fmt"
	"os"
)

func EnvVariablesDemo() {
	// Set Env variable  :
	os.Setenv("module8", "this is module8 value")
	// Get Env variable :
	envValue := os.Getenv("module8")
	fmt.Printf("got value for key module 8 :%v \n", envValue)
	// Unset Env variable :
	os.Unsetenv("module8")
	envValue = os.Getenv("module8")
	fmt.Printf("value for key module8 after unseting it :%v \n", envValue)
	// Lookup env variable :
	if _, ok := os.LookupEnv("module8"); !ok {
		fmt.Printf("did not find value for key module8\n")
	}

	//List All Env variables :
	envs := os.Environ()
	for _, env := range envs {
		fmt.Println(env)
	}

}
