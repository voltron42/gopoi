package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"

	"fmt"

	"./model"
	"./translate"
)

func main() {
	fmt.Println("reading config")
	bytes, err := ioutil.ReadFile("./config.json")
	check(err)
	conf := Config{}
	fmt.Println("unmarshalling config")
	err = json.Unmarshal(bytes, &conf)
	check(err)
	fmt.Println("reading template")
	bytes, err = ioutil.ReadFile(conf.In)
	check(err)
	doc := model.Document{}
	fmt.Println("unmarshalling template")
	err = xml.Unmarshal(bytes, &doc)
	check(err)
	translator := translate.New()
	fmt.Println("translating document")
	err = translator.Translate(doc, conf.Out)
	check(err)
	fmt.Println("translation complete")
}

type Config struct {
	In  string `json:"in"`
	Out string `json:"out"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
