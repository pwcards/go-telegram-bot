package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/pwcards/go-telegram-bot/internal/handler"
)

// глобальная переменная в которой храним токен
var telegramBotToken string

const TelegramToken = "5268346289:AAEjBhmqA1wyAjPMXd7S7mP2r_6t9pKp9TM"

func Execute() {
	handler.MessageHandler(TelegramToken)
}

func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", TelegramToken, "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}
