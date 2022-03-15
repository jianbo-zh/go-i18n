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

	files, err := glob(config.ExtractDir, ".go")
	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range files {

		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, filename, string(content), 0)
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

	defaultLangDir := path.Join(config.OutputDir, config.DefaultLang)
	fi, err := os.Stat(defaultLangDir)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(defaultLangDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else if !fi.IsDir() {
		log.Fatalf("[%s] is not dir", defaultLangDir)
	}

	defaultOutputFile := path.Join(defaultLangDir, "data.json")
	err = ioutil.WriteFile(defaultOutputFile, content, 0664)
	if err != nil {
		log.Fatal(err)
	}

	for _, lang := range strings.Split(config.SupportLangs, ",") {
		if lang == config.DefaultLang || lang == "" {
			continue
		}

		langDir := path.Join(config.OutputDir, lang)
		fi, err := os.Stat(langDir)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(langDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		} else if !fi.IsDir() {
			log.Fatalf("[%s] is not dir", langDir)
		}

		langFile := path.Join(langDir, "data.json")
		if _, err := os.Stat(langFile); errors.Is(err, os.ErrNotExist) {
			err = ioutil.WriteFile(langFile, content, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
