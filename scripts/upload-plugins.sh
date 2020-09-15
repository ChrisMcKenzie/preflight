#!/bin/bash

ROOT=$(pwd)
PLUGIN_PATH="plugin/plugins"
BUILD_PATH="build/plugins"

echo "# Uploading Plugins..."

for f in $BUILD_PATH/*; do
  echo $f
  # $f is a directory
  name=$(basename $f)
  plugin_name=$(echo $name | awk -F '_' '{print $1}')
  echo $plugin_name

  echo "## Uploading $name"
  GO_VERSION=$(go version | awk '{print $3}')
  echo $GO_VERSION
  set -x
  go run -trimpath scripts/upload_plugins.go $f library/$plugin_name/$name
  { set +x; } 2>/dev/null
done
