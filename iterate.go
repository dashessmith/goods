package goods

func All(N int, f func(index int) bool) bool {
	for i := 0; i < N; i++ {
		if !f(i) {
			return false
		}
	}
	return true
}
