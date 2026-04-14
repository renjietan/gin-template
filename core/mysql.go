package core

import (
	"time"

	"example.com/t/core/logger"
	"example.com/t/types"
	"example.com/t/utility"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NewGormLogger 返回适配后的 logger.Interface
func NewGormLogger(log *logrus.Logger) *logger.GormLogger {
	// 获取现有的 zap logger（注意 GetLogger 返回的是 *zap.SugaredLogger）
	gormLogger := &logger.GormLogger{
		Logger:        log,
		SlowThreshold: 200 * time.Millisecond, // 慢查询阈值
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

func NewMysql(config *gorm.Config, appConfig *types.AppConfig) (*gorm.DB, error) {
	template := `{{Username}}:{{Password}}@tcp({{Host}}:{{Port}})/{{DataBase}}?charset=utf8mb4&parseTime=True&loc=Local`
	dns := utility.StringByTemplate(template, appConfig.MysqlConfig)
	db, err := gorm.Open(mysql.Open(dns), config)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(32)
	sqlDB.SetMaxOpenConns(512)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, nil
}
