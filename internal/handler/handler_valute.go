package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pwcards/go-telegram-bot/internal/models"
)

const remoteSourceData = "https://www.cbr-xml-daily.ru/daily_json.js"

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetRemoteDataValute() (*models.ValuteData, error) {
	data := models.ValuteData{}
	err := getJson(remoteSourceData, &data)
	if err != nil {
		return &data, err
	}

	return &data, nil
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
