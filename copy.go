package util

func CopyInt64(src []int64) (ret []int64) {
	ret = make([]int64, len(src))
	copy(ret, src)
	return
}
