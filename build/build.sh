#!/bin/bash

[ -d .preflight ] || mkdir .preflight

for d in `find ./builtin/bin/* -type d`; do 
  echo Building \".preflight/provisioner-$(basename $d)\"...; 
  go build -o .preflight/provisioner-$(basename $d)
done
