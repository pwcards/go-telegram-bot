package handler

import "github.com/pwcards/go-telegram-bot/internal/repository"

type Handler struct {
	UserRepository         repository.UserRepository
	MessageUserRepository  repository.MessageUserRepository
	MessageReplyRepository repository.MessageReplyRepository
}

type Option func(*Handler)

func NewHandler(opts ...Option) *Handler {
	h := &Handler{}
	for _, opt := range opts {
		opt(h)
	}
	return h
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
