package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetStatsOverview
func GetStatsOverview(c *gin.Context) {
	var studentCount int64
	models.DB.Model(&models.Student{}).Where("status = ?", "active").Count(&studentCount)

	var classCount int64
	models.DB.Model(&models.Class{}).Count(&classCount)

	var pendingTrials int64
	models.DB.Model(&models.TrialRegistration{}).Where("status = ?", "pending").Count(&pendingTrials)

	var coachCount int64
	models.DB.Model(&models.Coach{}).Where("is_active = ?", true).Count(&coachCount)

	var totalRevenue float64
	models.DB.Model(&models.Invoice{}).Where("status = ?", "paid").Select("COALESCE(SUM(amount), 0)").Row().Scan(&totalRevenue)

	c.JSON(http.StatusOK, gin.H{
		"studentCount":  studentCount,
		"classCount":    classCount,
		"pendingTrials": pendingTrials,
		"coachCount":    coachCount,
		"totalRevenue":  totalRevenue,
	})
}

// GetStatsEnrollment - Line/Bar chart (enrollment trend by month)
func GetStatsEnrollment(c *gin.Context) {
	type MonthData struct {
		Name       string `json:"name"`
		NewStudents int    `json:"Học viên mới"`
		Trials     int    `json:"Học thử"`
	}

	now := time.Now()
	var data []MonthData

	for i := 5; i >= 0; i-- {
		month := now.AddDate(0, -i, 0)
		start := time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, time.Local)
		end := start.AddDate(0, 1, 0)

		var newStudents int64
		models.DB.Model(&models.Student{}).
			Where("join_date >= ? AND join_date < ? AND status = ?", start, end, "active").
			Count(&newStudents)

		var trials int64
		models.DB.Model(&models.TrialRegistration{}).
			Where("created_at >= ? AND created_at < ?", start, end).
			Count(&trials)

		data = append(data, MonthData{
			Name:        month.Format("Tháng 1"),
			NewStudents: int(newStudents),
			Trials:      int(trials),
		})
	}

	if data == nil {
		data = []MonthData{}
	}
	c.JSON(http.StatusOK, data)
}

// GetStatsDistribution - Pie chart (students per court)
func GetStatsDistribution(c *gin.Context) {
	type CourtData struct {
		Name       string `json:"name"`
		ClassCount int    `json:"Số lớp"`
	}

	var courts []models.Court
	models.DB.Find(&courts)

	var data []CourtData
	for _, court := range courts {
		var classCount int64
		models.DB.Model(&models.Class{}).Where("court_id = ?", court.ID).Count(&classCount)
		data = append(data, CourtData{
			Name:       court.Name,
			ClassCount: int(classCount),
		})
	}

	if data == nil {
		data = []CourtData{}
	}
	c.JSON(http.StatusOK, data)
}

// GetStatsCoach - Bar chart (attendance rate per coach)
func GetStatsCoach(c *gin.Context) {
	type CoachStats struct {
		Name           string `json:"name"`
		AttendanceRate int    `json:"Chuyên cần (%)"`
	}

	var coaches []models.Coach
	models.DB.Where("is_active = ?", true).Find(&coaches)

	var data []CoachStats
	for _, coach := range coaches {
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

		rate := 0
		if totalAtt > 0 {
			rate = int(presentAtt * 100 / totalAtt)
		}

		data = append(data, CoachStats{
			Name:           coach.Name,
			AttendanceRate: rate,
		})
	}

	if data == nil {
		data = []CoachStats{}
	}
	c.JSON(http.StatusOK, data)
}
