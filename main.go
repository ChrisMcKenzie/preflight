package main

import (
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

	var checklist preflight.CheckList
	err = hcl.Unmarshal(bytes, &checklist)
	if err != nil {
		log.Println(err)
		return
	}

	checklist.Resolve()
}
