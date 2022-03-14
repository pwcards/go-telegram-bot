package models

import "time"

type ValuteData struct {
	Date         time.Time `json:"Date"`
	PreviousDate time.Time `json:"PreviousDate"`
	PreviousURL  string    `json:"PreviousURL"`
	Timestamp    time.Time `json:"Timestamp"`
	Valute       struct {
		Usd struct{ ValuteItem } `json:"USD"`
		Eur struct{ ValuteItem } `json:"EUR"`
		Gbp struct{ ValuteItem } `json:"GBP"`
	} `json:"Valute"`
}

type ValuteItem struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}
