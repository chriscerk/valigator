package main

// not in use yet
func applySpecification(s specification) <-chan string {
	c := make(chan string)

	go func() {
		if s.Filename == "all" {
			allFilesChannel := searchAllFiles(s, c)
			go func() {
				for {
					c <- <-allFilesChannel
				}
			}()
		} else if fileExists(s.Filename) {
			searchFile(s, s.Filename)
		} else {
			c <- "|  ERROR  | " + s.Name + " | " + s.Filename + " does not exist! \n"
		}
	}()
	return c
}
