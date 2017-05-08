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

.PHONY: setup bundle_install bundle_update lint
