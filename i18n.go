package goi18n

import (
	"encoding/json"
	"io"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var defaultLang = language.AmericanEnglish

var langPrinters = make(map[language.Tag]*message.Printer, 0)

func init() {
	langPrinters[defaultLang] = message.NewPrinter(defaultLang)
}

type TmplRaw struct {
	Format string        `json:"f"`
	Params []interface{} `json:"p"`
}

func Setup(lang language.Tag, supports ...language.Tag) {

	supports = append(supports, lang)
	for _, val := range supports {
		if _, exists := langPrinters[val]; !exists {
			langPrinters[val] = message.NewPrinter(val)
		}
	}

	if lang != defaultLang {
		defaultLang = lang
	}
}

// Fprintf is like fmt.Fprintf, but using language-specific formatting.
func Fprintf(w io.Writer, key message.Reference, a ...interface{}) (n int, err error) {
	return langPrinters[defaultLang].Fprintf(w, key, a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...interface{}) string {
	return langPrinters[defaultLang].Sprintf(format, a...)
}

// Printf is like fmt.Printf, but using language-specific formatting.
func Printf(format string, a ...interface{}) {
	langPrinters[defaultLang].Printf(format, a...)
}

// Fprintfl is like fmt.Fprintf, but using language-specific formatting.
func Fprintfl(lang language.Tag, w io.Writer, key message.Reference, a ...interface{}) (n int, err error) {
	if _, exists := langPrinters[lang]; exists {
		return langPrinters[lang].Fprintf(w, key, a...)
	}
	return langPrinters[defaultLang].Fprintf(w, key, a...)
}

// Sprintfl formats according to a format specifier and returns the resulting string.
func Sprintfl(lang language.Tag, format string, a ...interface{}) string {
	if _, exists := langPrinters[lang]; exists {
		return langPrinters[lang].Sprintf(format, a...)
	}
	return langPrinters[defaultLang].Sprintf(format, a...)
}

// Printfl is like fmt.Printf, but using language-specific formatting.
func Printfl(lang language.Tag, format string, a ...interface{}) {
	if _, exists := langPrinters[lang]; exists {
		langPrinters[lang].Printf(format, a...)
		return
	}
	langPrinters[defaultLang].Printf(format, a...)
}

func Templatef(format string, a ...interface{}) string {
	bs, _ := json.Marshal(TmplRaw{Format: format, Params: a})
	return string(bs)
}

func ParseTemplate(lang language.Tag, tmplStr string) string {
	var tmplF TmplRaw
	json.Unmarshal([]byte(tmplStr), &tmplF)

	if _, exists := langPrinters[lang]; exists {
		return langPrinters[lang].Sprintf(tmplF.Format, tmplF.Params...)
	}

	return langPrinters[defaultLang].Sprintf(tmplF.Format, tmplF.Params...)
}
