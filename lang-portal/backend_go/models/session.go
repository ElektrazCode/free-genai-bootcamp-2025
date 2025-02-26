package models

type Session struct {
	ID         int    `json:"id"`
	GroupID    int    `json:"group_id"`
	CreatedAt  string `json:"created_at"`
	ActivityID int    `json:"activity_id"`
}
