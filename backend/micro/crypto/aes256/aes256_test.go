package aes256_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/7chip/little-bird/backend/micro/crypto/aes256"
)

var key1 = "123456789012345678901234567890ab"

var textToEncrypt = "A random string of some length, but which is not extremely long!"

func TestEncryptDecrypt(t *testing.T) {
	c, err := aes256.NewCrypto(key1)
	if err != nil {
		t.Fatal(err)
	}
	originalInput := "This is a test"
	encrypted, err := c.Encrypt(originalInput)
	decrypted, err := c.Decrypt(encrypted)

	assert.EqualValues(t, originalInput, decrypted)
}

func BenchmarkEncrypt(b *testing.B) {
	c, err := aes256.NewCrypto(key1)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Encrypt(textToEncrypt)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	c, err := aes256.NewCrypto(key1)
	if err != nil {
		b.Fatal(err)
	}
	encryptedText, err := c.Encrypt(textToEncrypt)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Decrypt(encryptedText)
	}
}
