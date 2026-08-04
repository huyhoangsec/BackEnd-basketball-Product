package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetInvoices retrieves all invoices (with optional filters)
func GetInvoices(c *gin.Context) {
	var invoices []models.Invoice
	query := models.DB.Preload("Student").Order("created_at desc")

	status := c.Query("status")
	if status != "" {
		query = query.Where("status = ?", status)
	}

	month := c.Query("month")
	if month != "" {
		query = query.Where("month = ?", month)
	}

	year := c.Query("year")
	if year != "" {
		query = query.Where("year = ?", year)
	}

	if err := query.Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

// CreateInvoice creates a new invoice
func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if invoice.Month == 0 {
		invoice.Month = int(time.Now().Month())
	}
	if invoice.Year == 0 {
		invoice.Year = time.Now().Year()
	}
	if invoice.DueDate.IsZero() {
		due := time.Date(invoice.Year, time.Month(invoice.Month+1), 5, 0, 0, 0, 0, time.Local)
		invoice.DueDate = due
	}

	invoice.Status = "unpaid"
	if err := models.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

// PayInvoice records a payment for an invoice and updates its status
func PayInvoice(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invoice ID required"})
		return
	}

	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var invoice models.Invoice
	if err := models.DB.First(&invoice, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	tx := models.DB.Begin()

	payment.InvoiceID = invoice.ID
	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record payment"})
		return
	}

	invoice.Status = "paid"
	if err := tx.Save(&invoice).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice status"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Payment recorded successfully", "invoice": invoice, "payment": payment})
}

// DeleteInvoice deletes an invoice and its associated payments safely
func DeleteInvoice(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	models.DB.Where("invoice_id = ?", id).Delete(&models.Payment{})
	if err := models.DB.Delete(&models.Invoice{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invoice"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted"})
}

// GenerateMonthlyInvoices generates invoices for all active students for the current month
func GenerateMonthlyInvoices(c *gin.Context) {
	now := time.Now()
	month := int(now.Month())
	year := now.Year()

	var students []models.Student
	if err := models.DB.Where("status = ?", "active").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active students"})
		return
	}

	createdCount := 0
	for _, student := range students {
		amount := 1000000.0

		var existingInvoice models.Invoice
		err := models.DB.Where("student_id = ? AND month = ? AND year = ? AND status = ?",
			student.ID, month, year, "unpaid").First(&existingInvoice).Error
		if err == nil {
			continue
		}

		dueDate := time.Date(year, time.Month(month+1), 5, 0, 0, 0, 0, time.Local)
		invoice := models.Invoice{
			StudentID: student.ID,
			Amount:    amount,
			Month:     month,
			Year:      year,
			DueDate:   dueDate,
			Status:    "unpaid",
		}

		if err := models.DB.Create(&invoice).Error; err == nil {
			createdCount++
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoices generated", "count": createdCount, "month": month, "year": year})
}
