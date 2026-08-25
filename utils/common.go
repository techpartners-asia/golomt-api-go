package utils

import (
	"fmt"
	"strconv"
)

func BoolToString(v bool) string {
	if v {
		return "Y"
	}
	return "N"
}

// int64Pointer Get int64 pointer
func int64Pointer(i int64) *int64 {
	return &i
}

// StrToUint String to Uint parser
func StrToUint(value string) (uint, error) {
	u64, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	result := uint(u64)
	return result, nil
}

func AppendAsString(args ...interface{}) string {
	appendedStr := ""
	for _, arg := range args {
		appendedStr = appendedStr + fmt.Sprintf("%v", arg)
	}

	return appendedStr
}

func GetValidString(source interface{}) string {
	switch v := source.(type) {
	case string:
		return v
	case float64:
		// encoding/json decodes every JSON number into a float64.
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		// nil, objects and arrays have no sensible string form.
		return ""
	}
}

func GetValidFloat(source interface{}) float64 {
	switch v := source.(type) {
	case float64:
		return v
	case string:
		// Some responses quote their numbers.
		num, _ := strconv.ParseFloat(v, 64)
		return num
	default:
		return 0
	}
}

func GetValidInt(source interface{}) int {
	switch v := source.(type) {
	case float64:
		return int(v)
	case string:
		num, _ := strconv.Atoi(v)
		return num
	default:
		return 0
	}
}
