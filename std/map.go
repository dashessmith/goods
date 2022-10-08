package std

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Map[K constraints.Ordered, V any] struct {
	keys Vector[K]
	dict map[K]V
}

func (z *Map[K, V]) Insert(p Pair[K, V]) Pair[MapIterator[K, V], bool] {
	z.init()
	if _, ok := z.dict[p.First]; ok {
		return Pair[MapIterator[K, V], bool]{First: MapIterator[K, V]{k: p.First}, Second: false}
	}
	z.dict[p.First] = p.Second
	z.upsertKey(p.First)
	return Pair[MapIterator[K, V], bool]{First: MapIterator[K, V]{k: p.First}, Second: true}
}

/*
Return value
1,2) The bool component is true if the insertion took place and false if the assignment took place. The iterator component is pointing at the element that was inserted or updated
3,4) Iterator pointing at the element that was inserted or updated
*/
func (z *Map[K, V]) InsertOrAssign(p Pair[K, V]) Pair[MapIterator[K, V], bool] {
	z.init()
	inserted := false
	if _, ok := z.dict[p.First]; !ok {
		inserted = true
		z.upsertKey(p.First)
	}
	z.dict[p.First] = p.Second
	return Pair[MapIterator[K, V], bool]{First: MapIterator[K, V]{k: p.First}, Second: inserted}
}

func (z *Map[K, V]) Upsert(pairs ...Pair[K, V]) {
	z.init()
	for _, p := range pairs {
		z.upsertKey(p.First)
		z.dict[p.First] = p.Second
	}
}

func (z *Map[K, V]) Set(key K, value V) {
	z.init()
	z.Upsert(Pair[K, V]{key, value})
}

func (z *Map[K, V]) Erase(key K) {
	z.init()
	delete(z.dict, key)
	z.deleteKey(key)
}

func (z *Map[K, V]) Keys() Vector[K] {
	z.init()
	return z.keys.ToSlice()
}

func (z *Map[K, V]) Get(key K) V {
	return z.dict[key]
}

func (z *Map[K, V]) Clear() {
	z.init()
	z.dict = map[K]V{}
	z.keys = Vector[K]{}
}

func (z *Map[K, V]) Contains(key K) (ok bool) {
	_, ok = z.dict[key]
	return
}

func (z *Map[K, V]) init() {
	if z.dict != nil {
		return
	}
	z.dict = map[K]V{}
}

type MapIterator[K constraints.Ordered, V any] struct {
	k K
}

func (z *Map[K, V]) deleteKey(key K) {
	idx := sort.Search(len(z.keys), func(i int) bool {
		return z.keys[i] == key
	})
	if idx < 0 {
		return
	}
	z.keys.EraseAt(idx)
}

func (z *Map[K, V]) upsertKey(k K) {
	if z.keys.Empty() {
		z.keys.PushBack(k)
		return
	}
	var impl func(s, e int)
	impl = func(s, e int) {
		m := (s + e) / 2
		midk := z.keys[m]
		if midk == k {
			return
		}

		if midk > k {
			if m == 0 {
				z.keys.PushFront(k)
				return
			}
			impl(m, e)
			return
		}
		if m == s {
			z.keys.PushBack(k)
			return
		}
		impl(s, m)
	}
	impl(0, z.keys.Size())
}
