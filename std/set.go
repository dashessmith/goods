package std

import "golang.org/x/exp/constraints"

type Set[T constraints.Ordered] map[T]bool

func (z *Set[T]) Empty() bool {
	return len(*z) <= 0
}

func (z *Set[T]) Size() int {
	return len(*z)
}

func (z *Set[T]) Clear() {
	(*z) = Set[T]{}
}

func (z *Set[T]) Insert() {

}
