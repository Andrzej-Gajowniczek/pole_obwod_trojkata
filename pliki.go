package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var mydir string = "/home/alfa/Documents/"

func main() {
	files, err := ioutil.ReadDir(mydir)
	if err != nil {
		log.Fatal(err)
	}
	var xx string = ""
	for _, file := range files {
		if file.IsDir() {
			xx = "/"
		} else {
			xx = ""
		}
		fmt.Println(mydir + file.Name() + xx)
	}
}
