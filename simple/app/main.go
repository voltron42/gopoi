package main

import (
	"../"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./sample.xml")
	if err != nil {
		panic(err)
	}
	d := simple.Document{}
	err = xml.Unmarshal(data, &d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("document: %v\n", d)
	err = d.Publish()
	if err != nil {
		panic(err)
	}
}
