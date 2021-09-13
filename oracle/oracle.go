package oracle

import (
	"dorado/config"
	"errors"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"
	log "github.com/sirupsen/logrus"
)

//var driverName = "oci8" //Oracle 驱动
//var dataSourceName = "ggs/123456@127.0.0.1:1521/ORCL" //数据库账号：ggs，密码：123456，实例服务：ORCL
var engine *xorm.Engine

func Init() error {
	log.Info("Oracle init starting..... ")
	var err error

	dataSourceName := config.Configs.Oracle.Username + "/" +
		config.Configs.Oracle.Password +
		"@" + config.Configs.Oracle.Ip + ":" +
		config.Configs.Oracle.Port +
		"/" + config.Configs.Oracle.Instance
	engine, err = xorm.NewEngine("oci8", dataSourceName)

	if err != nil {
		log.Error("Oracle init fail, err={}", err)
		return err
	}

	if engine == nil {
		log.Error("Oracle init fail engine is nil")
		return errors.New("engine is nil")
	}

	engine.SetMaxIdleConns(50)
	engine.SetMaxOpenConns(10)
	engine.SetConnMaxLifetime(10000)
	err = engine.Ping()
	if err != nil {
		log.Error("Oracle init fail , err = {}", err)
		return err
	}
	log.Info("Oracle init success")
	return nil
}

func GetEngine() *xorm.Engine {
	return engine
}
