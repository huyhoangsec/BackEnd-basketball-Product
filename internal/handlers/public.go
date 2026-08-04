package handlers

import (
	"net/http"

	"backend-ocean-basketball/internal/models"
	"github.com/gin-gonic/gin"
)

// Get all active coaches
func GetCoaches(c *gin.Context) {
	var coaches []models.Coach
	if err := models.DB.Where("is_active = ?", true).Find(&coaches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch coaches"})
		return
	}
	c.JSON(http.StatusOK, coaches)
}

// Get all courts
func GetCourts(c *gin.Context) {
	var courts []models.Court
	if err := models.DB.Find(&courts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courts"})
		return
	}
	c.JSON(http.StatusOK, courts)
}

// Get classes with schedules
func GetClasses(c *gin.Context) {
	var classes []models.Class
	if err := models.DB.Preload("Court").Preload("Coach").Preload("Schedules").Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch classes"})
		return
	}
	c.JSON(http.StatusOK, classes)
}

// Get tuition plans
func GetTuitionPlans(c *gin.Context) {
	var plans []models.TuitionPlan
	if err := models.DB.Find(&plans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tuition plans"})
		return
	}
	c.JSON(http.StatusOK, plans)
}

// Get FAQs
func GetFAQs(c *gin.Context) {
	var faqs []models.FAQ
	if err := models.DB.Order("\"order\" asc").Find(&faqs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FAQs"})
		return
	}
	c.JSON(http.StatusOK, faqs)
}

// Get visible Reviews
func GetReviews(c *gin.Context) {
	var reviews []models.Review
	if err := models.DB.Where("is_visible = ?", true).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// Get active banners
func GetBanners(c *gin.Context) {
	var banners []models.Banner
	if err := models.DB.Where("is_active = ?", true).Order("\"order\" asc").Find(&banners).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banners"})
		return
	}
	c.JSON(http.StatusOK, banners)
}
