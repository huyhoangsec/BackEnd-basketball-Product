package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOverdueInvoices(c *gin.Context) {
	var invoices []models.Invoice
	if err := models.DB.
		Where("status = ? AND due_date < ?", "unpaid", time.Now()).
		Preload("Student").
		Order("due_date asc").
		Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch overdue invoices"})
		return
	}

	// Auto-mark as overdue
	for i := range invoices {
		if invoices[i].Status == "unpaid" && invoices[i].DueDate.Before(time.Now()) {
			invoices[i].Status = "overdue"
			models.DB.Save(&invoices[i])
		}
	}

	c.JSON(http.StatusOK, invoices)
}

func GetInvoiceStats(c *gin.Context) {
	var totalPaid int64
	models.DB.Model(&models.Invoice{}).Where("status = ?", "paid").Count(&totalPaid)

	var totalUnpaid int64
	models.DB.Model(&models.Invoice{}).Where("status = ?", "unpaid").Count(&totalUnpaid)

	var totalOverdue int64
	models.DB.Model(&models.Invoice{}).Where("status = ? AND due_date < ?", "unpaid", time.Now()).Count(&totalOverdue)

	var totalRevenue float64
	models.DB.Model(&models.Invoice{}).Where("status = ?", "paid").Select("COALESCE(SUM(amount), 0)").Row().Scan(&totalRevenue)

	var pendingRevenue float64
	models.DB.Model(&models.Invoice{}).Where("status IN ?", []string{"unpaid", "overdue"}).Select("COALESCE(SUM(amount), 0)").Row().Scan(&pendingRevenue)

	c.JSON(http.StatusOK, gin.H{
		"total_paid":      totalPaid,
		"total_unpaid":    totalUnpaid,
		"total_overdue":   totalOverdue,
		"total_revenue":   totalRevenue,
		"pending_revenue": pendingRevenue,
	})
}
