package cron

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/pwcards/go-telegram-bot/internal/handler"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

type cronWorker struct {
	handler *handler.Handler
}

func (w *cronWorker) Run() {
	s := gocron.NewScheduler(time.UTC)

	log.Print("[CRON] start")

	// Получение данных из удаленного источника
	_, err := s.Every(1).Hour().Do(
		func() {
			_, err := w.handler.GetCurrentValute()
			if err != nil {
				return
			}
			log.Print("[CRON] execute: GetRemoteData")
		})

	// Ежедневная сводка (08:00)
	_, err = s.Every(1).Day().At(models.GetTimeMapValue(models.TimeKeySend08)).Do(
		func() {
			w.handler.SendSummaryList(models.TimeKeySend08)
			log.Print("[CRON] execute: SendSummary_08:00")
		})

	// Ежедневная сводка (09:00)
	_, err = s.Every(1).Day().At(models.GetTimeMapValue(models.TimeKeySend09)).Do(
		func() {
			w.handler.SendSummaryList(models.TimeKeySend09)
			log.Print("[CRON] execute: SendSummary_09:00")
		})

	// Ежедневная сводка (10:00)
	_, err = s.Every(1).Day().At(models.GetTimeMapValue(models.TimeKeySend10)).Do(
		func() {
			w.handler.SendSummaryList(models.TimeKeySend10)
			log.Print("[CRON] execute: SendSummary_10:00")
		})

	if err != nil {
		return
	}

	if err != nil {
		return
	}

	s.StartAsync()
}

type WorkerOption func(c *cronWorker)

func NewCronWorker(opts ...WorkerOption) *cronWorker {
	w := &cronWorker{}
	for _, opt := range opts {
		opt(w)
	}
	return w
}

func WithValutesRepository(h *handler.Handler) WorkerOption {
	return func(w *cronWorker) {
		w.handler = h
	}
}
