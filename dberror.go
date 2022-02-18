package goods

import (
	"regexp"
	"strings"

	"github.com/go-sql-driver/mysql"
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
	case *mysql.MySQLError:
		if et.Number == 1062 {
			if len(columns) <= 0 {
				return true
			}
		}
	}
	// ERROR: duplicate key value violates unique constraint "idx_t_users_token" (SQLSTATE 23505)
	if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") &&
		All(len(columns), func(index int) bool {
			matched, _ := regexp.MatchString("\".*"+columns[index]+".*\"", err.Error())
			return matched
		}) {
		return true
	}
	// 错误: 重复键违反唯一约束"t_goods_name_key"
	if strings.Contains(err.Error(), "错误: 重复键违反唯一约束") &&
		All(len(columns), func(index int) bool {
			matched, _ := regexp.MatchString("\".*"+columns[index]+".*\"", err.Error())
			return matched
		}) {
		return true
	}
	return false
}
