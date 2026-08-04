package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-ocean-basketball/config"
	"backend-ocean-basketball/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := &config.Config{
		JWTSecret: "testsecret",
	}

	validToken, _ := utils.GenerateJWT("uuid-1", "test@test.com", "admin", "testsecret")

	tests := []struct {
		name          string
		header        string
		requiredRoles []string
		expectedCode  int
	}{
		{
			name:          "No Auth Header",
			header:        "",
			requiredRoles: []string{},
			expectedCode:  http.StatusUnauthorized,
		},
		{
			name:          "Invalid Format",
			header:        "BearerTokenWithoutSpace",
			requiredRoles: []string{},
			expectedCode:  http.StatusUnauthorized,
		},
		{
			name:          "Invalid Token",
			header:        "Bearer invalidtoken",
			requiredRoles: []string{},
			expectedCode:  http.StatusUnauthorized,
		},
		{
			name:          "Valid Token, No Role Required",
			header:        "Bearer " + validToken,
			requiredRoles: []string{},
			expectedCode:  http.StatusOK,
		},
		{
			name:          "Valid Token, Valid Role",
			header:        "Bearer " + validToken,
			requiredRoles: []string{"admin"},
			expectedCode:  http.StatusOK,
		},
		{
			name:          "Valid Token, Invalid Role",
			header:        "Bearer " + validToken,
			requiredRoles: []string{"coach"},
			expectedCode:  http.StatusForbidden,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := gin.Default()
			r.Use(AuthMiddleware(cfg, tc.requiredRoles...))
			r.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tc.header != "" {
				req.Header.Set("Authorization", tc.header)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedCode, w.Code)
		})
	}
}
