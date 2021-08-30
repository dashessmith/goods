package goods

import (
	"fmt"
	"sync"
)

var onceDict = sync.Map{}

func Once(cond func() bool, then func()) {
	if !cond() {
		return
	}
	caller := GetFrame(1)
	key := fmt.Sprintf("%s(%d)", caller.Function, caller.Line)
	o, _ := onceDict.LoadOrStore(key, &sync.Once{})
	oc := o.(*sync.Once)
	oc.Do(func() {
		then()
	})
}
