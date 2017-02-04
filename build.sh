#!/bin/sh

if [ $# != 1 ]; then
	echo "Usage: $0 [binary name]"
	exit 0
fi

mkdir -p bin/zip

GOOS=linux GOARCH=amd64 go build -o ./bin/linux64/$1
zip ./bin/zip/linux64.zip ./bin/linux64/$1
GOOS=linux GOARCH=386 go build -o ./bin/linux386/$1
zip ./bin/zip/linux386.zip ./bin/linux386/$1

GOOS=windows GOARCH=386 go build -o ./bin/windows386/$1.exe
zip ./bin/zip/windows386.zip ./bin/windows386/$1.exe
GOOS=windows GOARCH=amd64 go build -o ./bin/windows64/$1.exe
zip ./bin/zip/windows64.zip ./bin/windows64/$1.exe

GOOS=darwin GOARCH=386 go build -o ./bin/darwin386/$1
zip ./bin/zip/darwin386.zip ./bin/darwin386/$1
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin64/$1
zip ./bin/zip/darwin64.zip ./bin/darwin64/$1
