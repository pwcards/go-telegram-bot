package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type MessageUserRepository interface {
	InsertMessage(userID int, text string) error
}

type messageUser struct {
	db *sqlx.DB
}

func NewMessageUser(db *sqlx.DB) MessageUserRepository {
	return &messageUser{db: db}
}

func (h messageUser) InsertMessage(userID int, messageText string) error {
	_, err := h.db.Exec(`
		INSERT INTO message_user (user_id, message_text, date_action) 
		VALUES(?, ?, ?)`,
		userID, messageText, time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
