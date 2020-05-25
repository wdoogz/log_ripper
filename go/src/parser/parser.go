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

	r, _ := regexp.Compile("error")
	results := r.FindAllString(string(logFile), -1)
	fmt.Println(len(results))

}

// ParseLogDir parses dir
func ParseLogDir(dirName string) {
	logDir, err := ioutil.ReadDir(dirName)
	for _, dirLogFile := range logDir {
		logFileName, err := ioutil.ReadFile(dirName + "/" + dirLogFile.Name())
		fmt.Println(string(logFileName))

		if err != nil {
			fmt.Println(err)
		}
	}

	if err != nil {
		fmt.Println(err)
	}
}
