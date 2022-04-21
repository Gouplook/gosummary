package goredis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {

	re := new(RedisMgr)
	RedisInit(0, re)
	//err := re.Rpush("name", "jim2")
	err := re.Setex("cache-1", 90, "uid")
	defer re.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("sucessfull....")

}
