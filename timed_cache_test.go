package goods_test

import (
	"testing"
	"time"

	"github.com/dashessmith/goods"
)

func Test_NewTimedcache(t *testing.T) {
	cache := goods.NewTimedcache()
	calltimes := 0
	fetchf := func() (x interface{}, kt time.Duration, err error) {
		calltimes++
		return `cachevalue`, 10 * time.Millisecond, nil
	}
	v, err := cache.Get("", fetchf)
	goods.AssertNoError(t, err)
	goods.AssertEqual(t, v, `cachevalue`)
	goods.AssertEqual(t, 1, calltimes)
	_, _ = cache.Get("", fetchf)
	goods.AssertNoError(t, err)
	goods.AssertEqual(t, v, `cachevalue`)
	goods.AssertEqual(t, 1, calltimes)
	_, _ = cache.Get("", fetchf)
	goods.AssertNoError(t, err)
	goods.AssertEqual(t, v, `cachevalue`)
	goods.AssertEqual(t, 1, calltimes)
	time.Sleep(20 * time.Millisecond)
	_, _ = cache.Get("", fetchf)
	goods.AssertNoError(t, err)
	goods.AssertEqual(t, v, `cachevalue`)
	goods.AssertEqual(t, 2, calltimes)
}

func Test_get(t *testing.T) {
	// 56 seconds
	cache := goods.NewTimedcache()
	d := goods.Elapse(func() {
		goods.Together(func(int, int) {
			for i := 0; i < 1000000; i++ {
				_, err := cache.Get("213", func() (x interface{}, kt time.Duration, err error) {
					return 1, time.Hour, nil
				})
				if err != nil {
					t.Fatalf("%v\n", err)
				}
			}
		})
	})
	t.Logf("elapse = %v\n", d)
}

func Test_asyncget(t *testing.T) {
	// 56 seconds
	cache := goods.NewTimedcache()
	x, err := cache.AsyncGet("test", 1, func() (x interface{}, kt time.Duration, err error) {
		time.Sleep(time.Second)
		return 2, time.Minute, nil
	})
	goods.AssertNoError(t, err)
	goods.AssertEqual(t, x, 1)
	time.Sleep(2 * time.Second)
	x, err = cache.AsyncGet("test", 2, func() (x interface{}, kt time.Duration, err error) {
		time.Sleep(time.Second)
		return 2, time.Minute, nil
	})
	goods.AssertNoError(t, err)
	goods.AssertEqual(t, x, 2)
}
