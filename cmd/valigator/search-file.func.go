package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func searchFile(s specification, fileName string) <-chan string {
	c := make(chan string)

	if isTestingEnvironment() && fileName != "" {
		fileName = "../../test/" + fileName
	}

	go func() {
		if s.Filename == "all" {
			allFilesChannel := searchAllFiles(s, c)
			go func() {
				for {
					c <- <-allFilesChannel
				}
			}()
		} else if fileExists(fileName) {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			defer file.Close()

			re := regexp.MustCompile(s.Pattern)
			totalMatches := 0

			scanner := bufio.NewScanner(file)
			lineNumber := 1
			for scanner.Scan() {
				line := scanner.Text()
				matches := re.FindStringSubmatch(line)

				if len(matches) > 0 {
					if s.Contains {
						c <- fmt.Sprintf("| SUCCESS | %s | Matches %s in %s on line %d \n", s.Name, matches[0], fileName, lineNumber)
					} else {
						c <- fmt.Sprintf("|  ERROR  | %s | %s in %s on line %d \n", s.Name, s.Error, fileName, lineNumber)
					}

					totalMatches++
				}
				lineNumber++
			}

			if totalMatches == 0 && s.Contains == true {
				c <- " | ERROR | " + s.Name + s.Error + " in " + s.Filename + " - " + s.Resolution + "\n"
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
			//c <- s.Name + "-DONE \n"
			close(c)
		} else {
			c <- "|  ERROR  | " + s.Name + " | " + s.Filename + " does not exist! \n"
		}
	}()
	return c
}
