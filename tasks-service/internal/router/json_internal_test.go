package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedroquerido/sword-challenge/tasks-service/internal/router/response"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteJSON(t *testing.T) {

	t.Run("should write header and json response", func(t *testing.T) {

		expectedResponse := response.NewErrorResponse(http.StatusBadRequest, messageBadRequest, []string{"error 1"}...)

		recorder := httptest.NewRecorder()

		writeJSON(recorder, http.StatusBadRequest, expectedResponse)
		outputResponse := response.ErrorResponse{}
		err := json.NewDecoder(recorder.Body).Decode(&outputResponse)
		require.Nil(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, expectedResponse.Code, outputResponse.Code)
		assert.Equal(t, expectedResponse.Message, outputResponse.Message)
		assert.Equal(t, expectedResponse.Errors, outputResponse.Errors)
	})
	t.Run("should only write header without data", func(t *testing.T) {

		recorder := httptest.NewRecorder()

		writeJSON(recorder, http.StatusBadRequest, nil)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, "", recorder.Body.String())
	})
	t.Run("should only write header with data that cannot be represented in json", func(t *testing.T) {

		recorder := httptest.NewRecorder()

		writeJSON(recorder, http.StatusBadRequest, make(chan string))

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, "", recorder.Body.String())
	})
}
