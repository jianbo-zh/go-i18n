package goi18n

import (
	"encoding/json"
	"io"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// *message.Printer
var p = message.NewPrinter(language.AmericanEnglish)

type TmplRaw struct {
	Format string        `json:"f"`
	Params []interface{} `json:"p"`
}

func Setup(defaultLang language.Tag) {
	if defaultLang != language.AmericanEnglish {
		p = message.NewPrinter(defaultLang)
	}
}

// Fprintf is like fmt.Fprintf, but using language-specific formatting.
func Fprintf(w io.Writer, key message.Reference, a ...interface{}) (n int, err error) {
	return p.Fprintf(w, key, a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...interface{}) string {
	return p.Sprintf(format, a...)
}

// Printf is like fmt.Printf, but using language-specific formatting.
func Printf(format string, a ...interface{}) {
	_, _ = p.Printf(format, a...)
}

func Templatef(format string, a ...interface{}) string {
	bs, _ := json.Marshal(TmplRaw{Format: format, Params: a})
	return string(bs)
}

func ParseTemplate(lang language.Tag, tmplStr string) string {
	var tmplF TmplRaw
	json.Unmarshal([]byte(tmplStr), &tmplF)
	return message.NewPrinter(lang).Sprintf(tmplF.Format, tmplF.Params...)
}
