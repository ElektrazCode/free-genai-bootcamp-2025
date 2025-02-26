// filepath: backend_go/handlers/words.go
package handlers

import (
    "backend_go/database"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetWords(c *gin.Context) {
    rows, err := database.DB.Query("SELECT id, french, english, parts FROM words")
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

func GetWordByID(c *gin.Context) {
    id := c.Param("id")

    var word struct {
        ID      int    `json:"id"`
        French  string `json:"french"`
        English string `json:"english"`
        Parts   string `json:"parts"`
    }

    err := database.DB.QueryRow("SELECT id, french, english, parts FROM words WHERE id = ?", id).Scan(&word.ID, &word.French, &word.English, &word.Parts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, word)
}