package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/pwcards/go-telegram-bot/internal/models"
)

type ValutesRepository interface {
	Create(now string, valutes *models.ValutesModelDB) (int64, error)
	FindValuteItem(now string) (*models.ValutesModelDB, error)
	UpdateItem(now string, valutes *models.ValutesModelDB) error
}

type valutes struct {
	db *sqlx.DB
}

func NewValutes(db *sqlx.DB) ValutesRepository {
	return &valutes{db: db}
}

func (h valutes) Create(now string, valutes *models.ValutesModelDB) (int64, error) {
	res, err := h.db.Exec(`
		INSERT INTO valutes (date_val, usd, eur, gbp) 
		VALUES(?, ?, ?, ?)`,
		now, valutes.Usd, valutes.Eur, valutes.Gbp,
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

func (h valutes) FindValuteItem(now string) (*models.ValutesModelDB, error) {
	item := &models.ValutesModelDB{}
	err := h.db.Get(item, `
		SELECT id, usd, eur, gbp
		FROM valutes 
		WHERE date_val = ?`, now)

	if err == sql.ErrNoRows {
		return item, nil
	}

	if err != nil {
		return item, errors.Wrap(err, "select valute item")
	}

	return item, nil
}

func (h valutes) UpdateItem(now string, valutes *models.ValutesModelDB) error {
	_, err := h.db.Exec(`
		UPDATE valutes 
		SET usd = ?,
		    eur = ?,
		    gbp = ?
		WHERE date_val = ?`,
		valutes.Usd, valutes.Eur, valutes.Gbp, now,
	)
	if err != nil {
		return errors.Wrap(err, "repository: update valute")
	}

	return nil
}
