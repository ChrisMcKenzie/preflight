package config

import (
	fmt "fmt"
	"io/ioutil"
	"testing"

	"github.com/hashicorp/hcl"
)

func readFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func TestParse(t *testing.T) {
	fName := "../examples/test.pf"
	data, err := readFile(fName)
	if err != nil {
		t.Error(err)
	}

	f, err := hcl.Parse(data)
	if err != nil {
		t.Error(err)
	}

	cfg, err := parse(fName, f)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v\n", cfg)
}
