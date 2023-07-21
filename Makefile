## build/api: build app binary
.PHONY: build
build:
	@echo 'Building...'
	GOOS=windows GOARCH=amd64 go build -ldflags='-s' -o=./tmp/glc.exe ./
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./tmp/glc ./