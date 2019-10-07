package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func initializeConfig(filename string) connCreds {
	var config connCreds
	source, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)

	if err != nil {
		panic(err)
	}

	return config
}
