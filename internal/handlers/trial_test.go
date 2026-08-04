package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"backend-ocean-basketball/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSubmitTrialRegistration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := models.MockDB()

	r := gin.Default()
	r.POST("/trial", SubmitTrialRegistration)

	t.Run("success", func(t *testing.T) {
		reqBody := map[string]interface{}{
			"student_name": "John",
			"parent_name":  "Jane",
			"phone":        "123456789",
		}
		body, _ := json.Marshal(reqBody)

		mock.ExpectBegin()
		mock.ExpectExec(`^INSERT INTO "trial_registrations"`).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		req, _ := http.NewRequest(http.MethodPost, "/trial", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestGetTrials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := models.MockDB()

	r := gin.Default()
	r.GET("/admin/trials", GetTrials)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "student_name", "parent_name", "status", "created_at"}).
			AddRow("uuid-1", "John", "Jane", "pending", time.Now())

		mock.ExpectQuery(`^SELECT \* FROM "trial_registrations" ORDER BY created_at desc`).WillReturnRows(rows)

		req, _ := http.NewRequest(http.MethodGet, "/admin/trials", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestUpdateTrialStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mock := models.MockDB()

	r := gin.Default()
	r.PUT("/admin/trials/:id/status", UpdateTrialStatus)

	t.Run("success", func(t *testing.T) {
		reqBody := map[string]string{"status": "contacted"}
		body, _ := json.Marshal(reqBody)

		rows := sqlmock.NewRows([]string{"id", "status"}).AddRow("uuid-1", "pending")
		mock.ExpectQuery(`^SELECT \* FROM "trial_registrations" WHERE id = \$1 ORDER BY "trial_registrations"\."id" LIMIT \$2`).WithArgs("uuid-1", 1).WillReturnRows(rows)

		mock.ExpectBegin()
		mock.ExpectExec(`^UPDATE "trial_registrations" SET`).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		req, _ := http.NewRequest(http.MethodPut, "/admin/trials/uuid-1/status", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
