package homebrew

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/ChrisMcKenzie/preflight/preflight"
)

type provisioner struct{}

type brewJSONInfo struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"desc"`
	Versions    struct {
		Stable *string `json:"stable"`
		Bottle bool    `json:"bottle"`
		Devel  *string `json:"devel"`
		Head   *string `json:"head"`
	} `json:"versions"`
	Installed []struct {
		Version    string   `json:"version"`
		Options    []string `json:"used_options"`
		FromBottle bool     `json:"poured_from_bottle"`
	} `json:"installed"`
}

// Provisioner ...
func Provisioner() preflight.Provisioner {
	return &provisioner{}
}

// Validate ...
func (*provisioner) Validate(t *preflight.Task) ([]string, []error) {
	return []string{}, []error{}
}

// Exists ...
func (*provisioner) Exists(t *preflight.Task) (bool, error) {
	var buf bytes.Buffer
	cmd := exec.Command("brew", "info", "--json=v1", t.Config["name"].(string))
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		return false, err
	}

	result := []brewJSONInfo{}
	d := json.NewDecoder(&buf)
	err = d.Decode(&result)
	if err != nil {
		return false, err
	}

	if len(result[0].Installed) > 0 {
		return true, nil
	}

	return false, nil
}
