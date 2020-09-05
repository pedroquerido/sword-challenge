package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	pkgError "github.com/pedroquerido/sword-challenge/tasks-service/pkg/error"
)

var _ Encryptor = (*encryptor)(nil)

// Encryptor ...
type Encryptor interface {
	Encrypt(value string) (string, error)
	Decrypt(value string) (string, error)
}

type encryptor struct {
	aesGCM cipher.AEAD
}

// NewEncryptor ...
func NewEncryptor(key string) (Encryptor, error) {

	// convert to bytes
	keyBytes := []byte(key)

	// validate len (16, 24 or 32)
	switch len(keyBytes) {
	case 16:
		fallthrough
	case 24:
		fallthrough
	case 32:
		// do nothing
	default:
		return nil, ErrorInvalidKey
	}

	// create new cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	// Create new gcm
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &encryptor{
		aesGCM: aesGCM,
	}, nil
}

func (e *encryptor) Encrypt(value string) (string, error) {

	valueBytes := []byte(value)

	// create nonce from GCM
	nonce := make([]byte, e.aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// encrypt the data
	encrypted := e.aesGCM.Seal(nonce, nonce, valueBytes, nil)
	return fmt.Sprintf("%x", encrypted), nil
}

func (e *encryptor) Decrypt(value string) (string, error) {

	encrypted, err := hex.DecodeString(value)
	if err != nil {
		return "", pkgError.NewError(ErrorEncryptedWithOtherEncryptor).WithDetails(err.Error())
	}

	// get nonce size and validate value length
	nonceSize := e.aesGCM.NonceSize()
	if nonceSize > len(encrypted) {
		return "", pkgError.NewError(ErrorEncryptedWithOtherEncryptor).WithDetails("invalid value size")
	}

	// extract nonce from encrypted
	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]

	// decrypt
	decrypted, err := e.aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", pkgError.NewError(ErrorEncryptedWithOtherEncryptor).WithDetails(err.Error())
	}

	return fmt.Sprintf("%s", decrypted), nil
}
