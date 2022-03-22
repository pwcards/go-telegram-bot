package handler

import (
	"fmt"

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

func (h *Handler) MessageHandler() error {
	// Логирование пользователя, который зашел в bot.
	h.Log.Info().
		Str("channel", "application").
		Msgf("Authorized on account %s", h.Bot.Self.UserName)

	// Структура с конфигом для получения updates
	u := telegramApi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, _ := h.Bot.GetUpdatesChan(u)

	// В канал updates приходят структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		if update.CallbackQuery != nil {
			err := h.CallBackSwitch(h.Bot, update.CallbackQuery)
			if err != nil {
				return errors.Wrap(err, "save callback start_time")
			}
			continue
		}

		if update.Message == nil { // ignore non-Message updates
			continue
		}

		err := h.SaveUser(update)
		if err != nil {
			return errors.Wrap(err, "save user local db")
		}

		err = h.SaveMessageUser(update)
		if err != nil {
			return err
		}

		// Лог сообщения, которое написал пользователь.
		h.Log.Log().
			Str("channel", "application").
			Msgf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := telegramApi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			// Обработка команд.
			// Команда - сообщение, начинающееся с "/"
			switch update.Message.Command() {
			case models.CommandStart:
				msg.Text = fmt.Sprintf(
					models.ReplyWelcome,
					update.Message.From.FirstName,
					models.GetValuteItemNameEmoji(models.ValuteUSD),
					models.GetValuteItemNameEmoji(models.ValuteEUR),
					models.GetValuteItemNameEmoji(models.ValuteGBP),
				)
			case models.CommandValutes:
				msg.Text = models.ReplySelectValute
				msg.ReplyMarkup = h.GetKeyboardValute()

			case models.CommandSummary:
				msg.Text = models.ReplySummaryRequest
				msg.ReplyMarkup = h.GetKeyboardTime()
			}
		} else {
			msg.ParseMode = telegramApi.ModeHTML

			valute, err := h.GetCurrentValute()
			if err != nil {
				return errors.Wrap(err, "get valute remote source")
			}

			switch update.Message.Text {
			case models.ValuteUSD:
				msg.Text = fmt.Sprintf(
					models.ReplyValute,
					models.GetValuteItemName(update.Message.Text),
					valute.Usd,
				)
			case models.ValuteEUR:
				msg.Text = fmt.Sprintf(
					models.ReplyValute,
					models.GetValuteItemName(update.Message.Text),
					valute.Eur,
				)
			case models.ValuteGBP:
				msg.Text = fmt.Sprintf(
					models.ReplyValute,
					models.GetValuteItemName(update.Message.Text),
					valute.Gbp,
				)
			default:
				msg.Text = models.ReplyUndefined
			}

			h.keyboardClose(&msg)
		}

		// Лог сообщения, которое ответил bot.
		h.Log.Log().
			Str("channel", "application").
			Msgf("[%s] %s", "BOT", msg.Text)

		if msg.Text != "" {
			err = h.SaveMessageReply(update, msg)
			if err != nil {
				return err
			}

			// Отправка сообщения
			if _, err := h.Bot.Send(msg); err != nil {
				h.Log.Panic().Err(err)
			}
		}
	}

	return nil
}

func (h Handler) SaveUser(update telegramApi.Update) error {
	item, err := h.UserRepository.FindUserItem(update.Message.From.ID)
	if err != nil {
		return errors.Wrap(err, "find user item")
	}
	if item.ID == 0 {
		userID, err := h.UserRepository.InsertUser(update.Message.From)
		if err != nil {
			return errors.Wrap(err, "insert user item")
		} else {
			h.Log.Info().
				Str("channel", "application").
				Msgf("Created user with id: %d", userID)
		}
	}

	return nil
}

func (h Handler) SaveMessageUser(update telegramApi.Update) error {
	err := h.MessageUserRepository.InsertMessage(update.Message.From.ID, update.Message.Text)
	if err != nil {
		return errors.Wrap(err, "insert message user")
	}

	return nil
}

func (h Handler) SaveMessageReply(update telegramApi.Update, msg telegramApi.MessageConfig) error {
	err := h.MessageReplyRepository.InsertMessage(update.Message.From.ID, msg.Text)
	if err != nil {
		return errors.Wrap(err, "insert message reply")
	}

	return nil
}

// keyboardClose will hide the keyboard.
func (h Handler) keyboardClose(msg *telegramApi.MessageConfig) {
	msg.ReplyMarkup = telegramApi.NewRemoveKeyboard(true)
}
