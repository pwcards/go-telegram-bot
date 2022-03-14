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
		// Универсальный ответ на любое сообщение
		reply := fmt.Sprintf("У меня нет ответа на твой вопрос, %s", update.Message.From.FirstName)
		if update.Message == nil {
			continue
		}

		// Лог сообщения, которое написал пользователь.
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Обработка команд.
		// Команда - сообщение, начинающееся с "/"
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				reply = "Привет! 🤗\n Я бот, который рассчитает твою зарплату по новому курсу 💰💰💰. \n Укажи свою текущую зарплату в рублях."
			case "usd_now":
				reply = fmt.Sprintf("Текущий курс доллара: %f руб.", valute.Valute.Usd.Value)
			}
		}

		// Создаем ответное сообщение
		sendMessage(bot, update, reply)
	}

	return nil
}

// sendMessage отправит сообщение в ответ.
func sendMessage(bot *telegramApi.BotAPI, update telegramApi.Update, reply string) {
	msg := telegramApi.NewMessage(update.Message.Chat.ID, reply)
	// Отправляем
	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}
