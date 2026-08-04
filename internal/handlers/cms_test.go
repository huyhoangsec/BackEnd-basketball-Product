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

// --- Banners ---

func TestGetAllBanners(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "title", "order"}).
		AddRow("1", "Banner 1", 1)
	
	mock.ExpectQuery(`^SELECT \* FROM "banners"`).
		WillReturnRows(rows)

	router := setupTestRouter()
	router.GET("/admin/banners", GetAllBanners)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/banners", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Banner
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Nil(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "Banner 1", response[0].Title)
}

func TestCreateBanner(t *testing.T) {
	mock := models.MockDB()
	
	banner := models.Banner{Title: "New Banner", Subtitle: "Sub", CTAText: "Click"}
	body, _ := json.Marshal(banner)

	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "banners"`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := setupTestRouter()
	router.POST("/admin/banners", CreateBanner)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/banners", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// --- FAQs ---

func TestGetAllFAQs(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "question"}).
		AddRow("1", "Q1")
	
	mock.ExpectQuery(`^SELECT \* FROM "faqs"`).
		WillReturnRows(rows)

	router := setupTestRouter()
	router.GET("/admin/faqs", GetAllFAQs)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/faqs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.FAQ
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Nil(t, err)
	assert.Len(t, response, 1)
	assert.Equal(t, "Q1", response[0].Question)
}

func TestCreateFAQ(t *testing.T) {
	mock := models.MockDB()
	
	faq := models.FAQ{Question: "Q", Answer: "A"}
	body, _ := json.Marshal(faq)

	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "faqs"`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := setupTestRouter()
	router.POST("/admin/faqs", CreateFAQ)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/faqs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// --- Reviews ---

func TestGetAllReviews(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "content"}).
		AddRow("1", "Good")
	
	mock.ExpectQuery(`^SELECT \* FROM "reviews"`).
		WillReturnRows(rows)

	router := setupTestRouter()
	router.GET("/admin/reviews", GetAllReviews)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/reviews", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateReview(t *testing.T) {
	mock := models.MockDB()
	
	review := models.Review{Content: "Good", Rating: 5}
	body, _ := json.Marshal(review)

	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "reviews"`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := setupTestRouter()
	router.POST("/admin/reviews", CreateReview)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/reviews", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

// --- Tuition Plans ---

func TestGetAllTuitionPlans(t *testing.T) {
	mock := models.MockDB()
	
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1", "Plan 1")
	
	mock.ExpectQuery(`^SELECT \* FROM "tuition_plans"`).
		WillReturnRows(rows)

	router := setupTestRouter()
	router.GET("/admin/tuition-plans", GetAllTuitionPlans)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/admin/tuition-plans", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTuitionPlan(t *testing.T) {
	mock := models.MockDB()
	
	plan := models.TuitionPlan{Name: "Plan 1", Price: 1000}
	body, _ := json.Marshal(plan)

	mock.ExpectBegin()
	mock.ExpectExec(`^INSERT INTO "tuition_plans"`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	router := setupTestRouter()
	router.POST("/admin/tuition-plans", CreateTuitionPlan)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/tuition-plans", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
