package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
)

// Encrypt encrypts SHA256 hex of `data` using a base64-encoded public key
func EncryptRSA(data string, base64PublicKey string) (string, error) {
	// Decode base64-encoded public key
	pubBytes, err := base64.StdEncoding.DecodeString(base64PublicKey)
	if err != nil {
		return "", err
	}

	// Parse X509-encoded public key
	pubInterface, err := x509.ParsePKIXPublicKey(pubBytes)
	if err != nil {
		return "", err
	}
	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return "", errors.New("not RSA public key")
	}

	// Hash input data with SHA256 and get hex string
	hash := sha256.Sum256([]byte(data))
	hexData := []byte(fmt.Sprintf("%x", hash))

	// Encrypt
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, hexData)
	if err != nil {
		return "", err
	}

	// Return base64
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts base64-encoded cipher text using a base64-encoded PKCS8 private key
func DecryptRSA(cipherText string, base64PrivateKey string) (string, error) {
	// Decode base64-encoded private key
	privBytes, err := base64.StdEncoding.DecodeString(base64PrivateKey)
	if err != nil {
		return "", err
	}

	// Parse PKCS8 private key
	privInterface, err := x509.ParsePKCS8PrivateKey(privBytes)
	if err != nil {
		return "", err
	}
	privKey, ok := privInterface.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not RSA private key")
	}

	// Decode base64 cipher text
	encryptedBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	// Decrypt
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, encryptedBytes)
	if err != nil {
		return "", err
	}

	return string(decryptedBytes), nil
}
