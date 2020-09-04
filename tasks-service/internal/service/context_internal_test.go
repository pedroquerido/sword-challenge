package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseContext(t *testing.T) {

	userID := "user1"
	isManager := true

	t.Run("should return service context", func(t *testing.T) {

		testContext := Context{
			UserID:    userID,
			IsManager: &isManager,
		}

		serviceContext, err := parseContext(context.WithValue(context.Background(), ContextKey, testContext))
		assert.Nil(t, err)
		assert.Equal(t, &testContext, serviceContext)
	})
	t.Run("should return missing context - bad context", func(t *testing.T) {

		testContext := "bad context"

		serviceContext, err := parseContext(context.WithValue(context.Background(), ContextKey, testContext))
		assert.Nil(t, serviceContext)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, ErrorMissingContext))
	})
	t.Run("should return missing context - missing user", func(t *testing.T) {

		testContext := Context{
			IsManager: &isManager,
		}

		serviceContext, err := parseContext(context.WithValue(context.Background(), ContextKey, testContext))
		assert.Nil(t, serviceContext)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, ErrorMissingContext))
	})
	t.Run("should return missing context - missing is_manager", func(t *testing.T) {

		testContext := Context{
			UserID: userID,
		}

		serviceContext, err := parseContext(context.WithValue(context.Background(), ContextKey, testContext))
		assert.Nil(t, serviceContext)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, ErrorMissingContext))
	})
}
