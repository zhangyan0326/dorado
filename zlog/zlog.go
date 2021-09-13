package zlog

import (
	"dorado/config"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"path"
	"time"
)

/*
	建议使用这一种
*/

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"-%Y-%m-%d %H:%M.log",
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})
	log.SetReportCaller(true) //将函数名和行数放在日志里面
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.AddHook(lfHook)
}

func getLog() *log.Logger {
	var l = log.New()
	c := new(log.TextFormatter)
	c.TimestampFormat = "2006-01-02 15:04:05"
	c.FullTimestamp = true
	l.SetFormatter(c)
	l.Hooks.Add(NewContextHook())
	return l
}

func Debug(args ...interface{}) {
	var logger = getLog()
	logger.Debug(args)
}

func Info(args ...interface{}) {
	var logger = getLog()
	logger.Info(args)
}

func Warn(args ...interface{}) {
	var logger = getLog()
	logger.Warn(args)
}

func Error(args ...interface{}) {
	var logger = getLog()
	logger.Error(args)
}

func Fatal(args ...interface{}) {
	var logger = getLog()
	logger.Fatal(args)
}

func Log(level log.Level, args ...interface{}) {
	var logger = getLog()
	logger.Log(level, args)
}

func Init(logPath string) {
	ConfigLocalFilesystemLogger(logPath, config.Configs.Log.SystemName, time.Second*60*60*24, time.Second*60*60)
}
