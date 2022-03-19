package models

type SummaryModel struct {
	ID         int    `db:"id" json:"id"`
	UserID     int    `db:"user_id" json:"user_id"`
	ChatID     int64  `db:"chat_id" json:"chat_id"`
	TimeAction string `db:"time_action" json:"time_action"`
}
