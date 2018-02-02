package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Scanning your project...")

	specifications := getSpecifications()
	var channels []<-chan string

	for _, specification := range specifications {
		c := searchFile(specification, specification.Filename)
		channels = append(channels, c)
	}

	mainChannel := fanIn(channels)

	i := 1
	for i < 100 {
		fmt.Printf(<-mainChannel)
		time.Sleep(1 * time.Millisecond)
		i++
	}

	fmt.Println("All Scanning complete")
	fmt.Println("Press 'Enter' to end...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
