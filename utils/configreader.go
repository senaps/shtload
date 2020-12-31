package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var CONFIG_PATH = os.Getenv("CONFIG_PATH")

func readFile(path string) string {
	filedata, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(filedata)
}

func ReadConfig() Config {
	confFile := readFile(CONFIG_PATH)
	var conf Config
	err := yaml.Unmarshal([]byte(confFile), &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
