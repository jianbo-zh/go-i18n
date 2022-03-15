package main

import (
	"encoding/json"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	config := ParseFlag()

	data := make(map[string]string)

	fi, err := ioutil.ReadDir(config.ExtractDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range fi {
		if v.IsDir() {
			continue
		}

		content, err := ioutil.ReadFile(v.Name())
		if err != nil {
			log.Fatal(err)
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, v.Name(), string(content), 0)
		if err != nil {
			log.Fatal(err)
		}

		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			fn, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			pack, ok := fn.X.(*ast.Ident)
			if !ok {
				return true
			}
			if pack.Name != config.Packname {
				return true
			}
			if len(call.Args) == 0 {
				return true
			}

			var expr ast.Expr
			// if Fprintf, we'll take second arg as template
			if fn.Sel.Name == "Fprintf" {
				expr = call.Args[1]
			} else { // include Printf, Sprintf
				expr = call.Args[0]
			}
			str, ok := expr.(*ast.BasicLit)
			if !ok {
				return true
			}

			// Keep this for later debug usage.
			// log.Printf("%v", str.Value)
			data[str.Value] = str.Value
			return true
		})
	}

	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// zh_CN
	defaultOutputFile := path.Join(config.OutputDir, config.DefaultLang, "data.json")
	err = ioutil.WriteFile(defaultOutputFile, content, 0664)
	if err != nil {
		log.Fatal(err)
	}

	for _, lang := range strings.Split(",", config.SupportLangs) {
		if lang == config.DefaultLang || lang == "" {
			continue
		}

		langFile := path.Join(config.OutputDir, lang, "data.json")
		if _, err := os.Stat(langFile); errors.Is(err, os.ErrNotExist) {
			err = ioutil.WriteFile(langFile, content, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
