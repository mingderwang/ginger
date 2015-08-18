/*
 * Copyright 2015 Ming-der Wang<ming@log4analytics.com> All right reserved.
 * Licensed by MIT License
 */
package gen

import (
	//"fmt"
	"os"
	"unicode"

	log "github.com/Sirupsen/logrus"
)

type GenType struct {
	TypeName     string // could be "Slack"
	VariableName string // could be "slack"
}

type AllType struct {
	Types []GenType
}

func upperFirstChar(in string) string {
	if in == "" {
		return ""
	}
	a := []rune(in)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

func lowerFirstChar(in string) string {
	if in == "" {
		return ""
	}
	a := []rune(in)
	a[0] = unicode.ToLower(a[0])
	return string(a)
}

func init() {
	log.SetFormatter(&log.TextFormatter{}) // or JsonFormatter
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
}

func checkError(err error, message string) {
	if err != nil {
		log.Fatal(message)
		panic(err)
	}
}

func findTypes(name string) []GenType {
	var types []GenType
	item := GenType{upperFirstChar(name), lowerFirstChar(name)}
	types = append(types, item)
	return types
}
