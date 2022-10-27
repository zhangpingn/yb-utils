// Copyright (c) 2023 YaoBase
// 描述:	初始化日志组件
// 创建者: 张平
// 创建时间: 2022/5/26 14:39

package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// LogData 日志参数相关结构体信息
type LogData struct {
	LogFilename string // 日志文件路径
	LogLevel    string // 日志级别
	MaxSize     int
	MaxBackups  int
	MaxAge      int
	Compress    bool
}

// InitLog
// 描述：初始化日志
// 创建者：张平
// 创建时间：2022/5/26 17:13
// 参数：
//	 [in]
//		无
//   [out]
//		无
//   [in/out]
//		无
func (logData *LogData) InitLog() {
	hook := lumberjack.Logger{
		Filename:   logData.LogFilename,
		MaxSize:    logData.MaxSize,
		MaxBackups: logData.MaxBackups,
		MaxAge:     logData.MaxAge,
		Compress:   logData.Compress,
	}

	var logLevel = zap.InfoLevel
	switch logData.LogLevel {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "fatal":
		logLevel = zap.FatalLevel
	case "panic":
		logLevel = zap.PanicLevel
	}

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(&hook),
		logLevel,
	)
	Logger = zap.New(core)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}
