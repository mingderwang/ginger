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
`{{range .Types}}
FROM debian:wheezy

RUN apt-get update -y && apt-get install --no-install-recommends -y -q \
                         curl \
                         build-essential \
                         ca-certificates \
                         git mercurial bzr \
               && rm -rf /var/lib/apt/lists/*

ENV GOVERSION 1.5
RUN mkdir /goroot && mkdir /gopath
RUN curl https://storage.googleapis.com/golang/go${GOVERSION}.linux-amd64.tar.gz \
           | tar xvzf - -C /goroot --strip-components=1

ENV GOPATH /gopath
ENV GOROOT /goroot
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

ENV foo /gopath/src/onion
WORKDIR ${foo}  
ADD . $foo
RUN go get github.com/mingderwang/ginger
RUN go get github.com/tools/godep
RUN godep restore
RUN go install
EXPOSE 8080

{{end}}
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
