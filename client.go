package main

import (
	"encoding/json"
	"github.com/flex1988/go-level-logger"
	"io/ioutil"
	"os/user"
)

type Configuration struct {
	ACCESS_KEY string
	SECRET_KEY string
}

var (
	config Configuration
	log    *logger.Logger
)

func loadConfig(path string) error {
	b, _ := ioutil.ReadFile(path)
	return json.Unmarshal(b, &config)
}

func init() {
	log, _ = logger.New()
	usr, _ := user.Current()
	err := loadConfig(usr.HomeDir + "/.imagerc")
	if err != nil {
		log.Error("没有发现配置文件")
	}
}

func main() {
	log.Info(config.ACCESS_KEY)
}
