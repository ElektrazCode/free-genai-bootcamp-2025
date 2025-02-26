// filepath: backend_go/handlers/words.go
package handlers

import (
	"backend_go/database"
	"backend_go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWords(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, french, english, parts FROM words")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var words []models.Word

	for rows.Next() {
		var word models.Word
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

	var word models.Word

	err := database.DB.QueryRow("SELECT id, french, english, parts FROM words WHERE id = ?", id).Scan(&word.ID, &word.French, &word.English, &word.Parts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, word)
}
