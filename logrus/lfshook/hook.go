package main

import (
	"bufio"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

// 创建log实例
var logger = logrus.New()

func main() {
	//TODO 使用sirupsen/logrus自定义hook
	// go get github.com/rifflock/lfshook
	// go get github.com/sirupsen/logrus

	// 创建钩子
	hook := lfshook.NewHook(
		// 将日志级别映射到文件路径 (多个级别可以共享一个文件，但多个文件不能用于一个级别)
		lfshook.PathMap{
			logrus.PanicLevel: "panic.log",
			logrus.FatalLevel: "fatal.log",
			logrus.ErrorLevel: "error.log",
			logrus.WarnLevel:  "warn.log",
			logrus.InfoLevel:  "info.log",
			logrus.DebugLevel: "debug.log",
			logrus.TraceLevel: "trace.log",
		}, &logrus.JSONFormatter{
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

	// 设置日志输出配置 输出格式只有文本和json两种格式
	// hook.SetFormatter(&logrus.TextFormatter{})

	// 以JSON而不是默认的ASCII格式记录
	// hook.SetFormatter()

	// 为没有任何定义的输出路径的级别设置默认路径
	// hook.SetDefaultPath("D:\\GoProject\\goStudy\\")

	// 为没有任何已定义writer的级别设置默认writer
	// hook.SetDefaultWriter(io.MultiWriter(os.Stdout, file))

	// 将钩子添加到logger实例
	logger.Hooks.Add(hook)

	// 记录指定的错误级别的错误，使用钩子的情况下也必须设置
	logger.SetLevel(logrus.TraceLevel)

	// 日志定位行号 如：func=main.main file="./xxx.go:38"
	logger.SetReportCaller(true)

	// 已追加模式打开时可以更安全的写入，自动调用互斥锁（线程安全）
	logger.SetNoLock()

	// 阻止logger进行标准化输出 (创建一个空的写对象)
	logger.SetOutput(bufio.NewWriter(&bufio.Writer{}))
	// 另一种写法（打开一个空设备）
	/*src, _ := os.OpenFile(os.DevNull, os.O_WRONLY, os.ModeTemporary)
	logger.SetOutput(bufio.NewWriter(src))*/

	// 日志写入
	// logrus鼓励通过Field机制进行精细化的、结构化的日志记录，而不是通过冗长的消息来记录日志
	entry := logger.WithFields(logrus.Fields{
		"event": "event",
		"topic": "topic",
		"key":   "key",
	})

	// 将日志文件写入定义的路径或使用定义的写入程序 SetDefaultPath、SetDefaultWriter
	//hook.Fire(entry)

	entry.Trace("trace to send event2")
	entry.Fatal("fatal to send event1")
}
