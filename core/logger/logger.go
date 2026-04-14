package logger

import (
	"io"
	"os"
	"path"
	"time"

	"example.com/t/types"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var GlobalLog *logrus.Logger // 全局logger实例，方便非DI场景使用

// NewLogger 创建并初始化一个logrus.Logger，支持按日期切割
func NewLogger(appConfig *types.AppConfig) (*logrus.Logger, error) {
	// 解析日志级别
	cfg := appConfig.LoggerConfig
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}

	// 配置rotatelogs（按日期切割）
	writer, err := rotatelogs.New(
		path.Join(path.Dir(cfg.FilePath), "%Y%m%d-"+path.Base(cfg.FilePath)),
		rotatelogs.WithLinkName(cfg.FilePath), // 生成软链指向最新日志
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(cfg.RotationTime)*time.Hour),
	)
	if err != nil {
		return nil, err
	}

	// 同时输出到控制台和文件
	log := logrus.New()
	log.SetLevel(level)
	log.SetOutput(io.MultiWriter(os.Stdout, writer))
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	GlobalLog = log
	return log, nil
}
