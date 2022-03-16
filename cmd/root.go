package cmd

import (
	"log"

	"github.com/pwcards/go-telegram-bot/internal/config"
	"github.com/pwcards/go-telegram-bot/internal/handler"
	"github.com/pwcards/go-telegram-bot/internal/repository"
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

	connect := GetConnect(cfg)

	h := handler.NewHandler(
		handler.WithUserRepository(repository.NewUser(connect)),
		handler.WithMessageUserRepository(repository.NewMessageUser(connect)),
		handler.WithMessageReplyRepository(repository.NewMessageReply(connect)),
	)

	err = h.MessageHandler(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
