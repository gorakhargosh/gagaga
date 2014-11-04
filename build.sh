#!/bin/sh

for f in *.go; do
  echo "go build $f";
  go build $f;
done;
