
Install:
```bash
# install extract tool
go install github.com/jianbo-zh/go-i18n/cmd/i18n@latest

# extract translation content
i18n -extract.dir=./cmd/

# import i18n package
import ("github.com/jianbo-zh/go-i18n")

# project init
goi18n.Setup(language.English, []language.Tag{language.English}, "path/to/translations")
```

Options:
---------------------------
-default.language string
    default language (default "en")
-extract.dir string
    The directory containing the files to be translated (default ".")
-output.dir string
    output directory for translation strings (default "translations")
-packname string
    pack name (default "goi18n")
-support.languages string
    all support languages (default "en,zh")
---------------------------

Description:
1. extract param string within diration `-extract.dir` packname (`goi18n`) method (`Fprintf`, `Printf`, `Sprintf`, `Templatef`)
2. generate translations files to diration through by `-output.dir` option