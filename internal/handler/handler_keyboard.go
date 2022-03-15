package handler

import (
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

var valuteKeyboard = telegramApi.NewReplyKeyboard(
	telegramApi.NewKeyboardButtonRow(
		telegramApi.NewKeyboardButton(
			models.GetValuteItemShortName(models.ValuteUSD),
		),
		telegramApi.NewKeyboardButton(
			models.GetValuteItemShortName(models.ValuteEUR),
		),
		telegramApi.NewKeyboardButton(
			models.GetValuteItemShortName(models.ValuteGBP),
		),
	),
)
