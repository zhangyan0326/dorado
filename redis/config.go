package redis

import (
	"dorado/bizerror"
	"dorado/config"
	"dorado/zlog"
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris/v12"
)

type RedisConn struct {
	Conn redis.Conn
}

var redisConn RedisConn

func Init() error {
	zlog.Info("redis初始化开始")
	//连接服务器
	address := config.Configs.Redis.Ip + ":" + config.Configs.Redis.Host
	if config.DefaultConfigs.System.Environment == "dev" {
		conn, err := redis.Dial("tcp", address)
		if err != nil {
			zlog.Error("redigo 设置失败")
			zlog.Error("redigo.Dial err=", err)
			return err
		}
		redisConn.Conn = conn
	} else {
		options := redis.DialPassword(config.Configs.Redis.Password)
		conn, err := redis.Dial("tcp", address, options)
		if err != nil {
			zlog.Error("redigo 设置失败")
			zlog.Error("redigo.Dial err=", err)
			return err
		}
		redisConn.Conn = conn
	}
	//设置数据库
	_, err := redisConn.Conn.Do("SELECT", config.Configs.Redis.Db)
	if err != nil {
		zlog.Error("redigo 设置失败")
		zlog.Error("redigo.Dial err=", err)
		return err
	}
	zlog.Info("set redis db " + config.Configs.Redis.Db)
	zlog.Info("redis初始化成功")
	return nil
}

func GetClient() (redis.Conn, error) {
	if redisConn.Conn == nil {
		zlog.Warn("redis conn is nil")
		return nil, errors.New("redigo connect is nil")
	} else {
		return redisConn.Conn, nil
	}
}

func GetRedisConn(context iris.Context, tokenValue interface{}, tokenKey string, operation string) (string, *bizerror.BizError) {
	zlog.Info("[redis] 开始")
	var replyString string
	conn, errRedis := GetClient()
	if errRedis != nil {
		zlog.Error("获取redis连接失败:", errRedis)
		return "", &bizerror.RedisConnError
	}
	if operation == "SET" {
		reply, errDo := redis.Bytes(conn.Do(operation, tokenKey, tokenValue))
		if errDo != nil {
			zlog.Error("【redis】获取的值:", reply)
			zlog.Error("redis set command error ", reply, errDo)
			return "", &bizerror.RedisConnError
		}
		zlog.Info("[redis] 设置redis成功")
		replyString = string(reply)
	} else if operation == "GET" || operation == "DEL" {
		reply, errDo := redis.Bytes(conn.Do(operation, tokenKey))
		if errDo != nil {
			zlog.Error("【redis】获取的值:", reply)
			zlog.Error("【redis】redis set command error ", reply, errDo)
			return "", &bizerror.RedisConnError
		}
		zlog.Info("[redis] ", operation, "redis成功")
		replyString = string(reply)
	}

	zlog.Info("[redis] 结束")
	return string(replyString), nil
}
