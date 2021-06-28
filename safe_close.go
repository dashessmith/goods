package util

import "reflect"

func SafeClose(x interface{}) {
	defer func() {
		recover()
	}()
	rv := reflect.ValueOf(x)
	if rv.Kind() == reflect.Chan {
		rv.Close()
	}
}
