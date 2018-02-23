// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
				fmt.Fprint(w, "illegal url")
				return
			}
		}
		detectSafeSearch(w, url)
	})

	nodego.TakeOver()
}
