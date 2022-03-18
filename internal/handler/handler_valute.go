package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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
		valute, err := h.GetRemoteDataValute()
		if err != nil {
			return &model, errors.Wrap(err, "get valute remote source")
		}

		dataModel := &models.ValutesModelDB{
			Usd: valute.Valute.Eur.Value,
			Eur: valute.Valute.Usd.Value,
			Gbp: valute.Valute.Gbp.Value,
		}

		// Запись значения валют за текущий день
		userID, err := h.ValutesRepository.Create(nowDate, dataModel)
		if err != nil {
			return &model, errors.Wrap(err, "insert user item")
		} else {
			log.Printf("Created user with id: %d", userID)
		}

		return dataModel, nil
	}

	return item, nil
}

func (h Handler) getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
