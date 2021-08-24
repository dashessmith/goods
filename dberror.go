package goods

import (
	"strings"

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
	// ERROR: duplicate key value violates unique constraint "idx_t_users_token" (SQLSTATE 23505)
	if len(columns) <= 0 && strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
		return true
	}
	return false
}
