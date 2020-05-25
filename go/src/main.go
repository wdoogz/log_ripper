package main

import (
	"fmt"
	"os"
	"parser"
	"poster"
)

func main() {
	var numberOfErrors int
	if len(os.Args) != 6 {
		fmt.Println("\n\nUsage: \n>> log_ripper <path to logs> <hostname/url and port> <database name> <database username> <database password>\n\n")
		fmt.Println("\nExample: \n>> log_ripper /var/log/messages 192.168.1.10:8086 mydb myuser mypass")
	}

	pathToLogs := os.Args[1]
	databaseHost := os.Args[2]
	databaseName := os.Args[3]
	databaseUser := os.Args[4]
	databasePass := os.Args[5]

	testPath, err := os.Stat(pathToLogs)
	if err != nil {
		fmt.Println(err)
	}

	if testPath.IsDir() {
		numberOfErrors = parser.ParseLogDir(pathToLogs)
		fmt.Println(numberOfErrors)
		poster.Poster(databaseHost, databaseName, databaseUser, databasePass, numberOfErrors)
	} else {
		numberOfErrors = parser.ParseLogFile(pathToLogs)
		fmt.Println(numberOfErrors)
		poster.Poster(databaseHost, databaseName, databaseUser, databasePass, numberOfErrors)
	}

}
