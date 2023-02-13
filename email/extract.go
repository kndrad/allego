package email

import "regexp"

var (
	re = regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)
)

// Extract takes a string as input and returns the email address found within the string using regular expression pattern matching.
// If no email address is found, an empty string is returned.
func Extract(s string) string {
	email := re.FindString(s)
	return email
}
