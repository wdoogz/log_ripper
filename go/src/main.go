package main

import (
	"fmt"
	"os"
	"parser"
	"poster"
)

func main() {
	// "/Users/wdugan/Documents/PersonalGit/log_ripper/log_files_test"
	var numberOfErrors int
	pathToLogs := os.Args[1]
	testPath, err := os.Stat(pathToLogs)

	if testPath.IsDir() {
		numberOfErrors = parser.ParseLogDir(pathToLogs)
		fmt.Println(numberOfErrors)
	} else {
		numberOfErrors = parser.ParseLogFile(pathToLogs)
		fmt.Println(numberOfErrors)
	}
	if err != nil {
		fmt.Println(err)
	}
	poster.Poster("https://httpbin.org/post")
}
