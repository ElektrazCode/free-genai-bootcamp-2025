// filepath: backend_go/models/activity.go
package models

type Activity struct {
	ID        int    `json:"id"`
	SessionID int    `json:"session_id"`
	GroupID   int    `json:"group_id"`
	CreatedAt string `json:"created_at"`
}
