package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type Trans map[string]map[string]string

func generator(config CommandConfig) {

	trans := Trans{}
	for _, lang := range strings.Split(config.SupportLangs, ",") {
		filename := path.Join(config.OutputDir, lang, "data.json")
		_, err := os.Stat(filename)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			} else {
				log.Fatal(err)
			}
		}

		var keyvals map[string]string
		bsData, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(bsData, &keyvals)
		if err != nil {
			log.Fatal(err)
		}

		for key, val := range keyvals {
			if trans[key] == nil {
				trans[key] = map[string]string{}
			}

			trans[key][lang] = val
		}
	}

	goFile, err := os.OpenFile(path.Join(config.OutputDir, "init.go"), os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer goFile.Close()

	goFile.Truncate(0)

	transJson, _ := json.MarshalIndent(trans, "", "  ")
	err = i18nTmpl.Execute(goFile, struct{ Trans string }{string(transJson)})
	if err != nil {
		log.Fatal(err)
	}
}

var funcs = template.FuncMap{
	"funcName": func(lang string) string {
		lang = strings.ReplaceAll(lang, "_", "")
		lang = strings.ToUpper(lang[:1]) + lang[1:]
		return lang
	},
}

var i18nTmpl = template.Must(template.New("i18n").Funcs(funcs).Parse(`package translations

import (
	"encoding/json"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var transJson = ` + "`{{ .Trans }}`" + `

type Trans map[string]map[string]string

func init() {
	var trans Trans
	err := json.Unmarshal([]byte(transJson), &trans)
	if err != nil {
		panic(err)
	}

	var langTag language.Tag
	for key, items := range trans {
		for lang, msg := range items {
			err = langTag.UnmarshalText([]byte(lang))
			if err != nil {
				panic(err)
			}
			err = message.SetString(langTag, key, msg)
			if err != nil {
				panic(err)
			}
		}
	}
}
`))
