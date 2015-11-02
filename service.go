package main

import "net/http"
import "os"
import "log"
import "strings"
import "io/ioutil"
import "github.com/takama/daemon"

type Service struct {
	daemon.Daemon
}

var dependencies = []string{"dummy.service"}
var stdlog, errlog *log.Logger

func (service *Service) Manage() (string, error) {
	usage := "Usage:image-assistant install | remove | start | stop | status"

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		url := r.Form["url"][0]
		stdlog.Println("Image Url: " + url)

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
		ioutil.WriteFile("/Users/gaochenfei/Dropbox/Pictures/"+fileName+"."+fileType, body, 0644)
		w.Write([]byte("/Users/gaochenfei/Dropbox/Pictures/" + fileName + "." + fileType))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		errlog.Println("Server error")
		os.Exit(0)
	}

	return usage, nil
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func main() {
	srv, err := daemon.New("image-assistantd", "image", dependencies...)
	if err != nil {
		errlog.Println(err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(err)
		os.Exit(1)
	}
	stdlog.Println(status)
}
