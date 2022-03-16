package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type MessageReplyRepository interface {
	InsertMessage(userID int, text string) error
}

type messageReply struct {
	db *sqlx.DB
}

func NewMessageReply(db *sqlx.DB) MessageReplyRepository {
	return &messageReply{db: db}
}

func (h messageReply) InsertMessage(userID int, messageText string) error {
	_, err := h.db.Exec(`
		INSERT INTO message_reply (user_id, message_text, date_action) 
		VALUES(?, ?, ?)`,
		userID, messageText, time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
