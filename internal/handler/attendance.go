package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type AttendanceRecord struct {
	StudentID uint   `json:"student_id" binding:"required"`
	Status    string `json:"status" binding:"required"` // present, absent, late
	Note      string `json:"note"`
}

type BatchAttendanceRequest struct {
	ClassID uint               `json:"class_id" binding:"required"`
	Date    string             `json:"date" binding:"required"`
	Records []AttendanceRecord `json:"records" binding:"required"`
}

func SubmitAttendance(c *gin.Context) {
	var req BatchAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid request body", "error": err.Error()})
		return
	}

	// Save logic placeholder
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Điểm danh thành công",
		"data": gin.H{
			"total_processed": len(req.Records),
			"class_id":        req.ClassID,
		},
	})
}