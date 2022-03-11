package cmd

import (
	"log"

	"github.com/pwcards/go-telegram-bot/internal/config"
	"github.com/pwcards/go-telegram-bot/internal/handler"
)

func Execute() {
	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := config.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	handler.MessageHandler(cfg)
}
