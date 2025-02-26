package models

import "time"

type Activity struct {
	ID        int       `json:"id"`
	SessionID int       `json:"session_id"`
	GroupID   int       `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Session struct {
	ID         int       `json:"id"`
	GroupID    int       `json:"group_id"`
	CreatedAt  time.Time `json:"created_at"`
	ActivityID int       `json:"activity_id"`
}

type Word struct {
	ID      int    `json:"id"`
	French  string `json:"french"`
	English string `json:"english"`
	Parts   string `json:"parts"`
}

type WordReview struct {
	WordID         int       `json:"word_id"`
	StudySessionID int       `json:"study_session_id"`
	Correct        bool      `json:"correct"`
	CreatedAt      time.Time `json:"created_at"`
}
