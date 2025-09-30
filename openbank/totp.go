package openbank

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"strings"
	"time"
)

const (
	DefaultTimeStepSeconds = 30
	NumDigitsOutput        = 6
)

// GenerateCurrentNumber generates the current TOTP code.
func GenerateCurrentNumber(base32Secret string) (int, error) {
	return GenerateNumber(base32Secret, time.Now().Unix()*1000, DefaultTimeStepSeconds)
}

// GenerateCurrentNumberString generates the current TOTP code as zero-padded string.
func GenerateCurrentNumberString(base32Secret string) (string, error) {
	n, err := GenerateCurrentNumber(base32Secret)
	if err != nil {
		return "", err
	}
	return ZeroPrepend(n, NumDigitsOutput), nil
}

// ValidateCurrentNumber checks if a given code is valid within +/- windowMillis.
func ValidateCurrentNumber(base32Secret string, authNumber int, windowMillis int) (bool, error) {
	return ValidateAtTime(base32Secret, authNumber, windowMillis, time.Now().Unix()*1000, DefaultTimeStepSeconds)
}

// ValidateAtTime validates a TOTP code for a given timestamp and step.
func ValidateAtTime(base32Secret string, authNumber int, windowMillis int, timeMillis int64, timeStepSeconds int) (bool, error) {
	fromTimeMillis := timeMillis
	toTimeMillis := timeMillis
	if windowMillis > 0 {
		fromTimeMillis = timeMillis - int64(windowMillis)
		toTimeMillis = timeMillis + int64(windowMillis)
	}
	timeStepMillis := int64(timeStepSeconds * 1000)

	for millis := fromTimeMillis; millis <= toTimeMillis; millis += timeStepMillis {
		generated, err := GenerateNumber(base32Secret, millis, timeStepSeconds)
		if err != nil {
			return false, err
		}
		if generated == authNumber {
			return true, nil
		}
	}
	return false, nil
}

// GenerateNumber generates the TOTP code for given timestamp and step.
func GenerateNumber(base32Secret string, timeMillis int64, timeStepSeconds int) (int, error) {
	key, err := DecodeBase32(base32Secret)
	if err != nil {
		return 0, err
	}

	value := timeMillis / 1000 / int64(timeStepSeconds)
	data := make([]byte, 8)
	for i := 7; value > 0; i-- {
		data[i] = byte(value & 0xFF)
		value >>= 8
	}

	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	hash := mac.Sum(nil)

	offset := hash[len(hash)-1] & 0x0F
	truncatedHash := int64(0)
	for i := 0; i < 4; i++ {
		truncatedHash <<= 8
		truncatedHash |= int64(hash[int(offset)+i] & 0xFF)
	}
	truncatedHash &= 0x7FFFFFFF
	truncatedHash %= 1000000

	return int(truncatedHash), nil
}

// ZeroPrepend pads a number with zeros to required digits.
func ZeroPrepend(num, digits int) string {
	return fmt.Sprintf("%0*d", digits, num)
}

// DecodeBase32 decodes a Base32 string (case-insensitive, no padding).
func DecodeBase32(str string) ([]byte, error) {
	str = strings.ToUpper(str)
	return base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(str)
}
