package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

type SummaryRepository interface {
	Create(userID int, value string) (int64, error)
	FindItem(userID int) (*models.SummaryModel, error)
	UpdateItem(userID int, value string) error
}

type summary struct {
	db *sqlx.DB
}

func NewSummary(db *sqlx.DB) SummaryRepository {
	return &summary{db: db}
}

func (h summary) Create(userID int, value string) (int64, error) {
	res, err := h.db.Exec(`
		INSERT INTO summary (user_id, time_action) 
		VALUES(?, ?)`,
		userID, value,
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

func (h summary) FindItem(userID int) (*models.SummaryModel, error) {
	itemUser := &models.SummaryModel{}
	err := h.db.Get(itemUser, `
		SELECT id 
		FROM summary 
		WHERE user_id = ?`, userID)

	if err == sql.ErrNoRows {
		return itemUser, nil
	}

	if err != nil {
		return itemUser, errors.Wrap(err, "select summary item")
	}

	return itemUser, nil
}

func (h summary) UpdateItem(userID int, value string) error {
	_, err := h.db.Exec(`
		UPDATE summary 
		SET time_action = ?
		WHERE user_id = ?`,
		value, userID,
	)
	if err != nil {
		return errors.Wrap(err, "repository: save summary")
	}

	return nil
}
