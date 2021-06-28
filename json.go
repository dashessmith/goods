package util

import "encoding/json"

func JsonString(obj interface{}) string {
	//bs, _ := json.Marshal(obj)
	bs, _ := json.MarshalIndent(obj, "", "    ")
	return string(bs)
}
