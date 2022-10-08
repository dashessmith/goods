package std

import (
	"sort"
	"testing"

	"github.com/dashessmith/goods"
)

func TestMap(t *testing.T) {
	m := Map[string, string]{}
	m.Upsert(Pair[string, string]{`2`, `1`})
	m.Upsert(Pair[string, string]{`1`, `1`})
	m.Upsert(Pair[string, string]{`3`, `1`})
	m.Set(`4`, `2`)
	t.Logf("keys = %v", m.Contains(`4`))

	goods.AssertTrue(t, sort.IsSorted(VectorSortable[string](m.Keys())))
}
