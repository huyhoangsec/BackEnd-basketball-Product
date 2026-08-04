package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudentProfile(c *gin.Context) {
	studentID := c.Param("id")

	var student models.Student
	if err := models.DB.Preload("Class").First(&student, "id = ?", studentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	var attendanceCount int64
	models.DB.Model(&models.AttendanceRecord{}).
		Where("student_id = ? AND status = ?", studentID, "present").
		Count(&attendanceCount)

	var totalClasses int64
	models.DB.Model(&models.AttendanceRecord{}).
		Where("student_id = ?", studentID).
		Count(&totalClasses)

	attendanceRate := 0
	if totalClasses > 0 {
		attendanceRate = int(attendanceCount * 100 / totalClasses)
	}

	c.JSON(http.StatusOK, gin.H{
		"student":         student,
		"attendance_rate": attendanceRate,
		"total_classes":   totalClasses,
	})
}

func TransferStudent(c *gin.Context) {
	studentID := c.Param("id")

	var req struct {
		NewClassID string `json:"new_class_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var student models.Student
	if err := models.DB.First(&student, "id = ?", studentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	var targetClass models.Class
	if err := models.DB.First(&targetClass, "id = ?", req.NewClassID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target class not found"})
		return
	}

	student.ClassID = req.NewClassID
	if err := models.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transfer student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student transferred successfully", "student": student})
}

func GetStudentAttendanceHistory(c *gin.Context) {
	studentID := c.Param("id")

	var records []models.AttendanceRecord
	if err := models.DB.
		Where("student_id = ?", studentID).
		Order("date desc").
		Preload("Student").
		Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance history"})
		return
	}

	c.JSON(http.StatusOK, records)
}
