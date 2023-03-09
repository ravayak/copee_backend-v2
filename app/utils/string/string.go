package stringutils

import (
	"errors"
	"strconv"
	"strings"
)

func IsEmpty(message string) bool {
	if strings.TrimSpace(message) == "" {
		return true
	}
	return false
}

func StrToInt(str string) (int64, error) {
	if strings.TrimSpace(str) == "" {
		err := errors.New("invalid number")
		return 0, err
	}

	number, getErr := strconv.ParseInt(str, 10, 64)
	if getErr != nil {
		err := errors.New("invalid number")
		return 0, err
	}
	return number, nil
}
