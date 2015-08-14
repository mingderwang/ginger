.PHONY: all clean

all: main.go
	go build
	./go-template

clean:
	@go clean

test:
	@go test
