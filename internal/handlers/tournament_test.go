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

func TestGetTournaments(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow("1", "Summer Cup", "upcoming")
	
	mock.ExpectQuery(`^SELECT \* FROM "tournaments"`).
		WillReturnRows(rows)

	router := setupTestRouter()
	router.GET("/api/admin/tournaments", GetTournaments)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/tournaments", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Tournament
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Nil(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "Summer Cup", response[0].Name)
}
