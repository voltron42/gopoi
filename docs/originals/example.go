package example

import (
	"fmt"
)

func Summary(err error, txt string) {
	if err != nil {
		panic(err)
	}
	fmt.Println(txt)
}

func Filename(txt string) string {
	return txt + ".pdf"
}

func FontDir() string {
	return ""
}
