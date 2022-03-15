package handler

import (
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

var valuteKeyboard = telegramApi.NewReplyKeyboard(
	telegramApi.NewKeyboardButtonRow(
		telegramApi.NewKeyboardButton(models.ValuteUSD),
		telegramApi.NewKeyboardButton(models.ValuteEUR),
		telegramApi.NewKeyboardButton(models.ValuteGBP),
	),
)
