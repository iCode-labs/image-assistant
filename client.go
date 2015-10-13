package main

import (
	"encoding/json"
	"github.com/flex1988/go-level-logger"
	"io/ioutil"
	"os"
	"os/user"
)

type Configuration struct {
	ACCESS_KEY string
	SECRET_KEY string
}

var (
	config Configuration
)

func main() {
	logger, _ := logger.New()
	for _, v := range os.Args {
		logger.Info(v)
	}
	usr, _ := user.Current()

	b, _ := ioutil.ReadFile(usr.HomeDir + "/.imagerc")
	json.Unmarshal(b, &config)
	logger.Info(config.ACCESS_KEY)
}
