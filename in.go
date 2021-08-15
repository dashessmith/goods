package goods

import "reflect"

func In(x interface{}, them ...interface{}) bool {
	for _, him := range them {
		if reflect.DeepEqual(x, him) {
			return true
		}
		rt := reflect.TypeOf(him)
		if rt.Kind() == reflect.Slice {
			rv := reflect.ValueOf(him)
			L := rv.Len()
			for i := 0; i < L; i++ {
				if reflect.DeepEqual(x, rv.Index(i).Interface()) {
					return true
				}
			}
		}
	}
	return false
}
