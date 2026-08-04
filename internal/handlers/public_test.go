package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-ocean-basketball/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCoaches(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "name", "is_active"}).
		AddRow("1", "Coach Test", true)
	
	mock.ExpectQuery(`^SELECT \* FROM "coaches" WHERE is_active = \$1`).
		WithArgs(true).
		WillReturnRows(rows)

	router := setupTestRouter()
	router.GET("/api/public/coaches", GetCoaches)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/public/coaches", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Coach
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Nil(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "Coach Test", response[0].Name)
}

func TestGetCourts(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1", "Test Court")
	
	mock.ExpectQuery(`^SELECT \* FROM "courts"`).
		WillReturnRows(rows)

	router := setupTestRouter()
	router.GET("/api/public/courts", GetCourts)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/public/courts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Court
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Nil(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "Test Court", response[0].Name)
}
