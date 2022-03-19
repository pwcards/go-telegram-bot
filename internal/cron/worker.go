package cron

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/pwcards/go-telegram-bot/internal/handler"
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
