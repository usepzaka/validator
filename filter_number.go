package validator

import (
	"regexp"
	"strings"
)

func IsNumberOnly(value string) bool {
	value = regexp.MustCompile(`\d`).ReplaceAllString(value, "")
	return value == ""
}

func IsNumberPlus(value string) bool {
	value = regexp.MustCompile(`^\+|[0-9]`).ReplaceAllString(value, "")
	return value == ""
}

func IsNumberAndDot(value string) bool {
	value = regexp.MustCompile(`[0-9\.\s]`).ReplaceAllString(value, "")
	return value == ""
}

// Parsing string into phone number string without 0 or 62
func ParseNo(number string) string {
	//Replace spasi
	number = strings.Replace(number, " ", "", -1)
	// remove any non-digit character, included the +
	number = regexp.MustCompile(`\D`).ReplaceAllString(number, "")
	// remove 0 in front
	number = regexp.MustCompile(`^0+`).ReplaceAllString(number, "")
	// remove 62 in front
	number = regexp.MustCompile(`^62+`).ReplaceAllString(number, "")

	return number
}

// Check msisdn provided is Indosat Number or not
// Return is boolean
func IsNumberIsat(msisdn string) bool {
	switch {
	case strings.HasPrefix(msisdn, "815"),
		strings.HasPrefix(msisdn, "816"),
		strings.HasPrefix(msisdn, "855"),
		strings.HasPrefix(msisdn, "856"),
		strings.HasPrefix(msisdn, "857"),
		strings.HasPrefix(msisdn, "858"):
		return true
	default:
		return false
	}
}
