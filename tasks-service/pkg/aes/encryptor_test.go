package aes_test

import (
	"errors"
	"testing"

	"github.com/pedroquerido/sword-challenge/tasks-service/pkg/aes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEncryptor(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {

		testKey := "N6bLoi7wcPl6egUs10FdfQgOtqBFHWKc" // 32
		encryptor, err := aes.NewEncryptor(testKey)
		require.Nil(t, err)
		assert.NotNil(t, encryptor)

		testKey = "DUlXlVyp10FdfQgOtqBFHWKc" // 24
		encryptor, err = aes.NewEncryptor(testKey)
		require.Nil(t, err)
		assert.NotNil(t, encryptor)

		testKey = "10FdfQgOtqBFHWKc" // 16
		encryptor, err = aes.NewEncryptor(testKey)
		require.Nil(t, err)
		assert.NotNil(t, encryptor)
	})
	t.Run("should return error invalid key", func(t *testing.T) {

		testKey := "DUlXlVyp" // 8
		encryptor, err := aes.NewEncryptor(testKey)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, aes.ErrorInvalidKey))
		assert.Nil(t, encryptor)
	})
}

func TestEncryptor_Encrypt(t *testing.T) {

	testKey := "10FdfQgOtqBFHWKc"
	encryptor, err := aes.NewEncryptor(testKey)
	require.Nil(t, err)

	t.Run("should return no error", func(t *testing.T) {

		testString := "some random string"

		encrypted, err := encryptor.Encrypt(testString)
		require.Nil(t, err)
		assert.NotEqual(t, testString, encrypted)
	})
}

func TestEncryptor_Decrypt(t *testing.T) {

	testKey := "10FdfQgOtqBFHWKc"
	encryptor, err := aes.NewEncryptor(testKey)
	require.Nil(t, err)

	toBeEncrypted := "some random string"
	encrypted, err := encryptor.Encrypt(toBeEncrypted)
	require.Nil(t, err)

	t.Run("should return no error", func(t *testing.T) {

		decrypted, err := encryptor.Decrypt(encrypted)
		require.Nil(t, err)
		assert.Equal(t, toBeEncrypted, decrypted)
	})
	t.Run("should return error encrypted with other encryptor - invalid hex", func(t *testing.T) {

		testString := "."

		decrypted, err := encryptor.Decrypt(testString)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, aes.ErrorEncryptedWithOtherEncryptor))
		assert.NotEqual(t, testString, decrypted)
	})
	t.Run("should return error encrypted with other encryptor - invalid value size", func(t *testing.T) {

		testString := "19fca073c"

		decrypted, err := encryptor.Decrypt(testString)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, aes.ErrorEncryptedWithOtherEncryptor))
		assert.NotEqual(t, testString, decrypted)
	})
	t.Run("should return error encrypted with other encryptor - invalid message auth", func(t *testing.T) {

		testKey2 := "10FdfQgOtqBFHWKa"
		encryptor2, err := aes.NewEncryptor(testKey2)
		require.Nil(t, err)

		encrypted2, err := encryptor2.Encrypt(toBeEncrypted)
		require.Nil(t, err)

		decrypted, err := encryptor.Decrypt(encrypted2)
		require.NotNil(t, err)
		assert.True(t, errors.Is(err, aes.ErrorEncryptedWithOtherEncryptor))
		assert.NotEqual(t, toBeEncrypted, decrypted)
	})
}
