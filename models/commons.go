package models

import "strings"

func GetStringPointer(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}

func toLower(str *string) {
	if str == nil {
		return
	}

	strLocal := strings.ToLower(*str)
	*str = strLocal
}
