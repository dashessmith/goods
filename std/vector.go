package std

import (
	"math"
	"sort"

	"golang.org/x/exp/constraints"
)

type Vector[T any] []T

func (z *Vector[T]) Front() *T {
	return &(*z)[0]
}

func (z *Vector[T]) Back() *T {
	return &(*z)[len(*z)-1]
}

func (z *Vector[T]) Clear() {
	(*z) = Vector[T]{}
}

func (z *Vector[T]) Copy() Vector[T] {
	ret := make(Vector[T], len(*z))
	copy(ret, *z)
	return ret
}

func (z *Vector[T]) Size() int {
	return len(*z)
}

func (z Vector[T]) Len() int {
	return len(z)
}

func (z *Vector[T]) Empty() bool {
	return len(*z) <= 0
}

func (z *Vector[T]) PushBack(x T) {
	*z = append((*z), x)
}

func (z *Vector[T]) PopBack() {
	z.EraseAt(z.Size() - 1)
}

func (z *Vector[T]) PushFront(x T) {
	z.InsertBefore(0, x)
}

func (z *Vector[T]) Begin() VectorIterator[T] {
	return VectorIterator[T]{
		v:   z,
		pos: 0,
	}
}

func (z *Vector[T]) End() VectorIterator[T] {
	return VectorIterator[T]{
		v:   z,
		pos: math.MaxInt,
	}
}

func (z *Vector[T]) Insert(itr VectorIterator[T], x T) {
	z.InsertBefore(itr.pos, x)
}

func (z *Vector[T]) InsertBefore(pos int, x T) {
	if pos <= 0 {
		(*z) = append(Vector[T]{x}, (*z)...)
		return
	}
	if pos >= z.Size() {
		z.PushBack(x)
		return
	}
	(*z) = append((*z)[:pos+1], (*z)[pos:]...)
	(*z)[pos] = x
}

func (z *Vector[T]) Sort(less func(a, b *T) bool) {
	sort.Slice((*z), func(i, j int) bool {
		return less(&(*z)[i], &(*z)[j])
	})
}

func (z *Vector[T]) ToSlice() (ret []T) {
	ret = make([]T, z.Size())
	copy(ret, *z)
	return
}

func (z *Vector[T]) EraseAt(idx int) {
	if idx < 0 || idx >= z.Size() {
		return
	}
	if idx == z.Size()-1 {
		(*z) = (*z)[:z.Size()-1]
		return
	}
	(*z) = append((*z)[:idx], (*z)[idx+1:]...)
}

func (z *Vector[T]) Find(equal func(x *T) bool) VectorIterator[T] {
	for pos := range *z {
		if equal(&(*z)[pos]) {
			return VectorIterator[T]{
				v:   z,
				pos: pos,
			}
		}
	}
	return z.End()
}

/*
Return value
Iterator following the last removed element.
*/
func (z *Vector[T]) Erase(itr VectorIterator[T]) VectorIterator[T] {
	if itr.pos < 0 || itr.pos >= z.Size() {
		return z.End()
	}
	z.EraseAt(itr.pos)
	if itr.pos == z.Size()-1 {
		return z.End()
	}
	return VectorIterator[T]{
		v:   z,
		pos: itr.pos,
	}
}

type VectorSortable[T constraints.Ordered] Vector[T]

func (z VectorSortable[T]) Less(i, j int) bool {
	return z[i] < z[j]
}

func (z VectorSortable[T]) Len() int {
	return len(z)
}

func (z VectorSortable[T]) Swap(i, j int) {
	z[i], z[j] = z[j], z[i]
}

func (z *VectorSortable[T]) EraseAt(idx int) {
	if idx < 0 || idx >= z.Len() {
		return
	}
	(*z) = append((*z)[:idx], (*z)[idx+1:]...)
}

func (z *VectorSortable[T]) ToSlice() (ret []T) {
	ret = make([]T, z.Len())
	copy(ret, *z)
	return
}

type VectorIterator[T any] struct {
	v   *Vector[T]
	pos int
}
