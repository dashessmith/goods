package util

import "errors"

func SameError(e1, e2 error) bool {
	if errors.Is(e1, e2) {
		return true
	}
	if e1 != nil && e2 != nil && e1.Error() == e2.Error() {
		return true
	}
	return false
}
