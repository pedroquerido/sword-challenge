package aes

import "errors"

var (
	// ErrorInvalidKey represents the error obtained if a key does not meet requirements
	ErrorInvalidKey = errors.New("invalid key")

	// ErrorEncryptedWithOtherEncryptor represents the error obtained if a value was not encrypted by this Encryptor
	ErrorEncryptedWithOtherEncryptor = errors.New("encryted by other encryptor")
)
