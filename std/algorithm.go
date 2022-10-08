package std

type Iterator[T any] interface {
	Next() Iterator[T]
	Prev() Iterator[T]
	Value() *T
	NotEqual(Iterator[T]) bool
}

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func SortIterator[T any](first, last Iterator[T], less func(a, b Iterator[T]) bool) {
	if !first.Next().NotEqual(last) {
		return
	}
	lessEqual := func(a, b Iterator[T]) bool {
		return !less(b, a)
	}
	pivot := first
	e := last.Prev()
	f := first.Next()
f1:
	for f.NotEqual(e) {
		for ; lessEqual(pivot, e); e = e.Prev() {
			if !f.NotEqual(e) {
				break f1
			}
		}
		for ; lessEqual(f, pivot); f = f.Next() {
			if !f.NotEqual(e) {
				break f1
			}
		}
		Swap(f.Value(), e.Value())
	}
	Swap(first.Value(), pivot.Value())

	SortIterator(first, pivot, less)
	SortIterator(pivot, last, less)
}
