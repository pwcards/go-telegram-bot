package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	telegramApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

type UserRepository interface {
	InsertUser(user *telegramApi.User) (int64, error)
	FindUserItem(userID int) (*models.User, error)
}

type user struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) UserRepository {
	return &user{db: db}
}

func (h user) InsertUser(user *telegramApi.User) (int64, error) {
	res, err := h.db.Exec(`
		INSERT INTO users (user_id, nickname, first_name, last_name, date_action) 
		VALUES(?, ?, ?, ?, ?)`,
		user.ID, user.UserName, user.FirstName, user.LastName, time.Now(),
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (h user) FindUserItem(userID int) (*models.User, error) {
	itemUser := &models.User{}
	err := h.db.Get(itemUser, `
		SELECT id 
		FROM users 
		WHERE user_id = ?`, userID)

	if err == sql.ErrNoRows {
		return itemUser, nil
	}

	if err != nil {
		return itemUser, errors.Wrap(err, "select user item")
	}

	return itemUser, nil
}
