package models

type User struct {
	ID        int    `db:"id" json:"id"`
	UserID    int    `db:"user_id" json:"user_id"`
	NickName  string `db:"nickname" json:"nickname"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}
