check_deps:
	$(eval GLIDE:=$(shell which glide))
	if [ -n "$(GLIDE)" -a -x "$(GLIDE)" ]; then echo "check_deps: OK"; else echo "Please install glide c.f. TODO"; fi;

setup: check_deps
	go get honnef.co/go/tools/cmd/staticcheck
	go get github.com/golang/lint/golint

bundle_install: setup
	glide install

bundle_update: setup
	glide update

lint: setup
	staticcheck $$(glide novendor)
	golint $$(glide novendor)
	go fmt $$(glide novendor)

build_darwin: lint
	GOOS=darwin GOARCH=amd64 go build main.go

build_windows: lint
	GOOS=windows GOARCH=amd64 go build

build_linux: lint
	GOOS=linux GOARCH=amd64 go build

.PHONY: setup bundle_install bundle_update lint build_darwin build_windows
