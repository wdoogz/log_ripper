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
	databaseName := os.Args[2]
	databaseUser := os.Args[3]
	databasePass := os.Args[4]
	testPath, err := os.Stat(pathToLogs)

	if testPath.IsDir() {
		numberOfErrors = parser.ParseLogDir(pathToLogs)
		fmt.Println(numberOfErrors)
		poster.Poster("http://192.168.1.10:8086", databaseName, databaseUser, databasePass)
	} else {
		numberOfErrors = parser.ParseLogFile(pathToLogs)
		fmt.Println(numberOfErrors)
		poster.Poster("http://192.168.1.10:8086", databaseName, databaseUser, databasePass)
	}
	if err != nil {
		fmt.Println(err)
	}
}
