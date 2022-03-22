package handler

import (
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pwcards/go-telegram-bot/internal/models"
	"github.com/pwcards/go-telegram-bot/internal/repository"
	"github.com/rs/zerolog"
)

type Handler struct {
	Cfg                    *models.Config
	Bot                    *telegramApi.BotAPI
	Log                    zerolog.Logger
	UserRepository         repository.UserRepository
	MessageUserRepository  repository.MessageUserRepository
	MessageReplyRepository repository.MessageReplyRepository
	ValutesRepository      repository.ValutesRepository
	SummaryRepository      repository.SummaryRepository
}

type Option func(*Handler)

func NewHandler(opts ...Option) *Handler {
	h := &Handler{}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func WithCfg(r *models.Config) Option {
	return func(h *Handler) {
		h.Cfg = r
	}
}

func WithBot(r *telegramApi.BotAPI) Option {
	return func(h *Handler) {
		h.Bot = r
	}
}

func WithUserRepository(r repository.UserRepository) Option {
	return func(h *Handler) {
		h.UserRepository = r
	}
}

func WithMessageUserRepository(r repository.MessageUserRepository) Option {
	return func(h *Handler) {
		h.MessageUserRepository = r
	}
}

func WithMessageReplyRepository(r repository.MessageReplyRepository) Option {
	return func(h *Handler) {
		h.MessageReplyRepository = r
	}
}

func WithValutesRepository(r repository.ValutesRepository) Option {
	return func(h *Handler) {
		h.ValutesRepository = r
	}
}

func WithSummaryRepository(r repository.SummaryRepository) Option {
	return func(h *Handler) {
		h.SummaryRepository = r
	}
}

func WithLogger(log zerolog.Logger) Option {
	return func(h *Handler) {
		h.Log = log
	}
}
