package mysql

import (
	"dorado/config"
	"dorado/zlog"
	"errors"
	_ "github.com/go-sql-driver/mysql"

	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func Init() error {
	zlog.Info("mysql初始化开始")
	var err error
	dataSourceName := config.Configs.Mysql.Username +
		":" + config.Configs.Mysql.Password +
		"@(" + config.Configs.Mysql.Ip +
		":" + config.Configs.Mysql.Host +
		")/cancer?charset=utf8"
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		zlog.Error("mysql初始化失败 , err = {}", err)
		return err
	}
	if engine == nil {
		zlog.Error("mysql初始化失败 , engine is nil")
		return errors.New("engine is nil")
	}
	engine.SetMaxIdleConns(50)
	engine.SetMaxOpenConns(10)
	engine.SetConnMaxLifetime(10000)
	err = engine.Ping()
	if err != nil {
		zlog.Error("mysql初始化失败 , err = {}", err)
		return err
	}
	zlog.Info("mysql初始化成功")
	return nil
}

func GetEngine() *xorm.Engine {
	return engine
}
