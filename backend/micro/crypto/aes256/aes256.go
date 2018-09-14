package aes256

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

var encoding = base64.URLEncoding

// Creates an instance of TokenCrypto, used for AES encryption/decryption.
//
// Takes a 32 characters long string, which results in using AES-256
func NewCrypto(encryptionKey string) (*TokenCrypto, error) {
	if len(encryptionKey) != 32 {
		return nil, errors.New("aes256: token encryption key must be 32 characters")
	}
	block, err := aes.NewCipher([]byte(encryptionKey))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return &TokenCrypto{aesgcm}, nil
}

type TokenCrypto struct {
	aesgcm cipher.AEAD
}

// Generates random nonce, and uses this and the key to encrypt the plaintext message.
// The nonce and cipher byte slices are are concatenated, encoded as base64, and returned.
func (t *TokenCrypto) Encrypt(plaintext string) (string, error) {
	nonce := make([]byte, t.aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	cipherBytes := t.aesgcm.Seal(nil, nonce, []byte(plaintext), nil)

	outBytes := append(nonce, cipherBytes...)
	return encoding.EncodeToString(outBytes), nil
}

// Decodes the bytes from base64. Then reads the nonce from the first bytes, and uses this and the key
// to decrypt the plaintext message
func (t *TokenCrypto) Decrypt(cipherBytesBase64 string) (string, error) {
	cipherBytes, err := encoding.DecodeString(cipherBytesBase64)
	if err != nil {
		return "", err
	}
	nonce := cipherBytes[:t.aesgcm.NonceSize()]
	input := cipherBytes[t.aesgcm.NonceSize():]
	plainBytes, err := t.aesgcm.Open(nil, nonce, input, nil)
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}
