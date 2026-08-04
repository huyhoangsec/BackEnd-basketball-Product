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

func TestGetInvoices(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "student_id", "amount", "status"}).
		AddRow("1", "s-123", 500000, "paid")
	
	mock.ExpectQuery(`^SELECT \* FROM "invoices"`).
		WillReturnRows(rows)

	mock.ExpectQuery(`^SELECT \* FROM "students" WHERE "students"\."id" = \$1`).
		WithArgs("s-123").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow("s-123", "John Doe"))

	router := setupTestRouter()
	router.GET("/api/admin/invoices", GetInvoices)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/invoices", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Invoice
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Nil(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "s-123", response[0].StudentID)
	assert.Equal(t, "paid", response[0].Status)
}

func TestCreateInvoice(t *testing.T) {
	mock := models.MockDB()

	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "invoices"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := setupTestRouter()
	router.POST("/api/admin/invoices", CreateInvoice)

	payload := map[string]interface{}{
		"student_id": "s-123",
		"amount":     500000,
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/admin/invoices", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPayInvoice(t *testing.T) {
	mock := models.MockDB()

	// SELECT invoice
	mock.ExpectQuery(`^SELECT \* FROM "invoices" WHERE id = \$1 ORDER BY "invoices"\."id" LIMIT \$2`).
		WithArgs("inv-1", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "status"}).AddRow("inv-1", "unpaid"))

	mock.ExpectBegin()
	
	// INSERT payment
	mock.ExpectExec(`^INSERT INTO "payments"`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// UPDATE invoice
	mock.ExpectExec(`^UPDATE "invoices" SET`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	router := setupTestRouter()
	router.PUT("/api/admin/invoices/:id/pay", PayInvoice)

	payload := map[string]interface{}{
		"amount": 500000,
		"method": "cash",
		"note":   "Test note",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/admin/invoices/inv-1/pay", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGenerateMonthlyInvoices(t *testing.T) {
	mock := models.MockDB()

	mock.ExpectQuery(`^SELECT \* FROM "students" WHERE status = \$1`).
		WithArgs("active").
		WillReturnRows(sqlmock.NewRows([]string{"id", "status"}).AddRow("s-123", "active"))

	mock.ExpectQuery(`^SELECT \* FROM "invoices"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"})) // Not found -> triggers create

	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "invoices"`).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := setupTestRouter()
	router.POST("/api/admin/invoices/generate", GenerateMonthlyInvoices)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/admin/invoices/generate", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
