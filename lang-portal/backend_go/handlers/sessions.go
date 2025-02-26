// filepath: backend_go/handlers/sessions.go
package handlers

import (
    "backend_go/database"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetSessions(c *gin.Context) {
    rows, err := database.DB.Query("SELECT id, group_id, created_at, activity_id FROM sessions")
    if (err != nil) {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var sessions []struct {
        ID        int    `json:"id"`
        GroupID   int    `json:"group_id"`
        CreatedAt string `json:"created_at"`
        ActivityID int   `json:"activity_id"`
    }

    for rows.Next() {
        var session struct {
            ID        int    `json:"id"`
            GroupID   int    `json:"group_id"`
            CreatedAt string `json:"created_at"`
            ActivityID int   `json:"activity_id"`
        }
        if err := rows.Scan(&session.ID, &session.GroupID, &session.CreatedAt, &session.ActivityID); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        sessions = append(sessions, session)
    }

    c.JSON(http.StatusOK, sessions)
}

func GetSessionByID(c *gin.Context) {
    id := c.Param("id")

    var session struct {
        ID        int    `json:"id"`
        GroupID   int    `json:"group_id"`
        CreatedAt string `json:"created_at"`
        ActivityID int   `json:"activity_id"`
    }

    err := database.DB.QueryRow("SELECT id, group_id, created_at, activity_id FROM sessions WHERE id = ?", id).Scan(&session.ID, &session.GroupID, &session.CreatedAt, &session.ActivityID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, session)
}

func GetSessionWords(c *gin.Context) {
    id := c.Param("id")

    rows, err := database.DB.Query("SELECT w.id, w.french, w.english, w.parts FROM words w JOIN review_words rw ON w.id = rw.word_id WHERE rw.session_id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var words []struct {
        ID      int    `json:"id"`
        French  string `json:"french"`
        English string `json:"english"`
        Parts   string `json:"parts"`
    }

    for rows.Next() {
        var word struct {
            ID      int    `json:"id"`
            French  string `json:"french"`
            English string `json:"english"`
            Parts   string `json:"parts"`
        }
        if err := rows.Scan(&word.ID, &word.French, &word.English, &word.Parts); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        words = append(words, word)
    }

    c.JSON(http.StatusOK, words)
}

func ResetHistory(c *gin.Context) {
    _, err := database.DB.Exec("DELETE FROM sessions")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    _, err = database.DB.Exec("DELETE FROM review_words")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "History reset successfully"})
}

func FullReset(c *gin.Context) {
    tables := []string{"sessions", "review_words", "activities", "words_groups", "groups", "words"}

    for _, table := range tables {
        _, err := database.DB.Exec("DELETE FROM " + table)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Full reset successfully"})
}

func ReviewWord(c *gin.Context) {
    sessionID := c.Param("id")
    wordID := c.Param("word_id")
    var review struct {
        Correct bool `json:"correct"`
    }

    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := database.DB.Exec("INSERT INTO review_words (word_id, session_id, correct, created_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP)", wordID, sessionID, review.Correct)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Word reviewed successfully"})
}