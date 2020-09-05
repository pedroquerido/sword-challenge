package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetContentTypeJSON(t *testing.T) {

	t.Run("should set header content type as json and continue", func(t *testing.T) {

		testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			contentType := w.Header().Get("Content-Type")
			assert.Equal(t, "application/json", contentType)
		})

		handlerToTest := setContentTypeJSON(testHandler)
		req := httptest.NewRequest(http.MethodGet, "/tests", nil)

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	})
}

func TestRecoverFromPanic(t *testing.T) {

	t.Run("should recover from panic and return internal error response", func(t *testing.T) {

		expectedResponse := response.NewErrorResponse(http.StatusInternalServerError, messageInternal)

		panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("unexpected panic")
		})

		handlerToTest := recoverFromPanic(panicHandler)

		req := httptest.NewRequest(http.MethodGet, "/tests", nil)
		recorder := httptest.NewRecorder()
		handlerToTest.ServeHTTP(recorder, req)

		outputResponse := response.ErrorResponse{}
		err := json.NewDecoder(recorder.Body).Decode(&outputResponse)
		require.Nil(t, err)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Equal(t, expectedResponse.Code, outputResponse.Code)
		assert.Equal(t, expectedResponse.Message, outputResponse.Message)
		assert.Equal(t, expectedResponse.Errors, outputResponse.Errors)
	})
}

func TestRequireHeaders(t *testing.T) {

	t.Run("should do nothing", func(t *testing.T) {

		handlerToTest := requireHeaders(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

		req := httptest.NewRequest(http.MethodGet, "/tests", nil)
		req.Header.Set(headerUserID, "user_id")
		req.Header.Set(headerUserRole, "user_role")

		recorder := httptest.NewRecorder()
		handlerToTest.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "", recorder.Body.String())
	})
	t.Run("should return bad request response", func(t *testing.T) {

		missingHeader := "missing header %s"
		expectedResponse := response.NewErrorResponse(http.StatusBadRequest, messageBadRequest,
			[]string{fmt.Sprintf(missingHeader, headerUserID), fmt.Sprintf(missingHeader, headerUserRole)}...)

		handlerToTest := requireHeaders(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

		req := httptest.NewRequest(http.MethodGet, "/tests", nil)

		recorder := httptest.NewRecorder()
		handlerToTest.ServeHTTP(recorder, req)

		outputResponse := response.ErrorResponse{}
		err := json.NewDecoder(recorder.Body).Decode(&outputResponse)
		require.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, expectedResponse.Code, outputResponse.Code)
		assert.Equal(t, expectedResponse.Message, outputResponse.Message)
		assert.Equal(t, expectedResponse.Errors, outputResponse.Errors)
	})
}
