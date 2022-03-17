package handler

import (
	"fmt"
	"log" //nolint:goimports

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

const (
	ReplyWelcome = "Привет, %s!\nВас приветствует бот для отслеживания курсов валют.\nВы можете отслеживать, как отдельную валюту сами, или настроить ежедневное оповещение.\n\nСейчас мы отслеживаем курсы %s, %s и %s."
	ReplyValute  = "Текущий курс %s: <strong>%.2f руб.</strong>"
)

func (h *Handler) initBot(token string) (*telegramApi.BotAPI, error) {
	return telegramApi.NewBotAPI(token)
}

func (h *Handler) MessageHandler(cfg *models.Config) error {
	// Создание бота
	bot, err := h.initBot(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	valute, err := h.GetRemoteDataValute()
	if err != nil {
		return errors.Wrap(err, "get valute remote source")
	}

	// Логирование пользователя, который зашел в bot.
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Структура с конфигом для получения updates
	u := telegramApi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, _ := bot.GetUpdatesChan(u)

	// В канал updates приходят структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
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
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := telegramApi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			// Обработка команд.
			// Команда - сообщение, начинающееся с "/"
			switch update.Message.Command() {
			case "start":
				msg.Text = fmt.Sprintf(
					ReplyWelcome,
					update.Message.From.FirstName,
					models.GetValuteItemNameEmoji(models.ValuteUSD),
					models.GetValuteItemNameEmoji(models.ValuteEUR),
					models.GetValuteItemNameEmoji(models.ValuteGBP),
				)
			}
		} else {
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = h.GetKeyboard()
			case "close":
				msg.ReplyMarkup = telegramApi.NewRemoveKeyboard(true)

			case models.ValuteUSD, models.ValuteEUR, models.ValuteGBP:
				objectValute, err := valute.GetObject("Valute", update.Message.Text)
				if err != nil {
					return errors.Wrap(err, "failed to get currency structure")
				}

				value, err := objectValute.GetFloat64("Value")
				if err != nil {
					return errors.Wrap(err, "failed to get currency value")
				}

				msg.ParseMode = "html"
				msg.Text = fmt.Sprintf(
					ReplyValute,
					models.GetValuteItemName(update.Message.Text),
					value,
				)
			}
		}

		// Лог сообщения, которое ответил bot.
		log.Printf("[%s] %s", "BOT", msg.Text)

		if msg.Text != "" {
			err = h.SaveMessageReply(update, msg)
			if err != nil {
				return err
			}

			// Отправка сообщения
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
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
			log.Printf("Created user with id: %d", userID)
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
