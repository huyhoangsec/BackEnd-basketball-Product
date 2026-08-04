package handlers

import (
	"net/http"

	"backend-ocean-basketball/internal/models"
	"github.com/gin-gonic/gin"
)

// Public: Submit a new trial registration
func SubmitTrialRegistration(c *gin.Context) {
	var trial models.TrialRegistration
	if err := c.ShouldBindJSON(&trial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	trial.Status = "pending"
	
	if err := models.DB.Create(&trial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit trial registration"})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Trial registered successfully", "trial": trial})
}

// Admin: Get all trial registrations
func GetTrials(c *gin.Context) {
	var trials []models.TrialRegistration
	if err := models.DB.Order("created_at desc").Find(&trials).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trial registrations"})
		return
	}
	c.JSON(http.StatusOK, trials)
}

// Admin: Update trial status
func UpdateTrialStatus(c *gin.Context) {
	id := c.Param("id")
	
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var trial models.TrialRegistration
	if err := models.DB.First(&trial, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trial registration not found"})
		return
	}

	trial.Status = req.Status
	if err := models.DB.Save(&trial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update trial status"})
		return
	}

	c.JSON(http.StatusOK, trial)
}
