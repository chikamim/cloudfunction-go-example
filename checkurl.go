package main

import (
	"net/http"
	"time"
)

const RequestTimeoutSec = 3

func IsURLOK(url string) bool {
	client := &http.Client{Timeout: time.Duration(RequestTimeoutSec) * time.Second}
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
