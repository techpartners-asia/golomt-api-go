package mini_app

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// Step 1: Convert JSON string to SHA256 hex (64-character)
func sha256Hex(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// Step 2: Encrypt the 64-character hex string using RSA/PKCS1Padding
func encryptRSAWithPublicKey(hexText string, base64PublicKey string) (string, error) {
	// Decode Base64 encoded X.509 public key
	pubBytes, err := base64.StdEncoding.DecodeString(base64PublicKey)
	if err != nil {
		return "", fmt.Errorf("invalid base64 public key: %w", err)
	}

	// Parse public key
	pubInterface, err := x509.ParsePKIXPublicKey(pubBytes)
	if err != nil {
		return "", fmt.Errorf("invalid X.509 public key: %w", err)
	}

	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("not an RSA public key")
	}

	// Encrypt using PKCS1 v1.5 (ECB is default)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(hexText))
	if err != nil {
		return "", fmt.Errorf("RSA encryption failed: %w", err)
	}

	// Return Base64 string
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// Step 3: Complete function to generate X-Golomt-Signature
func generateGolomtSignature(jsonData string, base64PublicKey string) (string, error) {
	hashHex := sha256Hex(jsonData)
	return encryptRSAWithPublicKey(hashHex, base64PublicKey)
}
