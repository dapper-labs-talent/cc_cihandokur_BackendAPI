package helper

import "regexp"

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

func EmailCheck(email string) bool {
	return regexp.MustCompile(emailRegex).MatchString(email)
}
