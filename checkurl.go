package main

import (
	"net/http"
	"time"
)

// func main() {
// 	ok := IsURLOK("https://codereview.stackexchange.com/questi")
// 	fmt.Println(ok)
// }

func IsURLOK(url string) bool {
	client := &http.Client{Timeout: time.Duration(3) * time.Second}
	res, err := client.Head(url)
	if err != nil {
		return false
	}
	if res.StatusCode > 400 {
		return false
	}

	if res.ContentLength > 0 {
		return true
	}
	return false
}
