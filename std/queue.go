package std

type Container[T any] interface {
	PushBack(T)
	PopBack() T
	Front() *T
	Empty() bool
	Size() int
}

type PriorityQueue[T any, C Container[T]] struct {
	c C
}

func (z *PriorityQueue[T, C]) Empty() bool {
	return z.c.Empty()
}

func (z *PriorityQueue[T, C]) Size() int {
	return z.c.Size()
}
