package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
)

func NewGcmCipher(hexKeyStr string) (cipher.AEAD, error) {
	hexKey, err := hex.DecodeString(hexKeyStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(hexKey)
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(block)
}

func GcmEncryptWithRandomNonce(cipher cipher.AEAD, plain []byte) ([]byte, error) {
	if cipher == nil {
		return nil, errors.New("cipher not found")
	}
	nonceSize := cipher.NonceSize()
	nonce := make([]byte, nonceSize)

	_, err := rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to create random nonce %w", err)
	}

	return cipher.Seal(nonce, nonce, plain, nil), nil
}

func GcmEncryptWithZeroNonce(cipher cipher.AEAD, plain []byte) ([]byte, error) {
	if cipher == nil {
		return nil, errors.New("cipher not found")
	}
	nonceSize := cipher.NonceSize()
	nonce := make([]byte, nonceSize)

	return cipher.Seal(nonce, nonce, plain, nil), nil
}

func GcmDecrypt(cipher cipher.AEAD, enc []byte) ([]byte, error) {
	nonceSize := cipher.NonceSize()

	if len(enc) <= nonceSize {
		return nil, fmt.Errorf("invalid enc, current length %v", len(enc))
	}
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	return cipher.Open(nil, nonce, ciphertext, nil)
}

func GcmCompare(cipher cipher.AEAD, plain, enc []byte) (bool, error) {
	nonceSize := cipher.NonceSize()

	if len(enc) <= nonceSize {
		return false, fmt.Errorf("invalid enc, current length %v", len(enc))
	}

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plainOfEnc, err := cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, err
	}

	return bytes.Equal(plain, plainOfEnc), nil
}
