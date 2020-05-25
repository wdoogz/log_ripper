package main

import (
	"fmt"
	"os"
	"parser"
)

func main() {
	// "/Users/wdugan/Documents/PersonalGit/log_ripper/log_files_test"
	pathToLogs := os.Args[1]
	testPath, err := os.Stat(pathToLogs)

	if testPath.IsDir() {
		parser.ParseLogDir(pathToLogs)
	} else {
		parser.ParseLogFile(pathToLogs)
	}
	if err != nil {
		fmt.Println(err)
	}
}
