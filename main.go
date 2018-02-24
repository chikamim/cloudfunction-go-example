package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/cloud-functions-go/nodego"
)

func init() {
	nodego.OverrideLogger()
}

func main() {
	flag.Parse()
	http.HandleFunc(nodego.HTTPTrigger, func(w http.ResponseWriter, r *http.Request) {
		log.Println("This is a log message from Go!")
		url := r.URL.RawQuery
		if !strings.HasPrefix(url, "gs://") {
			ok := IsURLOK(url)
			if !ok {
				fmt.Fprint(w, "url is unreachable")
				return
			}
		}
		detectSafeSearch(w, url)
	})

	nodego.TakeOver()
}
