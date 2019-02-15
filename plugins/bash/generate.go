package bash

//go:generate echo "# Building plugin \"bash\"..."
//go:generate go build "-ldflags='-X main.Version=0.0.1'" -buildmode=plugin -o ../bash_plugin.so ./plugin
