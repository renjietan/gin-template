package core

import (
	"bytes"
	"os"
	"path/filepath"
	"time"

	"example.com/t/types"
	"github.com/BurntSushi/toml"
)

func NewDefaultConfig() *types.AppConfig {
	path_wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &types.AppConfig{
		Debug:     false,
		AppName:   "gin-template",
		StaticDir: "./static",
		StaticUrl: "http://localhost:8080/static",
		Listen:    "0.0.0.0:3344",
		MysqlConfig: types.MysqlConfig{
			MysqlHost:     "127.0.0.1",
			MysqlPort:     "3306",
			MysqlUsername: "root",
			MysqlPassword: "123456",
			MysqlDataBase: "springboot3",
		},
		LoggerConfig: types.LoggerConfig{
			LoggerLevel:  "debug",
			LoggerMaxAge: 30 * 24 * time.Hour, // 保留30天
			//FilePath: "static/logs/app.log",
			LoggerFilePath:     "static/logs/app",
			LoggerRotationTime: 24 * time.Hour, // 每天切割
		},
		UploadConfig: types.UploadConfig{
			UploadPath: "static/uploads",
		},
		NacosConfig: types.NacosConfig{
			Nacos_Host:                 "0.0.0.0",
			Nacos_Port:                 8848,
			Nacos_GrpcPort:             9848,
			Nacos_ContextPath:          "/nacos",
			Nacos_UserName:             "nacos",
			Nacos_PassWord:             "123456",
			Nacos_GroupName:            "DEFAULT_GROUP",
			Nacos_NameSpaceId:          "nacos-namespace-id",
			Nacos_DataId:               "nacos-dataId",
			Nacos_LogDir:               "/static/logs/nacos/log",
			Nacos_CacheDir:             "/static/logs/nacos/cache",
			Nacos_NotLoadCacheAtStart:  true,
			Nacos_RotateTime:           "24h",
			Nacos_MaxAge:               30,
			Nacos_MaxSize:              100,
			Nacos_LogLevel:             "info",
			Nacos_TimeoutMs:            500,
			Nacos_UpdateCacheWhenEmpty: true,
			Nacos_MaxBackups:           30,
			Nacos_Compress:             true,
			Nacos_LocalTime:            true,
		},
		RedisConfig: types.RedisConfig{
			Redis_Host:        "0.0.0.0",
			Redis_Port:        6379,
			Redis_Password:    "",
			Redis_DB:          0,
			Redis_PoolSize:    20,
			Redis_PoolTimeout: 5,
		},
		SqliteConfig: types.SqliteConfig{
			SqlitePath:          filepath.Join(path_wd, "static", "db"),
			SqliteEncryptionKey: "123456",
		},
	}
}

func LoadConfig(configFile string) (*types.AppConfig, error) {
	var config *types.AppConfig
	_, err := os.Stat(configFile)
	if err != nil {
		config = NewDefaultConfig()
		config.Path = configFile
		err := SaveConfig(config)
		if err != nil {
			return nil, err
		}
		return config, nil
	}
	_, err = toml.DecodeFile(configFile, &config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func SaveConfig(config *types.AppConfig) error {
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	if err := encoder.Encode(&config); err != nil {
		return err
	}
	return os.WriteFile(config.Path, buf.Bytes(), 0644)
}
