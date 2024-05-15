package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
)

func ParsePrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	privatePem, _ := pem.Decode(key)
	return x509.ParsePKCS1PrivateKey(privatePem.Bytes)
}

func ParsePrivateKeyFromBase64(key string) (*rsa.PrivateKey, error) {
	publicKeyPem, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKeyFromPEM(publicKeyPem)
}

func ParsePrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	rawPrivateKey, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKeyFromPEM(rawPrivateKey)
}

func ParsePublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		return publicKey, nil
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func ParsePublicKeyFromBase64(key string) (*rsa.PublicKey, error) {
	publicKeyPem, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return ParsePublicKeyFromPEM(publicKeyPem)
}

func ParsePublicKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	publicKeyPem, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return ParsePublicKeyFromPEM(publicKeyPem)
}

func VerifyBase64Signature(publicKey *rsa.PublicKey, signature string, plain []byte) error {
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	hashedPlain := sha1.Sum(plain)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, hashedPlain[:], signatureBytes)
}

func GenerateBase64Signature(privateKey *rsa.PrivateKey, plain []byte) (string, error) {
	hashed := sha1.Sum(plain)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed[:])
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}
