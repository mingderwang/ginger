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
	gingerTemplateShellCode = template.Must(template.New("ginger").Parse(
		`#!/bin/bash
set -o xtrace
docker build -t ming/chilli .
docker run -d -p 80:8080 ming/chilli
`))
)

func init() {
	log.SetFormatter(&log.TextFormatter{}) // or JsonFormatter
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
}

func GenShellCode(path string) {
	types := findTypes(path)
	output, err := os.OpenFile("dockerize.sh", os.O_WRONLY|os.O_CREATE, 0700)
	defer output.Close()
	checkError(err, "could not open output file")
	gingerTemplateShellCode.Execute(output, AllType{types})
}
