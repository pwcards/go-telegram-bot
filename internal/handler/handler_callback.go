package handler

import (
	"fmt"

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

func (h Handler) CallBackSwitch(bot *telegramApi.BotAPI, callback *telegramApi.CallbackQuery) error {
	switch callback.Data {
	case models.TimeKeySend08, models.TimeKeySend09, models.TimeKeySend10:
		err := h.CallBackTime(bot, callback)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h Handler) CallBackTime(bot *telegramApi.BotAPI, callback *telegramApi.CallbackQuery) error {
	findItem, err := h.SummaryRepository.FindItem(callback.From.ID)
	if err != nil {
		return err
	}

	if findItem.ID != 0 {
		err := h.SummaryRepository.UpdateItem(callback.From.ID, callback.Data)
		if err != nil {
			return err
		}
	} else {
		_, err := h.SummaryRepository.Create(callback.From.ID, callback.Data)
		if err != nil {
			return err
		}
	}

	// Отправка подтверждения
	msg := telegramApi.NewMessage(callback.Message.Chat.ID, callback.Data)
	msg.Text = fmt.Sprintf(models.ReplySelectTime, models.GetTimeMapValue(callback.Data))
	h.keyboardClose(&msg)

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

	return nil
}
