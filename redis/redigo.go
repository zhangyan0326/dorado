package redis

import (
	"dorado/zlog"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
)

func Get(key string) interface{} {
	redisClient, err := GetClient()
	if err != nil {
		log.Error("获取redis客户端失败")
		return nil
	}
	reply, err := redisClient.Do("GET", key)
	if err != nil {
		log.Warning("获取redis缓存失败，key=", key)
		return nil
	}
	return reply
}

func GetString(key string) string {
	redisClient, err := GetClient()
	if err != nil {
		log.Error("获取redis客户端失败")
		return ""
	}
	reply, err := redis.String(redisClient.Do("GET", key))
	if err != nil {
		log.Warning("获取redis缓存失败，key=", key)
		return ""
	}
	return reply
}

func GetMap(key string) map[string]string {
	redisClient, err := GetClient()
	if err != nil {
		log.Error("获取redis客户端失败")
		return nil
	}
	reply, err := redis.StringMap(redisClient.Do("GET", key))
	if err != nil {
		log.Error("获取redis缓存失败，key=", key)
		return nil
	}
	return reply
}

// redis set方法
func Set(key string, value interface{}) {
	log.Info("开始设置redis中的token键值对")
	redisClient, err := GetClient()
	if err != nil {
		log.Error("获取redis客户端失败")
		return
	}
	log.Info("redis set info [key:", key, "] [value:", value, "]")

	err = redisClient.Send("SET", key, value)
	if err != nil {
		log.Error("设置redis缓存失败，key={},value={}", key, value)
	}
	err = redisClient.Flush()
	if err != nil {
		log.Error("设置redis缓存失败，key={},value={}", key, value)
	}
	log.Info("成功设置redis中的token键值对")
}

// 设置带有效期的redis
// key    键
// value  值
// expire 存在时长，单位秒
func SetExpire(key string, value interface{}, expire string) {
	zlog.Info("设置redis缓存", key, value, expire)
	redisClient, err := GetClient()
	if err != nil {
		log.Error("获取redis客户端失败")
		return
	}
	err = redisClient.Send("SET", key, value, "EX", expire)
	if err != nil {
		log.Error("设置redis缓存失败，key="+key+",value=", value)
	}
	err = redisClient.Flush()
	if err != nil {
		log.Error("设置redis缓存失败，key={},value={}", key)
	}
}

// 删除redis中的key值
func Del(key string) {
	log.Info("开始删除redis中的token键值对")
	redisClient, err := GetClient()
	if err != nil {
		log.Error("获取redis客户端失败")
		return
	}
	err = redisClient.Send("DEL", key)
	if err != nil {
		log.Error("删除redis缓存失败，key={},value={}", key)
	}
	err = redisClient.Flush()
	if err != nil {
		log.Error("删除redis缓存失败，key={},value={}", key)
	}
	log.Info("成功结束删除redis中的token键值对")
}
