package main

import "github.com/flex1988/go-level-logger"
import "net/http"
import "io/ioutil"
import "strings"
import "os"

func main() {
	logger, _ := logger.New()

	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		url := r.Form["url"][0]
		logger.Info("Image Url: " + url)

		urlParts := strings.Split(url, "/")
		var fileName = ""
		if r.Form["name"] == nil {
			fileName = urlParts[len(urlParts)-1]
		} else {
			fileName = r.Form["name"][0]
		}

		resp, _ := http.Get(url)
		defer resp.Body.Close()
		fileType := strings.Split(resp.Header["Content-Type"][0], "/")[1]

		body, _ := ioutil.ReadAll(resp.Body)
		ioutil.WriteFile("./download/"+fileName+"."+fileType, body, 777)
	})

	logger.Info("Server started on:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		logger.Error("Server error")
		os.Exit(0)
	}
}
