package handlers

import (
	"fmt"
	"net/http"
	"time"

	"backend-ocean-basketball/internal/models"

	"github.com/gin-gonic/gin"
)

// --- Banners ---

func GetAllBanners(c *gin.Context) {
	var banners []models.Banner
	if err := models.DB.Order("\"order\" asc").Find(&banners).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banners"})
		return
	}
	c.JSON(http.StatusOK, banners)
}

func CreateBanner(c *gin.Context) {
	var banner models.Banner
	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if banner.ID == "" {
		banner.ID = fmt.Sprintf("b-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&banner).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create banner"})
		return
	}
	c.JSON(http.StatusCreated, banner)
}

func UpdateBanner(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Banner ID required"})
		return
	}
	var banner models.Banner
	if err := models.DB.First(&banner, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banner not found"})
		return
	}
	if err := c.ShouldBindJSON(&banner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	banner.ID = id
	if err := models.DB.Save(&banner).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update banner"})
		return
	}
	c.JSON(http.StatusOK, banner)
}

func DeleteBanner(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid banner ID"})
		return
	}
	if err := models.DB.Exec("DELETE FROM banners WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete banner"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Banner deleted"})
}

// --- FAQs ---

func GetAllFAQs(c *gin.Context) {
	var faqs []models.FAQ
	if err := models.DB.Order("\"order\" asc").Find(&faqs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FAQs"})
		return
	}
	c.JSON(http.StatusOK, faqs)
}

func CreateFAQ(c *gin.Context) {
	var item models.FAQ
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if item.ID == "" {
		item.ID = fmt.Sprintf("faq-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create FAQ"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func UpdateFAQ(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FAQ ID required"})
		return
	}
	var item models.FAQ
	if err := models.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	if err := models.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update FAQ"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func DeleteFAQ(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FAQ ID"})
		return
	}
	if err := models.DB.Exec("DELETE FROM faqs WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete FAQ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "FAQ deleted"})
}

// --- Reviews ---

func GetAllReviews(c *gin.Context) {
	var reviews []models.Review
	if err := models.DB.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func CreateReview(c *gin.Context) {
	var item models.Review
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if item.ID == "" {
		item.ID = fmt.Sprintf("rev-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Review ID required"})
		return
	}
	var item models.Review
	if err := models.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	if err := models.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	if err := models.DB.Exec("DELETE FROM reviews WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
}

// --- TuitionPlans ---

func GetAllTuitionPlans(c *gin.Context) {
	var plans []models.TuitionPlan
	if err := models.DB.Find(&plans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tuition plans"})
		return
	}
	c.JSON(http.StatusOK, plans)
}

func CreateTuitionPlan(c *gin.Context) {
	var item models.TuitionPlan
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if item.ID == "" {
		item.ID = fmt.Sprintf("tp-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tuition plan"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func UpdateTuitionPlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tuition plan ID required"})
		return
	}
	var item models.TuitionPlan
	if err := models.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tuition plan not found"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	if err := models.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tuition plan"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func DeleteTuitionPlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tuition plan ID"})
		return
	}
	if err := models.DB.Exec("DELETE FROM tuition_plans WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tuition plan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tuition plan deleted"})
}
