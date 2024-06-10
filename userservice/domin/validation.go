package user

import "regexp"

var usernameRegex = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]$`)
var passwordRegex = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]$`)

func ValidateUserName(username string) bool {
	return usernameRegex.MatchString(username)
}

func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)
}
