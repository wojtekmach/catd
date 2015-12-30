#!/bin/bash

set -e

echo "building for OSX..."
GOOS=darwin go build
ls
mv catd bin/catd_mac
echo "done."

echo "building for Linux..."
GOOS=linux go build
ls
mv catd bin/catd_linux
echo "done."
