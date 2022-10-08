package goods

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](src []T) (ret []T) {
	ret = Slice(src)
	sort.Slice(ret, func(i, j int) bool { return ret[i] < ret[j] })
	return
}
