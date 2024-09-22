package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"
)

func init() {
	//设置日志格式为json格式 并且设置时间戳
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置是否显示文件名和行号
	logrus.SetReportCaller(false)

	// 设置日志级别（可以根据需要更改日志级别）
	//logrus.SetLevel(logrus.InfoLevel)
}

// 自定义打印日志方法 在想要调用的接口中调用
func Write(msg string, filename string) {
	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg) // 以info等级保存到文件中
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args)
}

func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.ErrorLevel, "panic")
	logrus.WithFields(fields).Panic(args)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(args)
}

func Trace(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.FatalLevel, "trace")
	logrus.WithFields(fields).Trace(args)
}

// 检查文件夹是否存在 不存在则创建
func setOutPutFile(level logrus.Level, logName string) {
	logdir := "./runtime/log"

	if _, err := os.Stat(logdir); os.IsNotExist(err) {
		// 检查日志文件夹是否存在 不存在则创建
		err = os.MkdirAll(logdir, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", logdir, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	filename := path.Join(logdir, logName+"_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file err", err)
	}

	// 将日志信息写入文件
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

// 记录访问日志
func LoggerToFile() gin.LoggerConfig {
	logdir := "./runtime/log"
	if _, err := os.Stat(logdir); os.IsNotExist(err) {
		// 检查日志文件夹是否存在 不存在则创建
		err = os.MkdirAll(logdir, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", logdir, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	filename := path.Join(logdir, "success_"+timeStr+".log")

	os.Stderr, _ = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s\" %s\n",
				params.TimeStamp.Format("2006-01-02 15:04:05"),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}

	return conf
}

func Recover(c *gin.Context) {
	defer func() {
		logdir := "./runtime/log"
		if err := recover(); err != nil {
			if _, errDir := os.Stat(logdir); os.IsNotExist(errDir) {
				// 检查日志文件夹是否存在 不存在则创建
				err = os.MkdirAll(logdir, 0777)
				if err != nil {
					panic(fmt.Errorf("create log dir '%s' error: %s", logdir, errDir))
				}
			}

			timeStr := time.Now().Format("2006-01-02")
			filename := path.Join(logdir, "error_"+timeStr+".log")

			f, errFile := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errFile != nil {
				fmt.Println(errFile)
			}
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic error time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})
			// 终止后续接口调用 如果注释掉后，recover到异常仍会继续执行后续代码
			c.Abort()

		}
	}()
	c.Next()
}
