package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Badges struct {
	Badges []Config
}

type Config struct {
	Name   string
	Type   string
	Logic  string
	Output string
	Values map[string]string
}

func getConfig(file string) []Config {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	expandedData := os.ExpandEnv(string(data))

	b := Badges{}

	err = yaml.Unmarshal([]byte(expandedData), &b)
	if err != nil {
		panic(err)
	}

	return b.Badges

}
