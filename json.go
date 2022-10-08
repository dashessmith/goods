package goods

import "encoding/json"

func JsonStringPretty(obj interface{}) string {
	bs, _ := json.MarshalIndent(obj, "", "    ")
	return string(bs)
}

func JsonString(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return string(bs)
}
