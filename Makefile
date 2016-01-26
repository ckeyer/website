PKG := github.com/ckeyer/website
APP := website
GODEP := godep
GO := $(GODEP) go 
VERSION := $(shell cat VERSION.txt)
LD_FLAGS := -X $(PKG)/version.version=$(VERSION)

default: build

build: 
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix nocgo -ldflags="$(LD_FLAGS)" -o bin/$(APP)

test: 
	$(GO) test -ldflags="$(LD_FLAGS)" ./...
