package handler

import (
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

// GetKeyboardValute keyboard valutes.
func (h Handler) GetKeyboardValute() telegramApi.ReplyKeyboardMarkup {
	return telegramApi.NewReplyKeyboard(
		telegramApi.NewKeyboardButtonRow(
			telegramApi.NewKeyboardButton(models.ValuteUSD),
			telegramApi.NewKeyboardButton(models.ValuteEUR),
			telegramApi.NewKeyboardButton(models.ValuteGBP),
		),
	)
}

// GetKeyboardTime keyboard times.
func (h Handler) GetKeyboardTime() telegramApi.InlineKeyboardMarkup {
	return telegramApi.NewInlineKeyboardMarkup(
		telegramApi.NewInlineKeyboardRow(
			telegramApi.NewInlineKeyboardButtonData(models.GetTimeMapValue(models.TimeKeySend08), models.TimeKeySend08),
			telegramApi.NewInlineKeyboardButtonData(models.GetTimeMapValue(models.TimeKeySend09), models.TimeKeySend09),
			telegramApi.NewInlineKeyboardButtonData(models.GetTimeMapValue(models.TimeKeySend10), models.TimeKeySend10),
		),
	)
}
