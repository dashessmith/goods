package std

import (
	"fmt"
)

type listNode[T any] struct {
	data *T
	next *listNode[T]
	prev *listNode[T]
}

type List[T any] struct {
	size      int
	beginNode *listNode[T]
	endNode   *listNode[T]
}

func (z *List[T]) Empty() bool {
	return z.size <= 0
}

func (z *List[T]) Size() int {
	return z.size
}

func (z *List[T]) PushFront(x T) {
	z.init()
	node := listNode[T]{
		data: &x,
	}
	z.size++
	first1 := z.beginNode.next
	first1.prev = &node
	node.next = first1
	node.prev = z.beginNode
	z.beginNode.next = &node
}

func (z *List[T]) PushBack(x T) {
	z.init()
	node := listNode[T]{
		data: &x,
	}
	z.size++
	last1 := z.endNode.prev
	last1.next = &node
	node.prev = last1
	node.next = z.endNode
	z.endNode.prev = &node
}

func (z *List[T]) Begin() ListIterator[T] {
	z.init()
	return ListIterator[T]{
		node: z.beginNode.next,
	}
}

func listdebug[T any](begin, end ListIterator[T]) {
	// x := []T{}
	for itr := begin; !itr.Equal(end); itr = itr.Next() {
		// x = append(x, *itr.Value())
		fmt.Printf("%v,", *itr.Value())
	}
	fmt.Printf("\n")
}

func (z *List[T]) Reverse() {
	z.init()
	listReverseImpl(z.Begin(), z.End())
}

func listReverseImpl[T any](begin, end ListIterator[T]) {
	if begin.Equal(end) || begin.Next().Equal(end) || begin.Next().Next().Equal(end) {
		return
	}
	for s, e := begin, end.Prev(); !s.Equal(e); s, e = s.Next(), e.Prev() {
		s.node.data, e.node.data = e.node.data, s.node.data
		if s.Next() == e {
			break
		}
	}
}

func (z *List[T]) End() ListIterator[T] {
	z.init()
	return ListIterator[T]{
		node: z.endNode,
	}
}

func (z *List[T]) Front() (x *T) {
	z.init()
	return z.beginNode.next.data
}

func (z *List[T]) Back() (x *T) {
	z.init()
	return z.endNode.prev.data
}

func (z *List[T]) Insert(itr ListIterator[T], x T) ListIterator[T] {
	z.init()
	z.size++
	node := listNode[T]{
		data: &x,
	}
	prevNode := itr.node.prev
	prevNode.next = &node
	itr.node.prev = &node
	node.prev = prevNode
	node.next = itr.node
	return ListIterator[T]{
		node: &node,
	}
}

func (z *List[T]) init() {
	if z.beginNode != nil {
		return
	}
	z.beginNode = &listNode[T]{}
	z.endNode = &listNode[T]{}
	z.beginNode.next = z.endNode
	z.endNode.prev = z.beginNode
	return
}

type ListIterator[T any] struct {
	node *listNode[T]
}

func (z ListIterator[T]) Next() ListIterator[T] {
	if z.node.next == nil {
		return z
	}
	return ListIterator[T]{
		node: z.node.next,
	}
}

func (z ListIterator[T]) Prev() ListIterator[T] {
	if z.node.prev.data == nil {
		return z
	}
	return ListIterator[T]{
		node: z.node.prev,
	}
}

func (z ListIterator[T]) Valid() bool {
	return z.node.data != nil
}

func (z ListIterator[T]) Value() *T {
	return z.node.data
}

func (z ListIterator[T]) Equal(x ListIterator[T]) bool {
	return z.node == x.node
}
