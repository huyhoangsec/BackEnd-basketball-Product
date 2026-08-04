package handlers

import (
	"backend-ocean-basketball/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTournament(c *gin.Context) {
	tournamentID := c.Param("id")

	var req struct {
		TeamName    string `json:"team_name" binding:"required"`
		ContactName string `json:"contact_name" binding:"required"`
		ContactPhone string `json:"contact_phone" binding:"required"`
		Players     int    `json:"players" binding:"required,min=1"`
		Notes       string `json:"notes"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	var tournament models.Tournament
	if err := models.DB.First(&tournament, "id = ?", tournamentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tournament not found"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Đăng ký giải đấu thành công",
		"tournament":   tournament.Name,
		"team_name":    req.TeamName,
		"contact_name": req.ContactName,
	})
}

type Match struct {
	ID          string `json:"id"`
	TournamentID string `json:"tournament_id"`
	Round       int    `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Score1      *int   `json:"score1"`
	Score2      *int   `json:"score2"`
	Status      string `json:"status"` // "scheduled", "in_progress", "completed"
	Date        string `json:"date"`
}

func GetTournamentMatches(c *gin.Context) {
	tournamentID := c.Param("id")

	var tournament models.Tournament
	if err := models.DB.First(&tournament, "id = ?", tournamentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tournament not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tournament": tournament,
		"matches":    []Match{},
	})
}
