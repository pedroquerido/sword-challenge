package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedroquerido/sword-challenge/mock-auth-server/internal/router"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func TestRouter_AuthHandler(t *testing.T) {

	headerAPIKey := "x-api-key"
	headerUserID := "x-user-id"
	headerUserRole := "x-user-role"

	apiKeyUser := map[string]string{
		"key1": "user1",
		"key2": "user2",
		"key3": "user3",
	}

	userRole := map[string]string{
		"user1": "manager",
		"user2": "technician",
	}

	testRouter := router.New(apiKeyUser, userRole)

	t.Run("should return status 401 - no api key", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/auth", nil)
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})
	t.Run("should return status 401 - bad api key", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/auth", nil)
		req.Header.Set(headerAPIKey, "key0")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})
	t.Run("should return status 401 - missing role for user", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/auth", nil)
		req.Header.Set(headerAPIKey, "key3")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})
	t.Run("should return status 200 - with headers", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/auth", nil)
		req.Header.Set(headerAPIKey, "key1")
		require.Nil(t, err)

		recorder := httptest.NewRecorder()
		testRouter.GetHTTPHandler().ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "user1", recorder.Header().Get(headerUserID))
		assert.Equal(t, "manager", recorder.Header().Get(headerUserRole))
	})
}
