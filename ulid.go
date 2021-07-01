package util

import (
	"io"
	"math/rand"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

var _entropy io.Reader

func init() {
	_entropy = ulid.Monotonic(rand.New(rand.NewSource(int64(ulid.Timestamp(time.Now())))), 0)
}

func ULID() string {
	for {
		u, err := ulid.New(ulid.Timestamp(time.Now()), _entropy)
		if err != nil {
			continue
		}
		now := time.Now()
		return strconv.FormatInt(now.UnixNano(), 10) + "_" + u.String()
	}
}
