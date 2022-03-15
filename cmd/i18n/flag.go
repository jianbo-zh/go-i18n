package main

import "flag"

type CommandConfig struct {
	ExtractDir   string `json:"extractDir"`   // 提取根目录
	OutputDir    string `json:"outputDir"`    // 输出目录
	DefaultLang  string `json:"defaultLang"`  // 默认语言
	SupportLangs string `json:"supportLangs"` // 支持语言
	Packname     string `json:"packname"`     // 包名
}

func ParseFlag() CommandConfig {
	var initConf CommandConfig

	flag.StringVar(&initConf.Packname, "packname", "i18n", "pack name")
	flag.StringVar(&initConf.ExtractDir, "extract.dir", ".", "The directory containing the files to be translated")
	flag.StringVar(&initConf.OutputDir, "output.dir", "translations", "output directory for translation strings")
	flag.StringVar(&initConf.DefaultLang, "default.language", "en_US", "default language")
	flag.StringVar(&initConf.SupportLangs, "support.languages", "en_US,zh_CN", "all support languages")

	flag.Parse()

	return initConf
}
