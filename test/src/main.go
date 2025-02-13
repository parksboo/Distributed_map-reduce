package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("test1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
