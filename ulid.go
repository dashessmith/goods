package goods

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	chid      chan string = make(chan string, 1024)
	chidint64 chan int64  = make(chan int64, 1024)
)

func init() {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(int64(ulid.Timestamp(time.Now())))), 0)
	var lastid string
	go func() {
		for {
			u, err := ulid.New(ulid.Timestamp(time.Now()), entropy)
			if err != nil {
				continue
			}
			now := time.Now()
			id := strconv.FormatInt(now.UnixNano(), 10) + "_" + u.String()
			if strings.Compare(id, lastid) > 0 {
				chid <- id
				lastid = id
			}
		}
	}()

	go func() {
		var lastid int64
		for {
			id := time.Now().UnixNano()
			if id > lastid {
				chidint64 <- id
				lastid = id
			}
		}
	}()
}

func ULID() string {
	return <-chid
}

func ULIDInt64() int64 {
	return <-chidint64
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
