package utils

import (
	"regexp"
)

const (
	alphaNumericUnderscore = `^[a-zA-Z0-9_]*$`
)

func AlphaNumericUnderscore(str string) (bool, error) {
	regex, err := regexp.Compile(alphaNumericUnderscore)
	if err != nil {
		return false, err
	}
	return regex.MatchString(str), nil
}
