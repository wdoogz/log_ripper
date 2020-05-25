package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func regexSearch(regFile string) int {
	r, _ := regexp.Compile("[eE][rR]{2}[oO][rR]")
	results := r.FindAllString(string(regFile), -1)
	return len(results)
}

// ParseLogFile opens file, and parses for error
func ParseLogFile(fileName string) int {
	logFile, _ := ioutil.ReadFile(fileName)

	res := regexSearch(string(logFile))
	return res
}

// ParseLogDir parses dir
func ParseLogDir(dirName string) int {
	var totalErrors int
	logDir, err := ioutil.ReadDir(dirName)
	for _, dirLogFile := range logDir {
		logFileName, _ := ioutil.ReadFile(dirName + "/" + dirLogFile.Name())

		//fmt.Println(dirLogFile.Name())
		results := regexSearch(string(logFileName))
		totalErrors = totalErrors + results
	}

	if err != nil {
		fmt.Println(err)
	}

	return totalErrors
}
