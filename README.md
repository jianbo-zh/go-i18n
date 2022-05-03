
Install:
```bash
# install extract tool
go install github.com/jianbo-zh/go-i18n/cmd/i18n@latest

# extract translation content
i18n extract -extract.dir=./cmd/ -output.dir=./translations

# generator translations to gofile
i18n generator -output.dir=./translations

# project init
import (
    _ "path/to/translations"
)

# import i18n package for use
import (
    goi18n "github.com/jianbo-zh/go-i18n"
)
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