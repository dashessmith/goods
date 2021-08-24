package goods

import (
	"github.com/lib/pq"
)

func IsViolatesUniqueError(err error, columns ...string) bool {
	if err == nil {
		return false
	}
	switch et := err.(type) {
	case *pq.Error:
		if et.Code == `23505` {
			if len(columns) <= 0 {
				return true
			}
		}
	}
	return false
}
