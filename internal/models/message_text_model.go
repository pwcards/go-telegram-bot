package models

type MessageUser struct {
	ID          int    `db:"id" json:"id"`
	UserID      int    `db:"user_id" json:"user_id"`
	MessageText string `db:"message_text" json:"message_text"`
	DateAction  string `db:"date_action" json:"date_action"`
}
