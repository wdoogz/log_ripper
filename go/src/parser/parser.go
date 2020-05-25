package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// ParseLogFile opens file, and parses for error
func ParseLogFile(fileName string) {
	logFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error!!... ->", err)
	}

	r, _ := regexp.Compile("[errorERROR]?[failFAIL]?")
	results := r.FindAllString(string(logFile), -1)
	fmt.Println(len(results))

}

// ParseLogDir parses dir
func ParseLogDir(dirName string) {
	logDir, err := ioutil.ReadDir(dirName)
	for _, dirLogFile := range logDir {
		logFileName, _ := ioutil.ReadFile(dirName + "/" + dirLogFile.Name())

		fmt.Println(dirLogFile.Name())
		r, _ := regexp.Compile("[errorERROR]?[failFAIL]?")
		results := r.FindAllString(string(logFileName), -1)
		fmt.Println(len(results))

		if err != nil {
			fmt.Println(err)
		}
	}

	if err != nil {
		fmt.Println(err)
	}
}
