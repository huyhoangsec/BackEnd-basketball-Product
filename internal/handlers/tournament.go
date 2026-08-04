package handlers

import (
	"fmt"
	"net/http"
	"time"

	"backend-ocean-basketball/internal/models"

	"github.com/gin-gonic/gin"
)

// GetTournaments
func GetTournaments(c *gin.Context) {
	var items []models.Tournament
	if err := models.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tournaments"})
		return
	}
	c.JSON(http.StatusOK, items)
}

// CreateTournament
func CreateTournament(c *gin.Context) {
	var item models.Tournament
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if item.ID == "" {
		item.ID = fmt.Sprintf("t-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tournament"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

// UpdateTournament
func UpdateTournament(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tournament ID required"})
		return
	}
	var item models.Tournament
	if err := models.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tournament not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	if err := models.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tournament"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// DeleteTournament
func DeleteTournament(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tournament ID"})
		return
	}
	if err := models.DB.Exec("DELETE FROM tournaments WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tournament"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tournament deleted"})
}
