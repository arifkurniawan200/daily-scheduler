package utils

import "regexp"

func IsValidDate(date string) bool {
	// Define a regular expression pattern for "YYYY-MM-DD" format
	datePattern := `^\d{4}-\d{2}-\d{2}$`
	r := regexp.MustCompile(datePattern)

	return r.MatchString(date)
}
