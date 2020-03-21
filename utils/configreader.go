package utils

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var CONFIG_PATH = os.Getenv("CONFIG_PATH")

func readFile(path string) []byte {
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return conf
}

func ReadConfig() Config {
	confFile := readFile(CONFIG_PATH)
	conf := Config{}
	err := yaml.Unmarshal(confFile, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
