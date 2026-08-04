package handlers

import (
	"fmt"
	"net/http"
	"time"

	"backend-ocean-basketball/internal/models"
	"github.com/gin-gonic/gin"
)

// -- Coaches --

func GetAllCoaches(c *gin.Context) {
	var coaches []models.Coach
	if err := models.DB.Find(&coaches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch coaches"})
		return
	}
	c.JSON(http.StatusOK, coaches)
}

func CreateCoach(c *gin.Context) {
	var coach models.Coach
	if err := c.ShouldBindJSON(&coach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	if coach.ID == "" {
		coach.ID = fmt.Sprintf("c-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&coach).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create coach"})
		return
	}
	c.JSON(http.StatusCreated, coach)
}

func UpdateCoach(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Coach ID required"})
		return
	}
	var coach models.Coach
	if err := models.DB.First(&coach, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coach not found"})
		return
	}

	if err := c.ShouldBindJSON(&coach); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	coach.ID = id // Keep original ID

	if err := models.DB.Save(&coach).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update coach"})
		return
	}
	c.JSON(http.StatusOK, coach)
}

// -- Courts --

func CreateCourt(c *gin.Context) {
	var court models.Court
	if err := c.ShouldBindJSON(&court); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	if court.ID == "" {
		court.ID = fmt.Sprintf("court-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&court).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create court"})
		return
	}
	c.JSON(http.StatusCreated, court)
}

func UpdateCourt(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Court ID required"})
		return
	}
	var court models.Court
	if err := models.DB.First(&court, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Court not found"})
		return
	}

	if err := c.ShouldBindJSON(&court); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	court.ID = id // Keep original ID

	if err := models.DB.Save(&court).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update court"})
		return
	}
	c.JSON(http.StatusOK, court)
}

// -- Classes --

func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	if class.ID == "" {
		class.ID = fmt.Sprintf("cls-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		return
	}
	c.JSON(http.StatusCreated, class)
}

func UpdateClass(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Class ID required"})
		return
	}
	var class models.Class
	if err := models.DB.First(&class, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	class.ID = id // Keep original ID

	if err := models.DB.Save(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update class"})
		return
	}
	c.JSON(http.StatusOK, class)
}

// -- DELETE handlers --

func DeleteCoach(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid coach ID"})
		return
	}

	// Set SQL NULL for foreign key in classes instead of empty string
	models.DB.Exec("UPDATE classes SET coach_id = NULL WHERE coach_id = ?", id)

	if err := models.DB.Exec("DELETE FROM coaches WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete coach"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Coach deleted"})
}

func DeleteCourt(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid court ID"})
		return
	}

	// Set SQL NULL for foreign key in classes instead of empty string
	models.DB.Exec("UPDATE classes SET court_id = NULL WHERE court_id = ?", id)

	if err := models.DB.Exec("DELETE FROM courts WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete court"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Court deleted"})
}

func DeleteClass(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	// Safely cleanup class schedules, attendance records & set SQL NULL for students
	models.DB.Exec("DELETE FROM class_schedules WHERE class_id = ?", id)
	models.DB.Exec("DELETE FROM attendance_records WHERE class_id = ?", id)
	models.DB.Exec("UPDATE students SET class_id = NULL WHERE class_id = ?", id)

	if err := models.DB.Exec("DELETE FROM classes WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete class"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Class deleted"})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	// Safely cleanup student attendance, payments & invoices
	models.DB.Exec("DELETE FROM attendance_records WHERE student_id = ?", id)
	models.DB.Exec("DELETE FROM payments WHERE invoice_id IN (SELECT id FROM invoices WHERE student_id = ?)", id)
	models.DB.Exec("DELETE FROM invoices WHERE student_id = ?", id)

	if err := models.DB.Exec("DELETE FROM students WHERE id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

// -- Students --

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	if student.ID == "" {
		student.ID = fmt.Sprintf("st-%d", time.Now().UnixNano())
	}
	if err := models.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func GetStudents(c *gin.Context) {
	var students []models.Student
	if err := models.DB.Preload("Class").Order("created_at desc").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	for i := range students {
		if students[i].Class.Name != "" {
			students[i].ClassName = students[i].Class.Name
		}
	}
	c.JSON(http.StatusOK, students)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student ID required"})
		return
	}
	var student models.Student
	if err := models.DB.First(&student, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}
	student.ID = id // Keep original ID

	if err := models.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}
	c.JSON(http.StatusOK, student)
}
