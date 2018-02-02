package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func searchAllFiles(s specification, c <-chan string) <-chan string {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	searchDir := path.Dir(filename)

	fileList := []string{}
	walkErr := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		isDirectory := f.IsDir()
		if !isDirectory {
			fileList = append(fileList, path)
		}
		return nil
	})

	if walkErr != nil {
		log.Fatal(walkErr)
	}

	var channels []<-chan string

	for _, file := range fileList {
		c := searchFile(s, file)
		channels = append(channels, c)
	}

	return fanIn(channels)
}
