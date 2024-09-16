package main

import (
	"log"
	"vpapidokha/emailvalidator/internal/configurator"
)

func main() {
	cfgPath, err := configurator.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := configurator.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Run()
}
