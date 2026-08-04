package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-ocean-basketball/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCoach(t *testing.T) {
	mock := models.MockDB()
	
	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "coaches"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := setupTestRouter()
	router.POST("/api/admin/coaches", CreateCoach)

	body := map[string]string{
		"name": "New Coach",
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/admin/coaches", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
