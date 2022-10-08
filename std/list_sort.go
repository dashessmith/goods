package std

import (
	"sort"
)

func (z *List[T]) Sort(less func(a, b *T) bool) {
	z.init()
	listSortImpl(z.Begin(), z.End(), func(a, b ListIterator[T]) bool {
		return less(a.Value(), b.Value())
	})
}

type listSortImplStruct[T any] struct {
	dict map[int]ListIterator[T]
	less func(a, b ListIterator[T]) bool
}

func (z *listSortImplStruct[T]) Len() int {
	return len(z.dict)
}

func (z *listSortImplStruct[T]) Less(i, j int) bool {
	return z.less(z.dict[i], z.dict[j])
}

func (z *listSortImplStruct[T]) Swap(i, j int) {
	Swap(z.dict[i].Value(), z.dict[j].Value())
}

func listSortImpl[T any](first, last ListIterator[T], less func(a, b ListIterator[T]) bool) {
	s := listSortImplStruct[T]{
		dict: map[int]ListIterator[T]{},
		less: less,
	}
	for i, x := 0, first; !x.Equal(last); x, i = x.Next(), i+1 {
		s.dict[i] = x
	}
	sort.Sort(&s)
}
