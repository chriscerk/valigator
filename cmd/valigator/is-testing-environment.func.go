package main

func isTestingEnvironment() bool {
	testFile := "../../test/specifications.json"
	if fileExists(testFile) {
		return true

	} else {
		return false
	}
}
