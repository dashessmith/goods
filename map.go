package util

import (
	"reflect"
	"sort"
)

func MapSortedKeys(m interface{}, dstSlice interface{}, less func(i, j int) bool) {
	rm := reflect.ValueOf(m)
	keys := rm.MapKeys()
	if len(keys) <= 0 {
		return
	}
	rd := reflect.ValueOf(dstSlice).Elem()
	rd.Set(reflect.MakeSlice(rd.Type(), 0, 0))
	rd.Set(reflect.Append(rd, keys...))
	sort.Slice(rd.Interface(), less)
}

func MapKeys(m interface{}, dstSlice interface{}, less func(i, j int) bool) {
	rm := reflect.ValueOf(m)
	keys := rm.MapKeys()
	if len(keys) <= 0 {
		return
	}
	rd := reflect.ValueOf(dstSlice).Elem()
	rd.Set(reflect.MakeSlice(rd.Type(), 0, 0))
	rd.Set(reflect.Append(rd, keys...))
}
