package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

var chid chan string = make(chan string, 1024)

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
}

func ULID() string {
	return <-chid
}
