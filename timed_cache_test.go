package util_test

import (
	"testing"
	"time"

	"github.com/dashessmith/util"
	"github.com/stretchr/testify/assert"
)

func Test_NewTimedcache(t *testing.T) {
	cache := util.NewTimedcache()
	calltimes := 0
	fetchf := func() (x interface{}, kt time.Duration, err error) {
		calltimes++
		return `cachevalue`, 10 * time.Millisecond, nil
	}
	v, err := cache.Get("", fetchf)
	assert.NoError(t, err)
	assert.Equal(t, v, `cachevalue`)
	assert.Equal(t, 1, calltimes)
	cache.Get("", fetchf)
	assert.NoError(t, err)
	assert.Equal(t, v, `cachevalue`)
	assert.Equal(t, 1, calltimes)
	cache.Get("", fetchf)
	assert.NoError(t, err)
	assert.Equal(t, v, `cachevalue`)
	assert.Equal(t, 1, calltimes)
	time.Sleep(20 * time.Millisecond)
	cache.Get("", fetchf)
	assert.NoError(t, err)
	assert.Equal(t, v, `cachevalue`)
	assert.Equal(t, 2, calltimes)
}

func Test_get(t *testing.T) {
	cache := util.NewTimedcache()
	d := util.Elapse(func() {
		util.Together(func(int, int) {
			for i := 0; i < 100000000; i++ {
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
