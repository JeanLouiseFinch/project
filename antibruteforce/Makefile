fmt:
	gofmt -w -s -d .
vet:
	go vet ./...
lint:
	golint ./...
imports:
	goimports -w .
tidy:
	go mod tidy
verify:
	go mod verify
pre-commit: fmt vet lint tidy verify imports
.PHONY: build
build:
	go build -o build/service
