.PHONY: all clean

all: main.go
	go build
	./ginger

clean:
	@go clean

test:
	@go test
