#!/bin/sh

go get
go generate
if [ $? -ne 0 ]; then
	echo "Go generate failed" >&2
	exit 1
fi

if [ -f countdown-server ]; then
	rm countdown-server
fi

os="${GOOS:-}"
arch="${GOARCH:-}"


GOOS="${1:-$os}" GOARCH="${2:-$arch}" go build -ldflags="-s -w" -v main.go data_generated.go
if [ $? -ne 0 ]; then
	echo "Go build failed" >&2
	exit 1
fi

mv main countdown-server