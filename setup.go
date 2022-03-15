package goi18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func init() {
	p = message.NewPrinter(language.AmericanEnglish)
}

func Setup(defaultLang language.Tag, supportLangs []language.Tag, translationDir string) error {

	if defaultLang != language.AmericanEnglish {
		p = message.NewPrinter(defaultLang)
	}

	fi, err := ioutil.ReadDir(translationDir)
	if err != nil {
		log.Fatal(err)
	}

	defaultLangName := defaultLang.String()

	for _, v := range fi {
		if !v.IsDir() {
			continue
		}

		curLang := defaultLang
		if v.Name() != defaultLangName {
			isSupport := false
			for _, lang := range supportLangs {
				if v.Name() == lang.String() {
					curLang = lang
					isSupport = true
					break
				}
			}

			if !isSupport {
				continue
			}
		}

		translationFile := path.Join(translationDir, v.Name(), "data.json")
		if _, err := os.Stat(translationFile); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return fmt.Errorf("get translation file [%s] error: %w", translationFile, err)
		}

		content, err := ioutil.ReadFile(translationFile)
		if err != nil {
			log.Fatal(err)
		}

		var trans map[string]string
		err = json.Unmarshal(content, &trans)
		if err != nil {
			log.Fatal(err)
		}

		for key, msg := range trans {
			_ = message.SetString(curLang, key, msg)
		}
	}

	return nil
}
