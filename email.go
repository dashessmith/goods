package goods

import (
	"net/mail"
	"strings"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func EmailTokens(email string) (prefix string, suffix string) {
	tokens := strings.Split(email, `@`)
	if len(tokens) > 1 {
		prefix = tokens[0]
	}
	if len(tokens) > 2 {
		suffix = tokens[1]
	}
	return
}
