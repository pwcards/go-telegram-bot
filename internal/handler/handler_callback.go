package handler

import (
	"fmt"

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

func (h Handler) CallBackSwitch(callback *telegramApi.CallbackQuery) error {
	switch callback.Data {
	case models.TimeKeySend08, models.TimeKeySend09, models.TimeKeySend10:
		err := h.CallBackTime(callback)
		if err != nil {
			return err
		}
	case models.TimeKeyNotActive:
		err := h.CallBackTimeNotActive(callback)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h Handler) CallBackTime(callback *telegramApi.CallbackQuery) error {
	findItem, err := h.SummaryRepository.FindItem(callback.From.ID, callback.Message.Chat.ID)
	if err != nil {
		return err
	}

	if findItem.ID != 0 {
		err := h.SummaryRepository.UpdateItem(callback.From.ID, callback.Message.Chat.ID, callback.Data)
		if err != nil {
			return err
		}
	} else {
		_, err := h.SummaryRepository.Create(callback.From.ID, callback.Message.Chat.ID, callback.Data)
		if err != nil {
			return err
		}
	}

	// Отправка подтверждения
	reply := fmt.Sprintf(models.ReplySelectTime, models.GetTimeMapValue(callback.Data))
	err = h.SendMessageCallback(callback, reply)
	if err != nil {
		return errors.Wrapf(err, "send message callback")
	}

	return nil
}

func (h Handler) CallBackTimeNotActive(callback *telegramApi.CallbackQuery) error {
	findItem, err := h.SummaryRepository.FindItem(callback.From.ID, callback.Message.Chat.ID)
	if err != nil {
		return err
	}

	if findItem.ID != 0 {
		err := h.SummaryRepository.UpdateItemNotActive(callback.From.ID, callback.Message.Chat.ID)
		if err != nil {
			return err
		}
	}

	// Отправка подтверждения
	reply := models.ReplySummaryNotActive
	err = h.SendMessageCallback(callback, reply)
	if err != nil {
		return errors.Wrapf(err, "send message callback")
	}

	return nil
}

// SendMessageCallback отправит подтверждение.
func (h Handler) SendMessageCallback(callback *telegramApi.CallbackQuery, reply string) error {
	msg := telegramApi.NewMessage(callback.Message.Chat.ID, callback.Data)
	msg.Text = reply
	h.keyboardClose(&msg)

	if _, err := h.Bot.Send(msg); err != nil {
		return err
	}

	return nil
}
