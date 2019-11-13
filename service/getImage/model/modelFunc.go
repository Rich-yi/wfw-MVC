package model

import (
	"github.com/garyburd/redigo/redis"

)

var RedisPool redis.Pool
//连接池
func InitRedis(){
	RedisPool=redis.Pool{
		MaxIdle:20,
		MaxActive:50,
		IdleTimeout:60 * 5,
		Dial: func() (redis.Conn, error){
			return redis.Dial("tcp","127.0.0.1:6379")
		} ,
	}
}
//存储验证码
func SaveImageRnd(uuid,rnd string)error  {
	//链接redis
	conn:=RedisPool.Get()
	//存储验证码
	_,err:=conn.Do("set",uuid,rnd)
	return err
}