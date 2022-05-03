package main

import (
	"log"
)

func main() {

	config := ParseFlag()

	if config.SubCommand == "extract" {
		extract(config)

	} else if config.SubCommand == "generator" {
		generator(config)

	} else {
		log.Fatal("i18n subcommand must be \"extract\" or \"generator\"")
	}
}
