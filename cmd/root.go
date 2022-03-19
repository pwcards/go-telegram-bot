package cmd

import (
	"log"

	"github.com/pwcards/go-telegram-bot/internal/config"
	"github.com/pwcards/go-telegram-bot/internal/cron"
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

	// Connect to DB
	connect := GetConnect(cfg)

	// Init handlers
	h := handler.NewHandler(
		handler.WithCfg(cfg),
		handler.WithUserRepository(repository.NewUser(connect)),
		handler.WithMessageUserRepository(repository.NewMessageUser(connect)),
		handler.WithMessageReplyRepository(repository.NewMessageReply(connect)),
		handler.WithValutesRepository(repository.NewValutes(connect)),
		handler.WithSummaryRepository(repository.NewSummary(connect)),
	)

	// Cron
	cronWorker := cron.NewCronWorker(
		cron.WithValutesRepository(h),
	)
	cronWorker.Run()

	err = h.MessageHandler()
	if err != nil {
		log.Fatal(err)
	}
}
