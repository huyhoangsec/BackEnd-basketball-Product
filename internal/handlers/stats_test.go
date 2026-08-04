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

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestGetStatsEnrollment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := models.MockDB()

	for i := 0; i < 12; i++ {
		mock.ExpectQuery(`^SELECT count\(\*\) FROM "students"`).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(10))
		mock.ExpectQuery(`^SELECT count\(\*\) FROM "trial_registrations"`).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))
	}

	router := setupTestRouter()
	router.GET("/api/admin/stats/enrollment", GetStatsEnrollment)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/stats/enrollment", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}
