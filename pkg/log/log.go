// log日志
package log

import (
	"admin-go-api/common/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var log *logrus.Logger
var logToFile *logrus.Logger

// 日志文件名
var loggerFile string

func setLogFile(file string) {
	loggerFile = file
}

// init
func init() {
	setLogFile(filepath.Join(config.Config.Log.Path, config.Config.Log.Name))
}

// 方法调用
func Log() *logrus.Logger {
	if config.Config.Log.Model == "file" {
		return logFile()
	} else {
		if log == nil {
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}
			log.SetLevel(logrus.DebugLevel)
		}
	}
	return log
}

// 日志方法
func logFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()
		logToFile.SetLevel(logrus.DebugLevel)
		//返回写日志对象logWriter
		logWriter, _ := rotatelogs.New(
			//分割后的文件名称
			loggerFile+".%Y%m%d.log",
			//设置最大保存时间
			rotatelogs.WithMaxAge(30*24*time.Hour),
			//设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		//设置时间格式
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
		//新增Hook
		logToFile.AddHook(lfHook)
	}
	return logToFile
}
