package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	contact.IsRead = false
	if err := models.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit contact"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Contact submitted successfully"})
}

func GetContacts(c *gin.Context) {
	var contacts []models.Contact
	if err := models.DB.Order("created_at desc").Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}
	c.JSON(http.StatusOK, contacts)
}

func MarkContactRead(c *gin.Context) {
	id := c.Param("id")
	var contact models.Contact
	if err := models.DB.First(&contact, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}
	contact.IsRead = true
	if err := models.DB.Save(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark as read"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact marked as read"})
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Contact{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted"})
}
