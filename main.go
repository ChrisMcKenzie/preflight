package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ChrisMcKenzie/preflight/preflight"
	"github.com/hashicorp/hcl"
)

func main() {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
	}

	file, err := hcl.ParseBytes(bytes)
	if err != nil {
		log.Println(err)
		return
	}

	cl, err := preflight.LoadHcl(file)
	if err != nil {
		fmt.Println(err)
	}
	cl.Resolve()
}
