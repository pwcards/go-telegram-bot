package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

type SummaryRepository interface {
	Create(userID int, chatID int64, value string) (int64, error)
	FindItem(userID int, chatID int64) (*models.SummaryModel, error)
	UpdateItem(userID int, chatID int64, value string) error
	GetUsersByKey(key string) ([]models.SummaryModel, error)
}

type summary struct {
	db *sqlx.DB
}

func NewSummary(db *sqlx.DB) SummaryRepository {
	return &summary{db: db}
}

func (h summary) Create(userID int, chatID int64, value string) (int64, error) {
	res, err := h.db.Exec(`
		INSERT INTO summary (user_id, chat_id, time_action) 
		VALUES(?, ?, ?)`,
		userID, chatID, value,
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

func (h summary) FindItem(userID int, chatID int64) (*models.SummaryModel, error) {
	item := &models.SummaryModel{}
	err := h.db.Get(item, `
		SELECT id 
		FROM summary 
		WHERE user_id = ?
		  AND chat_id = ?`, userID, chatID)

	if err == sql.ErrNoRows {
		return item, nil
	}

	if err != nil {
		return item, errors.Wrap(err, "select summary item")
	}

	return item, nil
}

func (h summary) UpdateItem(userID int, chatID int64, value string) error {
	_, err := h.db.Exec(`
		UPDATE summary 
		SET time_action = ?
		WHERE user_id = ?
		  AND chat_id = ?`, value, userID, chatID,
	)
	if err != nil {
		return errors.Wrap(err, "repository: update summary")
	}

	return nil
}

func (h summary) GetUsersByKey(key string) ([]models.SummaryModel, error) {
	rows, err := h.db.Queryx(`
		SELECT user_id, chat_id 
		FROM summary 
		WHERE time_action = ?`, key)

	if err != nil {
		return nil, errors.Wrap(err, "QueryContext")
	}

	users := make([]models.SummaryModel, 0)
	for rows.Next() {
		var item models.SummaryModel
		if err := rows.StructScan(&item); err != nil {
			return nil, errors.Wrap(err, "rows.Scan")
		}

		users = append(users, item)
	}
	if err := rows.Close(); err != nil {
		return nil, errors.Wrap(err, "rows.Close")
	}

	return users, nil
}
