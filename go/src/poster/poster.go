package poster

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Poster posts
func Poster(url string, databaseName string, databaseUser string, databasePass string) {
	var client = &http.Client{}
	var currentTime = strconv.FormatInt(time.Now().Unix(), 10)
	var data = strings.NewReader(`cpu_load_short,host=server01,region=us-west value=0.64` + currentTime + `000000000`)

	r, err := http.NewRequest("POST", url+"/write?db="+databaseName, data)
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
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
