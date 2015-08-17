/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse

import (
	//"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func Scan(src string, path string) []string {
	var result []string

	fset := token.NewFileSet()
	var f *ast.File
	var err error
	if src == "" {
		f, err = parser.ParseFile(fset, path, nil, parser.ParseComments)
	} else {
		f, err = parser.ParseFile(fset, path, src, parser.ParseComments)
	}
	if err != nil {
		panic(err)
	}
	//ast.Print(fset, f)
	for _, decl := range f.Decls {
		//fmt.Println(index)
		//spew.Dump(decl)
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Doc == nil {
				//fmt.Println("genDecl.Doc == nil")
			} else {
				//spew.Dump(genDecl.Doc)
				for _, comment := range genDecl.Doc.List {
					if strings.Contains(comment.Text, "@ginger") {
						for _, spec := range genDecl.Specs {
							//spew.Dump(genDecl.Specs)
							if typeSpec, ok := spec.(*ast.TypeSpec); ok {
								if typeSpec.Name != nil {
									//fmt.Println(typeSpec.Name.Name)
									result = append(result, typeSpec.Name.Name)
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}
