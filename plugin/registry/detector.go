package registry

import (
	fmt "fmt"
	"net/url"
	"path"
	"strings"
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
	if dir == "" {
		dir = "library"
	}
	pparts := strings.Split(plugin, ".")
	pName := pparts[0]
	soName := fmt.Sprintf("%s_plugin.so", pName)

	var repoURL url.URL
	repoURL.Scheme = "https"
	repoURL.Host = "plugins.pre-flight.io"
	repoURL.Path = path.Join(dir, pName, soName)
	repoURL.RawQuery = u.Query().Encode()

	return repoURL.String(), true, nil
}
