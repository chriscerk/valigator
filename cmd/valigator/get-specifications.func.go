package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func getSpecifications() []specification {

	var c []specification
	var rawFileContents []byte
	var err error

	prodFile := "./specifications.json"
	testFile := "../../test/specifications.json"

	if isTestingEnvironment() {
		rawFileContents, err = ioutil.ReadFile(testFile)
	} else if fileExists(prodFile) {
		rawFileContents, err = ioutil.ReadFile(prodFile)
	} else {
		log.Fatal("Specifications.json does not exist")
	}

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(rawFileContents, &c)
	return c
}
