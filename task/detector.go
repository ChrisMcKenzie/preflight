package task

import (
	fmt "fmt"
	"net/url"
	"os"
	"path"
	"strings"

	getter "github.com/hashicorp/go-getter"
)

const (
	BuiltinPluginDomain = "github.com/ChrisMcKenzie/preflight"
)

type PluginRegistryDetector struct{}

func (pd *PluginRegistryDetector) Detect(src, pwd string) (string, bool, error) {
	u, err := url.Parse(src)
	if err != nil {
		return "", false, err
	}

	dir, plugin := path.Split(u.String())
	if dir != "" {
		pparts := strings.Split(plugin, ".")
		pName := pparts[0]

		var repoURL url.URL
		repoURL.Scheme = "https"
		repoURL.Host = "plugins.pre-flight.io"
		repoURL.Path = path.Join(dir, pName)
		repoURL.RawQuery = u.Query().Encode()

		return repoURL.String(), true, nil
	}

	pparts := strings.Split(plugin, ".")
	pName := fmt.Sprintf("%s_plugin.so", pparts[0])
	d := &getter.FileDetector{}
	return d.Detect(path.Join(os.Getenv("GOPATH"), "src", BuiltinPluginDomain, "plugins", pName), pwd)
}
