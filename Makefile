.PHONY: all clean

all: main.go
	go build
	./ginger

clean:
	@go clean
	rm -f slack.go
	rm -f web_service.go

test:
	@go test
