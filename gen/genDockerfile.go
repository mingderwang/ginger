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
	gingerTemplateDockerfile = template.Must(template.New("ginger").Parse(
		`FROM asia.gcr.io/winter-wonder-647/base-onion:latest 
CMD ["/go/bin/onion","serve"]
EXPOSE 8080
`))
)

func init() {
	log.SetFormatter(&log.TextFormatter{}) // or JsonFormatter
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
}

func GenDockerfile(path string) {
	types := findTypes(path)
	output, err := os.OpenFile("Dockerfile", os.O_WRONLY|os.O_CREATE, 0600)
	defer output.Close()
	checkError(err, "could not open output file")
	gingerTemplateDockerfile.Execute(output, AllType{types})
}
