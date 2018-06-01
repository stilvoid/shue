#!/bin/bash

GOOS=linux go build -o shue-linux
GOOS=darwin go build -o shue-osx
GOOS=windows go build -o shue-windows.exe
