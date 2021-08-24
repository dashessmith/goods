package goods

import (
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
