package handler

import (
	"fmt"
	"log" //nolint:goimports

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/config"
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
				msg.Text = "Привет! 🤗\n Я бот, который рассчитает твою зарплату по новому курсу 💰💰💰. \n Укажи свою текущую зарплату в рублях."
			}
		} else {
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = valuteKeyboard
			case "close":
				msg.ReplyMarkup = telegramApi.NewRemoveKeyboard(true)

			case CommandUSD:
				msg.Text = fmt.Sprintf("Текущий курс доллара: %f руб.", valute.Valute.Usd.Value)
			case CommandEUR:
				msg.Text = fmt.Sprintf("Текущий курс евро: %f руб.", valute.Valute.Eur.Value)
			case CommandGBP:
				msg.Text = fmt.Sprintf("Текущий курс фунта стерлингов Соединенного королевства: %f руб.", valute.Valute.Gbp.Value)
			}
		}

		// Отправка сообщения
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	return nil
}
