// filepath: backend_go/handlers/groups.go
package handlers

import (
    "backend_go/database"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetGroups(c *gin.Context) {
    rows, err := database.DB.Query("SELECT id, name FROM groups")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var groups []struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }

    for rows.Next() {
        var group struct {
            ID   int    `json:"id"`
            Name string `json:"name"`
        }
        if err := rows.Scan(&group.ID, &group.Name); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        groups = append(groups, group)
    }

    c.JSON(http.StatusOK, groups)
}

func GetGroupByID(c *gin.Context) {
    id := c.Param("id")

    var group struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }

    err := database.DB.QueryRow("SELECT id, name FROM groups WHERE id = ?", id).Scan(&group.ID, &group.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, group)
}

func GetGroupWords(c *gin.Context) {
    id := c.Param("id")

    rows, err := database.DB.Query("SELECT w.id, w.french, w.english, w.parts FROM words w JOIN words_groups wg ON w.id = wg.word_id WHERE wg.group_id = ?", id)
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

func GetGroupSessions(c *gin.Context) {
    id := c.Param("id")

    rows, err := database.DB.Query("SELECT id, group_id, created_at FROM sessions WHERE group_id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var sessions []struct {
        ID        int    `json:"id"`
        GroupID   int    `json:"group_id"`
        CreatedAt string `json:"created_at"`
    }

    for rows.Next() {
        var session struct {
            ID        int    `json:"id"`
            GroupID   int    `json:"group_id"`
            CreatedAt string `json:"created_at"`
        }
        if err := rows.Scan(&session.ID, &session.GroupID, &session.CreatedAt); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        sessions = append(sessions, session)
    }

    c.JSON(http.StatusOK, sessions)
}