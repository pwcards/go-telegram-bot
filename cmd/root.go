package cmd

import (
	"os"

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/config"
	"github.com/pwcards/go-telegram-bot/internal/cron"
	"github.com/pwcards/go-telegram-bot/internal/handler"
	"github.com/pwcards/go-telegram-bot/internal/models"
	"github.com/pwcards/go-telegram-bot/internal/repository"
	"github.com/rs/zerolog"
)

func Execute() {
	// Start logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Generate our config based on the config supplied
	// by the user in the flags
	cfgPath, err := config.ParseFlags()
	if err != nil {
		logger.Panic().Err(err)
	}
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		logger.Panic().Err(err)
	}

	// Init bot
	bot, err := startBot(cfg)
	if err != nil {
		logger.Panic().Err(err)
	}

	// Connect to DB
	connect := GetConnect(cfg)

	// Init handlers
	h := handler.NewHandler(
		handler.WithCfg(cfg),
		handler.WithBot(bot),
		handler.WithLogger(logger),
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

	// Init general message handler
	err = h.MessageHandler()
	if err != nil {
		logger.Panic().Err(err)
	}
}

// startBot init bot telegram
func startBot(Cfg *models.Config) (*telegramApi.BotAPI, error) {
	bot, err := telegramApi.NewBotAPI(Cfg.Telegram.Token)
	if err != nil {
		return &telegramApi.BotAPI{}, err
	}

	return bot, nil
}
