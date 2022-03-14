package handler

import telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	CommandUSD = "USD 🇺🇸"
	CommandEUR = "EUR 🇪🇺"
	CommandGBP = "GBP 🇬🇧"
)

var valuteKeyboard = telegramApi.NewReplyKeyboard(
	telegramApi.NewKeyboardButtonRow(
		telegramApi.NewKeyboardButton(CommandUSD),
		telegramApi.NewKeyboardButton(CommandEUR),
		telegramApi.NewKeyboardButton(CommandGBP),
	),
)
