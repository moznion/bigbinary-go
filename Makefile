PKGS := $(shell go list ./...)

check: test lint vet fmt-check

test:
	go test -v $(PKGS)

lint:
	golint -set_exit_status $(PKGS)

vet:
	go vet $(PKGS)

fmt-check:
	gofmt -l -s *.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi; \
	goimports -l *.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi \

bootstrap:
	go get -u golang.org/x/lint/golint golang.org/x/tools/cmd/goimports

