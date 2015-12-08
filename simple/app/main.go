package main

import (
	"../"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	files := []string{}
	err = json.Unmarshal(bytes, &files)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		d := simple.Document{}
		err = xml.Unmarshal(data, &d)
		if err != nil {
			panic(err)
		}
		err = d.Publish()
		if err != nil {
			panic(err)
		}
	}
}
