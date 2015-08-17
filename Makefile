.PHONY: all clean

all: main.go
	go build
	./ginger

clean:
	@go clean
	rm -f gen/web_service.go

test:
	@go test
