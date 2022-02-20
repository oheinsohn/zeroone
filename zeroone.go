package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type SomeTestData struct {
	SomeTestData []ZeroOneData `json:"someTestData"`
}

type ZeroOneData struct {
	Name   string `json:"name"`
	Rating string `json:"rating"`
}

func main() {

	// Start up and versioning
	var version = 0
	fmt.Println("######### ZeroOne start ##############")
	fmt.Printf("Version: %09d \n", version)
	var zeroone_parent_pid = os.Getpid()
	fmt.Printf("ZeroOne parent pid: %d \n", zeroone_parent_pid)
	var open_session_number = 0

	// Starting with some test data
	readSomeTestData()

	// Make final checks before ending
	fmt.Println("Final number of open sessions: ", open_session_number)
	if open_session_number == 0 {
		fmt.Println("No open session left.")
	} else {
		fmt.Println("WARNING - There are open sessions left.")
	}
	time.Sleep(time.Second)
	fmt.Println("######### ZeroOne end ##############")
}

func readSomeTestData() {
	fmt.Println("Reading some test data.")

	scopeFile, err := os.Open("someTestData.json")
	if err != nil {
		fmt.Println("Could not read json file.", err)
	}
	defer scopeFile.Close()
	byteValue, _ := ioutil.ReadAll(scopeFile)
	var someTestData SomeTestData
	json.Unmarshal(byteValue, &someTestData)

	fmt.Println("Entries: ", len(someTestData.SomeTestData))
	for i := 0; i < len(someTestData.SomeTestData); i++ {
		fmt.Println("Name: " + someTestData.SomeTestData[i].Name)
		fmt.Println("Rating: " + someTestData.SomeTestData[i].Rating)
	}
}
