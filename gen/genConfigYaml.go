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
	gingerTemplateConfig = template.Must(template.New("ginger").Parse(
		`{{range .Types}}
svchost: :8080
dbuser: root
dbpassword: ""
dbhost: localhost
dbname: {{.TypeName}}.db
{{end}}
`))
)

func init() {
	log.SetFormatter(&log.TextFormatter{}) // or JsonFormatter
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
}

func GenConfig(path string) {
	types := findTypes(path)
	output, err := os.OpenFile("config.yaml", os.O_WRONLY|os.O_CREATE, 0600)
	defer output.Close()
	checkError(err, "could not open output file")
	gingerTemplateConfig.Execute(output, AllType{types})
}
