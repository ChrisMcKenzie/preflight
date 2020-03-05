package registry

import (
	fmt "fmt"
	"net/url"
	"os"
	"path"
	"plugin"
	"strings"
	"sync"

	plug "github.com/ChrisMcKenzie/preflight/plugin"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	getter "github.com/hashicorp/go-getter"
)

const (
	RegisterTaskFuncName = "RegisterTasks"

	TaskPackageVersion = "0.0.1"

	defaultTaskURL = "plugins.pre-flight.io"
)

var (
	mu           sync.Mutex
	taskRegistry = make(map[string]plug.TaskFunc)
)

func RegisterTask(name string, task plug.TaskFunc) {
	mu.Lock()
	defer mu.Unlock()
	taskRegistry[name] = task
}

func GetTask(meta plug.Meta) (plug.Task, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	in := meta.URL
	if meta.Version != "" {
		in = fmt.Sprintf("%s?version=%s", in, meta.Version)
	}

	d := []getter.Detector{&PluginRegistryDetector{}, &getter.FileDetector{}}
	rawURL, err := getter.Detect(in, dir, d)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	dPath := path.Join(".preflight/plugins", url.Path)
	if err := getter.GetFile(dPath, rawURL); err != nil {
		return nil, err
	}

	if err := LoadPlugin(dPath); err != nil {
		return nil, err
	}

	f := taskRegistry[meta.URL]
	return f(meta), nil
}

func getRemotePlugin(name string) plug.Task {
	return nil
}

func GetTypeName(task *types.Any) string {
	return strings.TrimSuffix(strings.TrimPrefix(task.GetTypeUrl(), "type.googleapis.com/"), ".Task")
}

func MarshalAny(pb proto.Message) (*types.Any, error) {
	value, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}

	url := path.Join(defaultTaskURL, proto.MessageName(pb))
	return &types.Any{TypeUrl: url, Value: value}, nil
}

func LoadPlugin(path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return err
	}

	fi, err := p.Lookup(RegisterTaskFuncName)
	if err != nil {
		return err
	}

	register, ok := fi.(func() map[string]plug.TaskFunc)
	if !ok {
		return err
	}

	tasks := register()
	for name, task := range tasks {
		RegisterTask(name, task)
	}

	return nil
}
