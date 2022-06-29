package main

import (
	"flag"
)

type CommandConfig struct {
	SubCommand   string `json:"subCommand"`   // 子命令
	ExtractDir   string `json:"extractDir"`   // 提取根目录
	OutputDir    string `json:"outputDir"`    // 输出目录
	DefaultLang  string `json:"defaultLang"`  // 默认语言
	SupportLangs string `json:"supportLangs"` // 支持语言
	Packname     string `json:"packname"`     // 包名
	Debug        bool   `json:"debug"`        // 是否打印调试日志
}

func ParseFlag() CommandConfig {
	var initConf CommandConfig

	flag.StringVar(&initConf.Packname, "packname", "goi18n", "pack name")
	flag.StringVar(&initConf.ExtractDir, "extract.dir", ".", "The directory containing the files to be translated")
	flag.StringVar(&initConf.OutputDir, "output.dir", "translations", "output directory for translation strings")
	flag.StringVar(&initConf.DefaultLang, "default.language", "en", "default language")
	flag.StringVar(&initConf.SupportLangs, "support.languages", "en,zh", "all support languages")
	flag.BoolVar(&initConf.Debug, "debug", false, "if open debug log")

	flag.Parse()

	initConf.SubCommand = flag.Arg(0)

	return initConf
}
