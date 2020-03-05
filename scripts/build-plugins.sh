#!/bin/bash

ROOT=$(pwd)
echo "# Building Plugins..."
# Generate various test protos.

for f in plugin/plugins/*; do
    if [ -d "$f" ]; then
        # $f is a directory
        cd $f/plugin
        name=$(basename $f)
        echo "## Building $name"
        version=$(cat VERSION)
        echo version=$version
        set -x
        go build "-ldflags='-X main.Version=$version'" -buildmode=plugin -o ../../${name}_plugin.so
        { set +x; } 2>/dev/null
    fi
done
