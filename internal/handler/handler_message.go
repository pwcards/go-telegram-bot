package handler

import (
	"fmt"
	"log" //nolint:goimports

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/config"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

const (
	ReplyWelcome = "Привет, %s!\nВас приветствует бот для отслеживания курсов валют.\nВы можете отслеживать, как отдельную валюту сами, или настроить ежедневное оповещение.\n\nСейчас мы отслеживаем курсы %s, %s и %s."
	ReplyValute  = "Текущий курс %s: <strong>%.2f руб.</strong>"
)

func initBot(token string) (*telegramApi.BotAPI, error) {
	return telegramApi.NewBotAPI(token)
}

func MessageHandler(cfg *config.Config) error {
	// Создание бота
	bot, err := initBot(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	valute, err := GetRemoteDataValute()
	if err != nil {
		return errors.Wrap(err, "get valute remote source")
	}

	// Логирование пользователя, который зашел в bot.
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Структура с конфигом для получения апдейтов
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
					models.GetValuteItemFullName(models.ValuteUSD),
					models.GetValuteItemFullName(models.ValuteEUR),
					models.GetValuteItemFullName(models.ValuteGBP),
				)
			}
		} else {
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = valuteKeyboard
			case "close":
				msg.ReplyMarkup = telegramApi.NewRemoveKeyboard(true)

			case models.GetValuteItemShortName(models.ValuteUSD):
				msg.ParseMode = "html"
				msg.Text = fmt.Sprintf(
					ReplyValute,
					models.GetValuteItem(models.ValuteUSD).Name,
					valute.Valute.Usd.Value,
				)
			case models.GetValuteItemShortName(models.ValuteEUR):
				msg.ParseMode = "html"
				msg.Text = fmt.Sprintf(
					ReplyValute,
					models.GetValuteItem(models.ValuteEUR).Name,
					valute.Valute.Eur.Value,
				)
			case models.GetValuteItemShortName(models.ValuteGBP):
				msg.ParseMode = "html"
				msg.Text = fmt.Sprintf(""+
					ReplyValute,
					models.GetValuteItem(models.ValuteGBP).Name,
					valute.Valute.Gbp.Value,
				)
			}
		}

		// Лог сообщения, которое ответил bot.
		log.Printf("[%s] %s", "BOT", msg.Text)

		// Отправка сообщения
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	return nil
}
