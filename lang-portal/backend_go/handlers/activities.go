// filepath: backend_go/handlers/activities.go
package handlers

import (
	"backend_go/database"
	"backend_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActivities(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, session_id, group_id, created_at FROM activities")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var activities []models.Activity

	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(&activity.ID, &activity.SessionID, &activity.GroupID, &activity.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		activities = append(activities, activity)
	}

	c.JSON(http.StatusOK, activities)
}

func GetActivityByID(c *gin.Context) {
	id := c.Param("id")

	var activity models.Activity

	err := database.DB.QueryRow("SELECT id, session_id, group_id, created_at FROM activities WHERE id = ?", id).Scan(&activity.ID, &activity.SessionID, &activity.GroupID, &activity.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func GetActivitySessions(c *gin.Context) {
	id := c.Param("id")

	rows, err := database.DB.Query("SELECT id, group_id, created_at FROM sessions WHERE activity_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var sessions []models.Session

	for rows.Next() {
		var session models.Session
		if err := rows.Scan(&session.ID, &session.GroupID, &session.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		sessions = append(sessions, session)
	}

	c.JSON(http.StatusOK, sessions)
}

func CreateActivity(c *gin.Context) {
	var activity models.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.DB.Exec("INSERT INTO activities (session_id, group_id) VALUES (?, ?)", activity.SessionID, activity.GroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
