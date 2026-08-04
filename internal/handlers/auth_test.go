package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-ocean-basketball/config"
	"backend-ocean-basketball/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLoginInvalidCredentials(t *testing.T) {
	mock := models.MockDB()
	
	// Mock the user query to return empty (not found)
	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE email = \$1`).
		WithArgs("wrong@email.com", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	router := setupTestRouter()
	cfg := &config.Config{JWTSecret: "test_secret"}
	authHandler := NewAuthHandler(cfg)

	router.POST("/api/login", authHandler.Login)

	body := map[string]string{
		"email":    "wrong@email.com",
		"password": "wrongpassword",
	}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
