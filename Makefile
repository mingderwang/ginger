.PHONY: all clean

all: main.go
	go build
	./ginger

clean:
	@go clean
	rm -f slack.go

test:
	@go test
