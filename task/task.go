package task

import (
	fmt "fmt"
	"net/url"
	"os"
	"path"
	"plugin"
	"strings"
	"sync"

	"github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	getter "github.com/hashicorp/go-getter"
)

const (
	RegisterTaskFuncName = "RegisterTasks"

	TaskPackageIsVersion1 = true
	TaskPackageVersion    = "0.0.1"

	taskURL = "plugins.pre-flight.io"
)

type Task interface {
	Evaluate()
	Apply()
}

type TaskFunc func(Meta) Task

var (
	mu           sync.Mutex
	taskRegistry = make(map[string]TaskFunc)
)

func RegisterTask(name string, task TaskFunc) {
	mu.Lock()
	defer mu.Unlock()
	taskRegistry[name] = task
}

func GetTask(meta Meta) (Task, error) {
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

func getRemotePlugin(name string) Task {
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

	url := path.Join(taskURL, proto.MessageName(pb))
	return &types.Any{TypeUrl: url, Value: value}, nil
}

func LoadPlugin(path string) error {
	plug, err := plugin.Open(path)
	if err != nil {
		return err
	}
	fi, err := plug.Lookup(RegisterTaskFuncName)
	if err != nil {
		return err
	}

	register, ok := fi.(func() map[string]TaskFunc)
	if !ok {
		return err
	}

	tasks := register()
	for name, task := range tasks {
		RegisterTask(name, task)
	}

	return nil
}
