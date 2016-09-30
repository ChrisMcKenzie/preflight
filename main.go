package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ChrisMcKenzie/preflight/preflight"
	"github.com/fatih/color"
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

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	for _, task := range cl.Tasks {
		b, _ := json.MarshalIndent(task.Config, "", "  ")
		fmt.Printf("===== TASK: (%s) %s =====\n\n%s\n\n", green(task.Type), task.Name, yellow(string(b)))
	}
	// cl.Resolve()
}
