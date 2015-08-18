.PHONY: all clean

all: main.go
	go build
	./ginger

clean:
	@go clean
	rm -f gen/web_service.go
	rm -f gen/slack_resource.go
	rm -f web_service.go
	rm -f person_resource.go

test:
	@go test
