NAME=$(shell basename $(CURDIR))

.PHONY: all
all: debug release

.PHONY: debug
debug: build/debug/$(NAME)

.PHONY: release
release: build/release/$(NAME)_linux_amd64

.PHONY: go-generate-all
go-generate-all:
	go generate -v ./...

build/debug/$(NAME): go-generate-all
	go build -o $@ -race

build/release/$(NAME)_linux_amd64: go-generate-all
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@

.PHONY: test
test:
	go test -v -race ./...

.PHONY: lint
lint:
	gometalinter -j $$(( $$(nproc) * 2 )) --deadline=90s ./...

.PHONY: dep-ensure
dep-ensure:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	dep ensure -v -vendor-only
