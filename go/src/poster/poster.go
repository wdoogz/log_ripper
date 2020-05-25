package poster

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//Poster posts
func Poster(host string, databaseName string, databaseUser string, databasePass string, errorCount int) []byte {
	var client = &http.Client{}
	var currentTime = strconv.FormatInt(time.Now().Unix(), 10)
	var hostname, _ = os.Hostname()
	var postURL string = "http://" + host + "/write?db=" + databaseName
	var newErrorCount string = strconv.Itoa(errorCount)
	var data = strings.NewReader(`log_errors,host=` + hostname + ` value=` + newErrorCount + ` ` + currentTime + `000000000`)

	r, err := http.NewRequest("POST", postURL, data)
	r.Header.Set("Authorization", "Token "+databaseUser+":"+databasePass)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return bodyText
}
