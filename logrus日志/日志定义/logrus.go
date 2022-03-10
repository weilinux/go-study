package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

// 创建log实例
var logger = logrus.New()

func main() {
	//TODO 使用sirupsen/logrus操作
	// go get github.com/sirupsen/logrus

	//logrus有7个日志级别，依次是Trace << Debug << Info << Warning << Error << Fatal << Panic
	// Panic：记录日志，然后panic
	// Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出
	// Error：错误日志，需要查看原因
	// Warn：警告信息，提醒程序员注意
	// Info：关键操作，核心流程的日志
	// Debug：一般程序中输出的调试信息
	// Trace：很细粒度的信息，一般用不到

	// 设置日志输出配置 输出格式只有文本和json两种格式
	// logger.SetFormatter(&logrus.TextFormatter{})

	// 以JSON而不是默认的ASCII格式记录
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05", // 定义时间格式
		DisableTimestamp:  false,                 // 输出中禁止自动写入时间戳
		DisableHTMLEscape: false,                 // 输出中禁用html转义
		DataKey:           "data",                // 将自定义字段放在给定的dataKey中
		FieldMap: logrus.FieldMap{ // 自定义默认字段的键名
			logrus.FieldKeyMsg:         "message",   // 错误消息 默认值：msg
			logrus.FieldKeyLevel:       "err_level", // 错误等级 默认值：level
			logrus.FieldKeyTime:        "create_at", // 创建时间 默认值：time
			logrus.FieldKeyLogrusError: "error",     // entry错误 默认值：logrus_error
			logrus.FieldKeyFunc:        "func_name", // 方法名称 默认值：func
			logrus.FieldKeyFile:        "file_name", // 文件名称 默认值：file
		},
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			// frame.PC PC是该帧中位置的程序计数器
			// frame.File File为文件名
			// frame.Line Line为行号
			// frame.Entry 函数的入口点程序计数器；可能是nil
			// frame.Func Func是此调用帧的Func值。这可能是nil
			// frame.Function Function是的包路径限定函数名

			// 去掉文件名前面的路径和后面的参数
			fileName := path.Base(frame.File)
			return frame.Function, fileName
		}, // 自定义Caller的返回
		PrettyPrint: true, // 缩进所有json日志
	})

	// 仅记录警告严重性或更高级别的错误
	logger.SetLevel(logrus.WarnLevel)

	// 日志定位行号 如：func=main.main file="./xxx.go:38"
	logger.SetReportCaller(true)

	// 已追加模式打开时可以更安全的写入，自动调用互斥锁（线程安全）
	logger.SetNoLock()

	// 打开文件
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Error(err)
	}

	//关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Error(err)
		}
	}(file)

	// 将日志输出在文件中
	//logger.SetOutput(file)

	// 输出到标准输出，而不是默认的标准错误
	// 将日志输出到控制台上、文件中
	logger.SetOutput(io.MultiWriter(os.Stdout, file))

	// 日志写入
	// logrus鼓励通过Field机制进行精细化的、结构化的日志记录，而不是通过冗长的消息来记录日志
	entry := logger.WithFields(logrus.Fields{
		"event": "event",
		"topic": "topic",
		"key":   "key",
	})
	entry.Log(logrus.FatalLevel, "fatal to send event2")
	entry.Fatal("fatal to send event1")
}
