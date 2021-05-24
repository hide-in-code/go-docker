package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"tda/docker"
)

func main() {
	yamlFile, err := ioutil.ReadFile("docker-compose.yml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	conf := new(docker.Yml)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(conf.Version)
	fmt.Println(conf.Services["db"].Environment)
}
