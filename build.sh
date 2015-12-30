#!/bin/bash

set -e

echo "building for OSX..."
GOOS=darwin go build
mv catd bin/catd_mac
echo "done."

echo "building for Linux..."
GOOS=linux go build
mv catd bin/catd_linux
echo "done."
