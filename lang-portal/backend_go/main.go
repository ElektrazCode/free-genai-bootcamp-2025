package main

import (
	"backend_go/database"
	"backend_go/handlers"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	dbPath := filepath.Join("database", "words.db")
	database.InitDB(dbPath)

	// Create a new Gin router
	r := gin.Default()

	// Define routes
	r.GET("/api/dashboard/last_study_session", handlers.GetLastStudySession)
	r.GET("/api/dashboard/study_progress", handlers.GetStudyProgress)
	r.GET("/api/dashboard/quick-stats", handlers.GetQuickStats)
	r.GET("/api/activities", handlers.GetActivities)
	r.GET("/api/activities/:id", handlers.GetActivityByID)
	r.GET("/api/activities/:id/sessions", handlers.GetActivitySessions)
	r.POST("/api/activities", handlers.CreateActivity)
	r.GET("/api/words", handlers.GetWords)
	r.GET("/api/words/:id", handlers.GetWordByID)
	r.GET("/api/groups", handlers.GetGroups)
	r.GET("/api/groups/:id", handlers.GetGroupByID)
	r.GET("/api/groups/:id/words", handlers.GetGroupWords)
	r.GET("/api/groups/:id/sessions", handlers.GetGroupSessions)
	r.GET("/api/sessions", handlers.GetSessions)
	r.GET("/api/sessions/:id", handlers.GetSessionByID)
	r.GET("/api/sessions/:id/words", handlers.GetSessionWords)
	r.POST("/api/reset_history", handlers.ResetHistory)
	r.POST("/api/full_reset", handlers.FullReset)
	r.POST("/api/sessions/:id/words/:word_id/review", handlers.ReviewWord)

	// Start the server
	r.Run(":8080")
}
