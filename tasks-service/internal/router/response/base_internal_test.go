package response

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBaseResponse(t *testing.T) {

	t.Run("should return base response with generated uuid", func(t *testing.T) {

		message := "message"
		code := 200

		baseResponse := newBaseResponse(message, code)
		assert.NotNil(t, baseResponse)
		assert.Equal(t, message, baseResponse.Message)
		assert.Equal(t, code, baseResponse.Code)
		assert.NotNil(t, baseResponse.ID)

		generatedUUID, err := uuid.Parse(baseResponse.ID)
		require.Nil(t, err)
		assert.NotNil(t, generatedUUID)
	})
}
