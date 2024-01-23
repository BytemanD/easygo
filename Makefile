
UPX=$(shell which upx 2> /dev/null)
VERSION := $(shell go run cmd/easygo.go -v |awk '{{print $$3}}')
GO_VERSION := $(shell go version |awk '{{print $$3}}')
BUILD_DATE := $(shell date +'%Y-%m-%d %H:%M:%S')
UNAME := $(shell uname -sriv)

build:
	go mod download
	mkdir -p dist
	go build  -o dist/ -ldflags " \
		-X 'main.Version=$(VERSION)' \
		-X 'main.GoVersion=$(GO_VERSION)' \
		-X 'main.BuildDate=$(BUILD_DATE)' \
		-X 'main.BuildPlatform=$(UNAME)' -s -w" \
		cmd/easygo.go
	
ifeq ("$(UPX)", "")
	echo "upx not install"
else
	$(UPX) -q dist/easygo > /dev/null
endif
