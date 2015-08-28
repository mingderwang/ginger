/*
 * Copyright 2015 Ming-der Wang<ming@log4analytics.com> All right reserved.
 * Licensed by MIT License
 */
package gen

import (
	//"fmt"
	"os"
	"text/template"

	log "github.com/Sirupsen/logrus"
)

var (
	gingerTemplateMakefile = template.Must(template.New("ginger").Parse(
		`{{range .Types}}
.PHONY: all clean

all: onion.go 
	@go build
	./onion migratedb
	./onion serve

clean:
	@go clean
	rm -f *_resource.go
	rm -f main.go
	rm -f web_service.go
	rm -f config.yaml
	rm -f Makefile
	rm -f Dockerfile
	rm -f dockerize.sh
	rm -f -r Godeps

cleandb:
	rm -r *.db

migrate:
	./onion migratedb

run:
	./onion serve
{{end}}
`))
)

func init() {
	log.SetFormatter(&log.TextFormatter{}) // or JsonFormatter
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
}

func GenMakefile(path string) {
	types := findTypes(path)
	output, err := os.OpenFile("Makefile", os.O_WRONLY|os.O_CREATE, 0600)
	defer output.Close()
	checkError(err, "could not open output file")
	gingerTemplateMakefile.Execute(output, AllType{types})
}
