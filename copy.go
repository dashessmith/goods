package goods

func CopyInt64(src []int64) (ret []int64) {
	ret = make([]int64, len(src))
	copy(ret, src)
	return
}

func CopyInt(src []int) (ret []int) {
	ret = make([]int, len(src))
	copy(ret, src)
	return
}

func CopyByte(src []byte) (ret []byte) {
	ret = make([]byte, len(src))
	copy(ret, src)
	return
}
