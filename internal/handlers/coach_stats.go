package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCoachWorkload(c *gin.Context) {
	coachID := c.Param("id")

	var coach models.Coach
	if err := models.DB.First(&coach, "id = ?", coachID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coach not found"})
		return
	}

	var classCount int64
	models.DB.Model(&models.Class{}).Where("coach_id = ?", coachID).Count(&classCount)

	var studentCount int64
	models.DB.Model(&models.Student{}).
		Joins("JOIN classes ON classes.id = students.class_id").
		Where("classes.coach_id = ? AND students.status = ?", coachID, "active").
		Count(&studentCount)

	var attendanceCount int64
	models.DB.Model(&models.AttendanceRecord{}).
		Joins("JOIN classes ON classes.id = attendance_records.class_id").
		Where("classes.coach_id = ?", coachID).
		Count(&attendanceCount)

	var presentCount int64
	models.DB.Model(&models.AttendanceRecord{}).
		Joins("JOIN classes ON classes.id = attendance_records.class_id").
		Where("classes.coach_id = ? AND attendance_records.status = ?", coachID, "present").
		Count(&presentCount)

	attendanceRate := 0
	if attendanceCount > 0 {
		attendanceRate = int(presentCount * 100 / attendanceCount)
	}

	c.JSON(http.StatusOK, gin.H{
		"coach":           coach,
		"class_count":     classCount,
		"student_count":   studentCount,
		"attendance_rate": attendanceRate,
	})
}

func GetAllCoachesStats(c *gin.Context) {
	var coaches []models.Coach
	if err := models.DB.Where("is_active = ?", true).Find(&coaches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch coaches"})
		return
	}

	type CoachStat struct {
		ID             string  `json:"id"`
		Name           string  `json:"name"`
		ClassCount     int64   `json:"class_count"`
		StudentCount   int64   `json:"student_count"`
		AttendanceRate float64 `json:"attendance_rate"`
	}

	var stats []CoachStat
	for _, coach := range coaches {
		var classCount int64
		models.DB.Model(&models.Class{}).Where("coach_id = ?", coach.ID).Count(&classCount)

		var studentCount int64
		models.DB.Model(&models.Student{}).
			Joins("JOIN classes ON classes.id = students.class_id").
			Where("classes.coach_id = ? AND students.status = ?", coach.ID, "active").
			Count(&studentCount)

		var totalAtt int64
		models.DB.Model(&models.AttendanceRecord{}).
			Joins("JOIN classes ON classes.id = attendance_records.class_id").
			Where("classes.coach_id = ?", coach.ID).
			Count(&totalAtt)

		var presentAtt int64
		models.DB.Model(&models.AttendanceRecord{}).
			Joins("JOIN classes ON classes.id = attendance_records.class_id").
			Where("classes.coach_id = ? AND attendance_records.status = ?", coach.ID, "present").
			Count(&presentAtt)

		rate := 0.0
		if totalAtt > 0 {
			rate = float64(presentAtt) / float64(totalAtt) * 100
		}

		stats = append(stats, CoachStat{
			ID:             coach.ID,
			Name:           coach.Name,
			ClassCount:     classCount,
			StudentCount:   studentCount,
			AttendanceRate: rate,
		})
	}

	if stats == nil {
		stats = []CoachStat{}
	}

	c.JSON(http.StatusOK, stats)
}
