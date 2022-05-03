package goredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestRedis(t *testing.T) {

	re := new(RedisMgr)
	RedisInit(0, re)
	//err := re.Rpush("name", "jim2")
	//err := re.Setex("cache-1", 90, "uid")

	//err := re.Hincrby("RPCCARDS:1001_0005", "104", 1)
	defer re.Close()

	v, err := re.Hget("RPCCARDS:1001_0005", "104")

	clik, _ := redis.Int(v, err)
	fmt.Println("获取点击量", clik)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("sucessfull....")

}

func TestRedisMgr_Hincrby(t *testing.T) {
	re := new(RedisMgr)
	RedisInit(0, re)

	// 新增点击量
	err := re.Hincrby("RPCCARDS:1001_0005", "104", 1)
	defer re.Close()

	v, err := re.Hget("RPCCARDS:1001_0005", "104")

	clik, _ := redis.Int(v, err)
	fmt.Println("获取点击量", clik)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("sucessfull....")

}
