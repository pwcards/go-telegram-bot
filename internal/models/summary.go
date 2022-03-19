package models

type SummaryModel struct {
	ID         int    `db:"id" json:"id"`
	UserID     int    `db:"user_id" json:"user_id"`
	TimeAction string `db:"time_action" json:"time_action"`
}
