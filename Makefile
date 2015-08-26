.PHONY: all clean

all: main.go
	go build
	./ginger

clean:
	@go clean
	rm -f gen/web_service.go
	rm -f gen/main.go
	rm -f gen/config.yaml
	rm -f gen/*_resource.go
	rm -f gen/Makefile
	rm -f web_service.go
	rm -f *_resource.go

test:
	@go test
