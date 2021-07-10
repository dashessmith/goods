package util_test

import (
	"testing"

	"github.com/dashessmith/util"
)

func Test_map(t *testing.T) {
	m := map[string]interface{}{
		`5`: `5`,
		`2`: `2`,
		`3`: `3`,
		`4`: `4`,
	}
	keys := []string{}
	util.MapKeys(m, &keys)
	t.Logf("%v\n", keys)
	util.MapSortedKeys(m, &keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	t.Logf("%v\n", keys)
}
