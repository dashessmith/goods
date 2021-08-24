package goods

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandScaleIntWithout(scale int, them ...int) int {
	if scale < 1 {
		panic("scale should >= 1")
	}
	var res int = 0
	for scale > 0 {
		n := rand.Intn(10)
		if In(n, them) {
			continue
		}
		if res <= 0 && n <= 0 {
			continue
		}
		res = res*10 + n
		scale--
	}
	return res
}

func RandScaleInt(scale int) int {
	if scale < 1 {
		panic("scale should >= 1")
	}
	high := int(math.Pow(10, float64(scale)))
	low := high / 10
	high -= low
	return rand.Intn(high) + low
}

const randstr = `abcdefghijklmnopqrstuvwxyz_ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`

func RandScaleStr(scale int) string {
	tmp := make([]rune, scale)
	for i := 0; i < scale; i++ {
		tmp[i] = rune(randstr[rand.Intn(len(randstr))])
	}
	return string(tmp)
}
