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
`{{range .Types}}
#!/bin/bash
set -o xtrace
echo "start dockerizing ..."
go generate
go build
godep save -r
{{end}}
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
