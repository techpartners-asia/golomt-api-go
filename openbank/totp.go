package openbank

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultTimeStepSeconds = 30
	NumDigitsOutput        = 6
)

// GenerateBase32Secret generates a random base32 secret (default length 16)
func GenerateBase32Secret() string {
	return GenerateBase32SecretWithLength(16)
}

// GenerateBase32SecretWithLength generates a random base32 secret of given length
func GenerateBase32SecretWithLength(length int) string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	bytes := make([]byte, length)
	_, _ = rand.Read(bytes)
	for i := range bytes {
		bytes[i] = alphabet[int(bytes[i])%len(alphabet)]
	}
	return string(bytes)
}

// GenerateCurrentNumber generates current TOTP integer
func GenerateCurrentNumber(base32Secret string) (int, error) {
	return GenerateNumber(base32Secret, time.Now().UnixMilli(), DefaultTimeStepSeconds)
}

// GenerateCurrentNumberString generates current TOTP string (zero-padded)
func GenerateCurrentNumberString(base32Secret string) (string, error) {
	num, err := GenerateCurrentNumber(base32Secret)
	if err != nil {
		return "", err
	}
	return zeroPrepend(num, NumDigitsOutput), nil
}

// ValidateCurrentNumber checks if a given TOTP is valid within the time window (milliseconds)
func ValidateCurrentNumber(base32Secret string, authNumber int, windowMillis int) (bool, error) {
	return ValidateNumberAtTime(base32Secret, authNumber, windowMillis, time.Now().UnixMilli(), DefaultTimeStepSeconds)
}

// ValidateNumberAtTime validates the TOTP number around a specific timestamp
func ValidateNumberAtTime(base32Secret string, authNumber int, windowMillis int, timeMillis int64, timeStepSeconds int) (bool, error) {
	from := timeMillis
	to := timeMillis
	if windowMillis > 0 {
		from = timeMillis - int64(windowMillis)
		to = timeMillis + int64(windowMillis)
	}

	stepMillis := int64(timeStepSeconds * 1000)
	for millis := from; millis <= to; millis += stepMillis {
		gen, err := GenerateNumber(base32Secret, millis, timeStepSeconds)
		if err != nil {
			return false, err
		}
		if gen == authNumber {
			return true, nil
		}
	}
	return false, nil
}

// GenerateNumber generates TOTP number for a given time
func GenerateNumber(base32Secret string, timeMillis int64, timeStepSeconds int) (int, error) {
	key, err := decodeBase32(base32Secret)
	if err != nil {
		return 0, err
	}

	value := timeMillis / 1000 / int64(timeStepSeconds)
	data := make([]byte, 8)
	for i := 7; value > 0; i-- {
		data[i] = byte(value & 0xFF)
		value >>= 8
	}

	h := hmac.New(sha1.New, key)
	h.Write(data)
	hash := h.Sum(nil)
	offset := hash[len(hash)-1] & 0x0F

	truncated := int64((int(hash[offset]&0x7F) << 24) |
		(int(hash[offset+1]&0xFF) << 16) |
		(int(hash[offset+2]&0xFF) << 8) |
		(int(hash[offset+3] & 0xFF)))

	truncated = truncated % 1000000
	return int(truncated), nil
}

// --- Helper functions ---

func zeroPrepend(num, digits int) string {
	s := strconv.Itoa(num)
	if len(s) >= digits {
		return s
	}
	return strings.Repeat("0", digits-len(s)) + s
}

// decodeBase32 нь SPEC Хавсралт 3-ын Java decodeBase32-ийн яг хуулбар:
// жижиг/том үсэг зөвшөөрнө, '=' дээр зогсоно, 8-д хуваагдахгүй урттай
// secret-ийн сүүлийн дутуу byte-ийг мөн нэмнэ.
func decodeBase32(secret string) ([]byte, error) {
	secret = strings.ReplaceAll(secret, " ", "")
	result := make([]byte, 0, (len(secret)*5+7)/8)
	which := 0
	working := 0
	for i := 0; i < len(secret); i++ {
		ch := secret[i]
		var val int
		switch {
		case ch >= 'a' && ch <= 'z':
			val = int(ch - 'a')
		case ch >= 'A' && ch <= 'Z':
			val = int(ch - 'A')
		case ch >= '2' && ch <= '7':
			val = 26 + int(ch-'2')
		case ch == '=':
			which = 0
			i = len(secret)
			continue
		default:
			return nil, errors.New("invalid base32 secret")
		}
		switch which {
		case 0:
			working = (val & 31) << 3
			which = 1
		case 1:
			working |= (val & 28) >> 2
			result = append(result, byte(working))
			working = (val & 3) << 6
			which = 2
		case 2:
			working |= (val & 31) << 1
			which = 3
		case 3:
			working |= (val & 16) >> 4
			result = append(result, byte(working))
			working = (val & 15) << 4
			which = 4
		case 4:
			working |= (val & 30) >> 1
			result = append(result, byte(working))
			working = (val & 1) << 7
			which = 5
		case 5:
			working |= (val & 31) << 2
			which = 6
		case 6:
			working |= (val & 24) >> 3
			result = append(result, byte(working))
			working = (val & 7) << 5
			which = 7
		case 7:
			working |= val & 31
			result = append(result, byte(working))
			which = 0
		}
	}
	if which != 0 {
		result = append(result, byte(working))
	}
	return result, nil
}
