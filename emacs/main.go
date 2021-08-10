package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {

	}

	wg.Wait()

}

func newRedis() redis.Conn {

	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		MaxIdle:     10,
		MaxActive:   20,
		IdleTimeout: 100,
	}

	return pool.Get()
}

const (
	cacheLockIndex = "redis:distribute:lock"
	AutoExpireTime = 60 * 10

	SetLockSuccess = "OK"
)

func AcquireLock(method, value string, expire time.Duration) bool {

	if setLock(method, value) {
		return true
	}

	return false
}

func ReleaseLock(method, value string) error {

	rd := newRedis()
	defer rd.Close()

	if getLock(method) != value {
		return fmt.Errorf("DeleteLock: fail to match redis key:%v", method)
	}

	cacheIndex := fmt.Sprintf("%v:%v", cacheLockIndex, method)
	_, err := redis.String(rd.Do("DEL", cacheIndex))
	if err != nil {
		return err
	}

	return nil

}

func setLock(method, value string) bool {

	rd := newRedis()
	defer rd.Close()

	cacheIndex := fmt.Sprintf("%v:%v", cacheLockIndex, method)
	msg, err := redis.String(rd.Do("SETNX", cacheIndex, value, "EX", AutoExpireTime))
	if err != nil || msg != SetLockSuccess {
		return false
	}

	return true

}

func getLock(method string) string {

	rd := newRedis()
	defer rd.Close()

	cacheIndex := fmt.Sprintf("%v:%v", cacheLockIndex, method)
	msg, err := redis.String(rd.Do("GET", cacheIndex))
	if err != nil {
		return ""
	}

	return msg

}
