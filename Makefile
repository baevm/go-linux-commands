## build/api: build app binary
.PHONY: build
build:
	@echo 'Building...'
	GOOS=windows GOARCH=amd64 go build -tags "windows" -ldflags="-s -w" -o=./tmp/glc.exe ./
	GOOS=linux GOARCH=amd64 go build -tags "linux" -ldflags="-s -w" -o=./tmp/glc ./

.PHONY: install
install:
	go install -ldflags '-X "main.version=dev build $(dev_build_version)"' ./...