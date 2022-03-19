package models

type ValutesModelDB struct {
	ID      int     `db:"id" json:"id"`
	DateVal string  `db:"date_val" json:"date_val"`
	Usd     float64 `db:"usd" json:"usd"`
	Eur     float64 `db:"eur" json:"eur"`
	Gbp     float64 `db:"gbp" json:"gbp"`
}
