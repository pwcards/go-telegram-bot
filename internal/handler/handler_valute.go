package handler

import (
	"net/http"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func (h *Handler) GetRemoteDataValute(cfg *models.Config) (*jason.Object, error) {
	r, err := myClient.Get(cfg.ServerData.Host)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	res, _ := jason.NewObjectFromReader(r.Body)

	return res, nil
}
