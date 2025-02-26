package models

import "time"

// Response Types
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
}

type WordWithStats struct {
	Japanese     string `json:"japanese"`
	Romaji       string `json:"romaji"`
	English      string `json:"english"`
	CorrectCount int    `json:"correct_count"`
	WrongCount   int    `json:"wrong_count"`
}

type WordDetailResponse struct {
	Japanese string `json:"japanese"`
	Romaji   string `json:"romaji"`
	English  string `json:"english"`
	Stats    struct {
		CorrectCount int `json:"correct_count"`
		WrongCount   int `json:"wrong_count"`
	} `json:"stats"`
	Groups []GroupWithStats `json:"groups"`
}

type GroupWithStats struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	WordCount int    `json:"word_count,omitempty"`
	Stats     struct {
		TotalWordCount int `json:"total_word_count,omitempty"`
	} `json:"stats,omitempty"`
}

type StudySessionResponse struct {
	ID               int       `json:"id"`
	ActivityName     string    `json:"activity_name"`
	GroupName        string    `json:"group_name"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
	ReviewItemsCount int       `json:"review_items_count"`
}

type StudyActivitySessionResponse struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
}

type LastStudySessionResponse struct {
	ID              int       `json:"id"`
	GroupID         int       `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
	StudyActivityID int       `json:"study_activity_id"`
	GroupName       string    `json:"group_name"`
}

type StudyProgressResponse struct {
	TotalWordsStudied   int `json:"total_words_studied"`
	TotalAvailableWords int `json:"total_available_words"`
}

type QuickStatsResponse struct {
	SuccessRate        float64 `json:"success_rate"`
	TotalStudySessions int     `json:"total_study_sessions"`
	TotalActiveGroups  int     `json:"total_active_groups"`
	StudyStreakDays    int     `json:"study_streak_days"`
}

type WordsResponse struct {
	Items      []WordWithStats `json:"items"`
	Pagination *Pagination     `json:"pagination"`
}

type GroupsResponse struct {
	Items      []GroupWithStats `json:"items"`
	Pagination *Pagination      `json:"pagination"`
}
