package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	Parser("/Users/wdugan/Documents/PersonalGit/log_ripper/test.txt")
}

// Parser opens file, and parses for error
func Parser(fileName string) {
	logFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error!... ->", err)
	}

	r, _ := regexp.Compile("error")
	results := r.FindAllString(string(logFile), -1)
	fmt.Println(len(results))

}
