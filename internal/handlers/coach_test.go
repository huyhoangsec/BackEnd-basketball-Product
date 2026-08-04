package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-ocean-basketball/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCoachClasses(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "name", "coach_id", "court_id"}).
		AddRow("cls-1", "Basic Class", "c-123", "court-1")
	
	mock.ExpectQuery(`^SELECT \* FROM "classes" WHERE coach_id = \$1`).
		WithArgs("c-123").
		WillReturnRows(rows)

	mock.ExpectQuery(`^SELECT \* FROM "courts" WHERE "courts"\."id" = \$1`).
		WithArgs("court-1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("court-1", "Test Court"))

	mock.ExpectQuery(`^SELECT \* FROM "class_schedules" WHERE "class_schedules"\."class_id" = \$1`).
		WithArgs("cls-1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "class_id", "day_of_week"}))

	router := setupTestRouter()
	
	router.Use(func(c *gin.Context) {
		c.Set("userId", "c-123")
		c.Set("role", "coach")
		c.Next()
	})
	router.GET("/api/coach/classes", GetCoachClasses)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/coach/classes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Class
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Nil(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "Basic Class", response[0].Name)
}
