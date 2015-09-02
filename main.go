/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package main

import (
	//	"fmt"
	"github.com/mingderwang/ginger/gen"
	"github.com/mingderwang/ginger/parse"
	"os"
)

// don't print
type Person2 struct {
	FirstName string
	LastName  string
	HairColor string
}

// @ginger test
type Person struct {
	FirstName string
	LastName  string
	HairColor string
}

func main() {
	if len(os.Args) > 1 {
		var types []string
		fileName := os.Args[1]
		if fileName == "" {
			return
		}
		types = parse.Scan("", fileName)
		for _, typeName := range types {
			gen.GenWebService(typeName)
			gen.GenResourceFile(typeName)
			gen.GenMain(typeName)
			//			gen.GenConfig(typeName)
			gen.GenMakefile(typeName)
			gen.GenDockerfile(typeName)
			gen.GenShellCode(typeName)
		}
	}
}
