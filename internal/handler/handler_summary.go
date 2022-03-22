package handler

import (
	"fmt"
	"time"

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

func (h Handler) SendSummaryList(key string) {
	usersByKey, err := h.SummaryRepository.GetUsersByKey(key)
	if err != nil {
		return
	}

	if len(usersByKey) > 0 {
		// Получение курсов
		valute, err := h.GetCurrentValute()
		if err != nil {
			return
		}

		for _, element := range usersByKey {
			msg := telegramApi.NewMessage(element.ChatID, "")
			msg.Text = fmt.Sprintf(
				models.ReplyEveryDaySummary,
				time.Now().Format("2006-01-02"),
				models.GetValuteItemNameEmoji(models.ValuteUSD),
				valute.Usd,
				models.GetValuteItemNameEmoji(models.ValuteEUR),
				valute.Eur,
				models.GetValuteItemNameEmoji(models.ValuteGBP),
				valute.Gbp,
			)

			h.Log.Info().
				Str("channel", "summary").
				Str("event", "scheduled").
				Int64("user_id", element.ChatID).
				Msg("Send summary message")

			if _, err := h.Bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
