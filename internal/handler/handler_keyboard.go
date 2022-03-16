package handler

import (
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

func (h Handler) GetKeyboard() telegramApi.ReplyKeyboardMarkup {
	return telegramApi.NewReplyKeyboard(
		telegramApi.NewKeyboardButtonRow(
			telegramApi.NewKeyboardButton(models.ValuteUSD),
			telegramApi.NewKeyboardButton(models.ValuteEUR),
			telegramApi.NewKeyboardButton(models.ValuteGBP),
		),
	)
}
