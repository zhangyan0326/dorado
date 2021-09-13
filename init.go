package dorado

import (
	"dorado/config"
	"dorado/mysql"
	"dorado/oracle"
	"dorado/redis"
	"dorado/zlog"
	"github.com/kataras/iris/v12"
)

func Init(ctx *iris.Application) error {

	err := config.Init(ctx)
	if err != nil {
		zlog.Error("dorado初始化配置失败", err)
		return err
	}

	if config.Configs.Switch.Redis == "true" {
		err := redis.Init()
		if err != nil {
			zlog.Error("dorado初始化redis失败", err)
			return err
		}
	}

	if config.Configs.Switch.Mysql == "true" {
		err := mysql.Init()
		if err != nil {
			zlog.Error("dorado初始化mysql失败", err)
			return err
		}
	}

	if config.Configs.Switch.Oracle == "true" {
		err := oracle.Init()
		if err != nil {
			zlog.Error("dorado初始化oracle失败", err)
			return err
		}
	}

	if config.Configs.Switch.Apollo == "true" {
		// TODO
	}

	if config.Configs.Switch.Grpc == "true" {
		// TODo
	}

	return nil

}
