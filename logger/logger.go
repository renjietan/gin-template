package logger

// * +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// * Copyright 2023 The Geek-AI Authors. All rights reserved.
// * Use of this source code is governed by a Apache-2.0 license
// * that can be found in the LICENSE file.
// * @Author yangjian102621@163.com
// * +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if sugarLogger != nil {
		return sugarLogger
	}

	logLevel := zap.NewAtomicLevelAt(getLogLevel(os.Getenv("GEEKAI_LOG_LEVEL")))
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	fileCore := zapcore.NewCore(encoder, writerSyncer, logLevel)
	consoleOutput := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(
		encoder,
		consoleOutput,
		logLevel,
	)
	core := zapcore.NewTee(fileCore, consoleCore)
	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
	return sugarLogger
}

// core 三个参数之  编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "static/logs/app.log", // 日志文件路径
		MaxSize:    100,                   // 单个文件最大 100 MB
		MaxBackups: 30,                    // 最多保留 30 个旧文件
		MaxAge:     30,                    // 最多保留 30 天
		Compress:   true,                  // 压缩旧文件
		LocalTime:  true,                  // 使用本地时间命名备份
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLogLevel(level string) zapcore.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return zapcore.DebugLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
