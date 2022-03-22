package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func (h Handler) GetRemoteDataValute() (*models.ValuteData, error) {
	data := models.ValuteData{}
	err := h.getJson(h.Cfg.ServerData.Host, &data)
	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (h Handler) GetCurrentValute() (*models.ValutesModelDB, error) {
	nowDate := time.Now().Format("2006-01-02")
	model := models.ValutesModelDB{}

	item, err := h.ValutesRepository.FindValuteItem(nowDate)
	if err != nil {
		return &model, errors.Wrap(err, "find user item")
	}

	if item.ID == 0 {
		// Получение удаленных данных
		dataModel, err := h.GetData()
		if err != nil {
			return nil, err
		}

		// Запись значения валют за текущий день
		_, err = h.ValutesRepository.Create(nowDate, dataModel)
		if err != nil {
			return &model, errors.Wrap(err, "insert user item")
		}

		return dataModel, nil
	} else {
		// Получение удаленных данных
		dataModel, err := h.GetData()
		if err != nil {
			return nil, err
		}

		// Отправка персональных оповещений пользователям (изменение конкретной валюты)
		h.CheckValueChanges(item, dataModel)

		err = h.ValutesRepository.UpdateItem(nowDate, dataModel)

		if err != nil {
			return nil, err
		}
	}

	return item, nil
}

func (h Handler) GetData() (*models.ValutesModelDB, error) {
	model := models.ValutesModelDB{}

	valute, err := h.GetRemoteDataValute()
	if err != nil {
		return &model, errors.Wrap(err, "get valute remote source")
	}

	return &models.ValutesModelDB{
		Usd: valute.Valute.Eur.Value,
		Eur: valute.Valute.Usd.Value,
		Gbp: valute.Valute.Gbp.Value,
	}, nil
}

func (h Handler) CheckValueChanges(localModel, remoteModel *models.ValutesModelDB) {
	listSummary, err := h.SummaryRepository.GetListSummary()
	if err != nil {
		return
	}

	if localModel.Usd != remoteModel.Usd {
		h.SendMessageCourseChangeRealtime(
			listSummary,
			models.ValuteUSD,
			h.getDifferentString(localModel.Usd, remoteModel.Usd),
			math.Abs(localModel.Usd-remoteModel.Usd),
			localModel.Usd,
			remoteModel.Usd,
		)
	}

	if localModel.Eur != remoteModel.Eur {
		h.SendMessageCourseChangeRealtime(
			listSummary,
			models.ValuteEUR,
			h.getDifferentString(localModel.Eur, remoteModel.Eur),
			math.Abs(localModel.Eur-remoteModel.Eur),
			localModel.Eur,
			remoteModel.Eur,
		)
	}

	if localModel.Gbp != remoteModel.Gbp {
		h.SendMessageCourseChangeRealtime(
			listSummary,
			models.ValuteGBP,
			h.getDifferentString(localModel.Gbp, remoteModel.Gbp),
			math.Abs(localModel.Gbp-remoteModel.Gbp),
			localModel.Gbp,
			remoteModel.Gbp,
		)
	}
}

func (h Handler) SendMessageCourseChangeRealtime(listSummary []models.SummaryModel, valuteItem string,
	differentString string, differentCount float64, lastValue float64, nowValue float64) {
	for _, element := range listSummary {
		msg := telegramApi.NewMessage(element.ChatID, "")
		msg.Text = fmt.Sprintf(
			models.ReplyChangeCourseData,
			models.GetValuteItemNameEmoji(valuteItem),
			differentString,
			differentCount,
			nowValue,
			lastValue,
		)

		h.Log.Info().
			Str("channel", "summary").
			Str("event", "different value").
			Int64("user_id", element.ChatID).
			Msg("Send summary message")

		if _, err := h.Bot.Send(msg); err != nil {
			panic(err)
		}
	}
}

func (h Handler) getDifferentString(local float64, remote float64) string {
	differentString := models.DifferenceGrown
	if local > remote {
		differentString = models.DifferenceDecreased
	}

	return differentString
}

func (h Handler) getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
