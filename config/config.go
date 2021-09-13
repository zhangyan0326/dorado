package config

import (
	"github.com/BurntSushi/toml"
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
)

const (
	pathTest    = "./config/test.toml"
	pathDev     = "./config/dev.toml"
	pathPro     = "./config/pro.toml"
	pathDefault = "./config/config.toml"
)

var Path string
var Configs Config
var DefaultConfigs DefaultConfig

type DefaultConfig struct {
	System SystemConfig `json:"system"`
}

type SystemConfig struct {
	Environment string
	Port        string
}

type Switch struct {
	Grpc   string
	Mysql  string
	Oracle string
	Apollo string
	Redis  string
}

type MysqlConfig struct {
	Ip       string
	Host     string
	Username string
	Password string
}

type Apollo struct {
	AppId         string
	Cluster       string
	NamespaceName string
	Ip            string
}

type RedisConfig struct {
	Ip       string
	Host     string
	Password string
	Db       string
}

type Log struct {
	LogUrl     string
	SystemName string
}

type OracleConfig struct {
	Ip       string
	Port     string
	Username string
	Password string
	Instance string
}

type Config struct {
	Apollo Apollo
	Mysql  MysqlConfig
	Redis  RedisConfig
	Log    Log
	Switch Switch
	Oracle OracleConfig
}

func Init(ctx *iris.Application) error {

	log.Info("配置文件初始化开始")

	// 根据系统判定配置文件使用哪个
	_, err := toml.DecodeFile(pathDefault, &DefaultConfigs)
	if DefaultConfigs.System.Environment == "dev" {
		Path = pathDev
	} else if DefaultConfigs.System.Environment == "test" {
		Path = pathTest
	} else if DefaultConfigs.System.Environment == "prod" {
		Path = pathPro
	}

	_, err = toml.DecodeFile(Path, &Configs)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("配置文件初始化成功")
	log.Info(Configs)
	return nil
}
