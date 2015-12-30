#!/bin/bash

set -e

GOOS=darwin go build
mv catd bin/catd_mac
shasum -a 256 bin/catd_mac

GOOS=linux go build
mv catd bin/catd_linux
shasum -a 256 bin/catd_linux
