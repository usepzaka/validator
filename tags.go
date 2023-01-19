package validator

import (
	"reflect"
	"strconv"
	"strings"
)

func validateStructTag(valKey string, jsonField string, valueField reflect.Value, typeField reflect.StructField) (bool, string) {
	strMessage := ""
	switch valKey {
	case "required":
		if valueField.IsZero() {
			strMessage := "--> " + jsonField + " <-- is Required"

			return false, strMessage
		}
	case "alphanumeric":
		if typeField.Type.String() == "string" && !IsAlphanumeric(valueField.String()) {
			strMessage := "--> " + jsonField + " <-- must be Alphanumeric."
			return false, strMessage
		}
	case "phone":
		if typeField.Type.String() == "string" && !IsNumberPlus(valueField.String()) {
			strMessage := "--> " + jsonField + " <-- must be Phone Number Format."
			return false, strMessage
		}
		parsedPhoneNumber := ParseNo(valueField.String())
		if len(parsedPhoneNumber) < 9 || len(parsedPhoneNumber) > 12 {
			strMessage := "--> " + jsonField + " <-- minimum 10 digit and maximum 13 digit for prefix 0. min 11 digit and 14 digit with prefix 62"
			return false, strMessage
		}
	case "letter":
		if typeField.Type.String() == "string" && !IsAlpha(valueField.String()) {
			strMessage := "--> " + jsonField + " <-- must be Letter a - z and A - Z."
			return false, strMessage
		}
	case "number":
		if typeField.Type.String() == "string" && !IsNumberOnly(valueField.String()) {
			strMessage := "--> " + jsonField + " <-- must be Number 0 - 9."
			return false, strMessage
		}
	case "letternumber":
		if typeField.Type.String() == "string" && !IsAlphanumeric(valueField.String()) {
			strMessage := "--> " + jsonField + " <-- must be Phone Number Format."
			return false, strMessage
		}
	case "lowercase":
		if typeField.Type.String() == "string" && !IsLowerCase(valueField.String()) {
			strMessage := "--> " + jsonField + " <-- must be lowercase Format."
			return false, strMessage
		}
	case "uppercase":
		if typeField.Type.String() == "string" && !IsUpperCase(valueField.String()) {
			strMessage := "--> " + jsonField + " <-- must be UPPERCASE Format."
			return false, strMessage
		}

	}

	switch {
	case strings.HasPrefix(valKey, "min="):
		min := strings.Split(valKey, "=")
		intg, _ := strconv.Atoi(min[1])
		strMessage := "--> " + jsonField + " <-- minimum " + min[1] + " character."
		if len(valueField.String()) < intg {
			return false, strMessage
		}
	case strings.HasPrefix(valKey, "max="):
		max := strings.Split(valKey, "=")
		intg, _ := strconv.Atoi(max[1])
		strMessage := "--> " + jsonField + " <-- maximum " + max[1] + " character."
		if len(valueField.String()) > intg {
			return false, strMessage
		}
	case strings.HasPrefix(valKey, "must="):
		val := strings.Split(valKey, "=")
		intg, _ := strconv.Atoi(val[1])
		strMessage := "--> " + jsonField + " <-- Must be " + val[1] + " character."
		if len(valueField.String()) != intg {
			return false, strMessage
		}
	}

	return true, strMessage
}
