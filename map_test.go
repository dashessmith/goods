package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

type UserLevel int

const (
	UserLevel_0   UserLevel = 0
	UserLevel_100 UserLevel = 100
	UserLevel_200 UserLevel = 200
	UserLevel_300 UserLevel = 300
)

func Test_map(t *testing.T) {
	m := map[UserLevel]interface{}{
		UserLevel_200: `5`,
		UserLevel_0:   `2`,
		UserLevel_100: `3`,
		UserLevel_300: `4`,
	}
	var keys []UserLevel
	goods.MapKeys(m, &keys)
	t.Logf("%v\n", keys)
	goods.MapSortedKeys(m, &keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	t.Logf("%v\n", keys)
}
