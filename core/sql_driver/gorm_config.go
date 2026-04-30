package sql_driver

import (
	"time"

	"example.com/t/core/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewGormLogger(log *logrus.Logger) *logger.GormLogger {
	gormLogger := &logger.GormLogger{
		Logger:        log,
		SlowThreshold: 200 * time.Millisecond,
	}
	return gormLogger
}

func NewGormConfig(log *logrus.Logger) *gorm.Config {
	return &gorm.Config{
		Logger: NewGormLogger(log),
		//Logger: logger2.GetLogger().Desugar(),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gin_", // 设置表前缀
			SingularTable: false,  // 使用单数表名形式
		},
	}
}
