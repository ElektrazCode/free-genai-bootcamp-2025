package models

type Word struct {
	ID      int    `json:"id"`
	French  string `json:"french"`
	English string `json:"english"`
	Parts   string `json:"parts"`
}
