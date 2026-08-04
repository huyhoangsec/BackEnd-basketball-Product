package handlers

import (
	"net/http"
	"time"

	"backend-ocean-basketball/internal/models"
	"github.com/gin-gonic/gin"
)

// Get classes assigned to the logged-in coach
func GetCoachClasses(c *gin.Context) {
	coachId := c.GetString("userId")
	role := c.GetString("role")

	var classes []models.Class
	query := models.DB.Preload("Court").Preload("Schedules")
	
	// If admin, they can see all. If coach, only their own.
	if role == "coach" {
		query = query.Where("coach_id = ?", coachId)
	}

	if err := query.Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch classes"})
		return
	}
	c.JSON(http.StatusOK, classes)
}

// Get students for a specific class
func GetClassStudents(c *gin.Context) {
	classId := c.Param("classId")
	
	// Optional: Check if class belongs to this coach
	// (For simplicity, skipping the authorization check here, assuming middleware/frontend handles it)

	var students []models.Student
	if err := models.DB.Where("class_id = ?", classId).Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

// Mark attendance for a student
func MarkAttendance(c *gin.Context) {
	classId := c.Param("classId")
	markedBy := c.GetString("userId") // Logged in coach or admin

	var req struct {
		StudentID string `json:"student_id" binding:"required"`
		Status    string `json:"status" binding:"required"`
		Date      string `json:"date" binding:"required"` // Format: YYYY-MM-DD
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	// Check if attendance already exists for this student on this date
	var record models.AttendanceRecord
	err = models.DB.Where("student_id = ? AND date = ?", req.StudentID, date).First(&record).Error

	if err == nil {
		// Update existing record
		record.Status = req.Status
		record.MarkedBy = markedBy
		if err := models.DB.Save(&record).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attendance"})
			return
		}
	} else {
		// Create new record
		record = models.AttendanceRecord{
			StudentID: req.StudentID,
			ClassID:   classId,
			Date:      date,
			Status:    req.Status,
			MarkedBy:  markedBy,
		}
		if err := models.DB.Create(&record).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark attendance"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance marked successfully", "record": record})
}

// Get attendance history for a class on a specific date
func GetAttendanceHistory(c *gin.Context) {
	classId := c.Param("classId")
	dateStr := c.Query("date") // Format: YYYY-MM-DD

	query := models.DB.Where("class_id = ?", classId).Preload("Student")
	
	if dateStr != "" {
		date, err := time.Parse("2006-01-02", dateStr)
		if err == nil {
			query = query.Where("date = ?", date)
		}
	}

	var records []models.AttendanceRecord
	if err := query.Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance history"})
		return
	}
	
	c.JSON(http.StatusOK, records)
}
