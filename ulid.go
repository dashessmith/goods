package goods

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	chid      chan string = make(chan string, 1024)
	chidint64 chan int64  = make(chan int64, 1024)
	ulidOnce  sync.Once
)

func ulidInit() {
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
	ulidOnce.Do(ulidInit)
	return <-chid
}

func ULIDInt64() int64 {
	ulidOnce.Do(ulidInit)
	return <-chidint64
}
