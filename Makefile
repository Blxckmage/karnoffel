.PHONY: build clean help

build:
	go mod tidy && go build -o bin/game ./cmd/main.go

clean:
	rm -rf bin/
