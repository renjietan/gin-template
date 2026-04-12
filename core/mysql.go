package core

import (
	"time"

	"example.com/t/types"
	"example.com/t/utility"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// NewGormLogger 返回适配后的 logger.Interface
// func NewGormLogger() gormLogger.Interface {
// 	// 获取你现有的 zap logger（注意 GetLogger 返回的是 *zap.SugaredLogger）
// 	sugar := logger2.GetLogger()
// 	zapLogger := sugar.Desugar() // 转为 *zap.Logger 供适配器使用

// 	// 配置 GORM 日志行为
// 	config := gormLogger.Config{
// 		SlowThreshold:             200 * time.Millisecond,
// 		LogLevel:                  gormLogger.Warn,
// 		IgnoreRecordNotFoundError: true,
// 		Colorful:                  false, // zap 自带结构化日志，不需要颜色
// 	}

// 	return logger2.NewGormZapAdapter(zapLogger, config)
// }

func NewGormConfig() *gorm.Config {
	return &gorm.Config{
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
