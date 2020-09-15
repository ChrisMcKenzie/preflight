#!/bin/bash

ROOT=$(pwd)
PLUGIN_PATH="plugin/plugins"
BUILD_PATH="build/plugins"

echo "# Building Plugins..."

mkdir -p $BUILD_PATH

for f in $PLUGIN_PATH/*; do
    if [ -d "$f" ]; then
        # $f is a directory
        name=$(basename $f)
        echo "## Building $name"
        version=$(cat $f/plugin/VERSION)
        echo version=${version:-"0.0.0"}
        GO_VERSION=$(go version | awk '{print $3}')
        echo goversion=$GO_VERSION
        set -x
        go build -ldflags="-X main.Version=$version" -gcflags="-trimpath=$GOPATH/src" -trimpath -buildmode=plugin -o $BUILD_PATH/${name}_${GO_VERSION}_plugin.so  ./$f/plugin
        { set +x; } 2>/dev/null
    fi
done
