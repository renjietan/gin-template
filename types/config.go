package types

import "time"

type AppConfig struct {
	// 忽略 Path 字段
	Path      string `toml:"-"`
	StaticDir string // 静态资源目录
	StaticUrl string
	Listen    string
	LogDir    string
	MysqlConfig
	LoggerConfig
}

type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DataBase string
}

// Logger Config 日志配置
type LoggerConfig struct {
	Level        string        // 日志级别 debug/info/warn/error
	FilePath     string        // 日志文件路径，支持时间格式，如 "logs/app-%Y%m%d.log"
	MaxAge       time.Duration // 日志保留时长，如 30*24*time.Hour
	RotationTime time.Duration // 切割时间间隔，如 24*time.Hour
}
