package poster

import (
	"fmt"
	"net/http"
)

//Poster posts
func Poster(url string) {
	r, err := http.Post(url, "hello", nil)
	r.Header.Add("Authorization", "Token username:password")
	r.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
