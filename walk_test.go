package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestWalk(t *testing.T) {

	root := "/Users/daveappleton/Documents/"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	t.Fail()
}
