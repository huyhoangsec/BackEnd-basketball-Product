package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAttendanceReport(c *gin.Context) {
	classID := c.Query("class_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	query := models.DB.Model(&models.AttendanceRecord{}).Preload("Student")

	if classID != "" {
		query = query.Where("class_id = ?", classID)
	}

	if startDate != "" {
		if t, err := time.Parse("2006-01-02", startDate); err == nil {
			query = query.Where("date >= ?", t)
		}
	}

	if endDate != "" {
		if t, err := time.Parse("2006-01-02", endDate); err == nil {
			query = query.Where("date <= ?", t)
		}
	}

	var records []models.AttendanceRecord
	if err := query.Order("date desc").Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance report"})
		return
	}

	type StudentAttendance struct {
		StudentID   string `json:"student_id"`
		StudentName string `json:"student_name"`
		Present     int    `json:"present"`
		Absent      int    `json:"absent"`
		Excused     int    `json:"excused"`
		Unexcused   int    `json:"unexcused"`
		Total       int    `json:"total"`
		Rate        int    `json:"rate"`
	}

	studentMap := make(map[string]*StudentAttendance)

	for _, r := range records {
		sid := r.StudentID
		if _, ok := studentMap[sid]; !ok {
			studentMap[sid] = &StudentAttendance{
				StudentID:   sid,
				StudentName: r.Student.Name,
			}
		}
		sa := studentMap[sid]
		sa.Total++
		switch r.Status {
		case "present":
			sa.Present++
		case "cancelled":
			sa.Total--
		case "excused":
			sa.Excused++
		case "unexcused":
			sa.Unexcused++
		case "dropped":
			sa.Total--
		}
	}

	var result []StudentAttendance
	for _, sa := range studentMap {
		if sa.Total > 0 {
			sa.Rate = sa.Present * 100 / sa.Total
		}
		result = append(result, *sa)
	}

	if result == nil {
		result = []StudentAttendance{}
	}

	c.JSON(http.StatusOK, gin.H{
		"records":  records,
		"summary":  result,
		"total":    len(records),
	})
}

func GetAttendanceByDate(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		dateStr = time.Now().Format("2006-01-02")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	var records []models.AttendanceRecord
	if err := models.DB.
		Where("date = ?", date).
		Preload("Student").
		Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance"})
		return
	}

	c.JSON(http.StatusOK, records)
}
