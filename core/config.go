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
	pathWd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &types.AppConfig{
		AppName:   "gin-template",
		Debug:     false,
		Version:   "1.0.0",
		StaticDir: "./static",
		Host:      "localhost",
		Port:      3344,
		Mysql: types.MysqlConfig{
			Enable:   false,
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "123456",
			Database: "springboot3",
		},
		Sqlite: types.SqliteConfig{
			Enable:              false,
			BasePath:            filepath.Join(pathWd, "static", "db"),
			SqliteEncryptionKey: "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99",
		},
		Redis: types.RedisConfig{
			Enable:      false,
			Host:        "0.0.0.0",
			Port:        6379,
			Password:    "",
			DB:          0,
			PoolSize:    20,
			PoolTimeout: 5,
		},
		Ns: types.NsConfig{
			Enable:               false,
			Host:                 "0.0.0.0",
			Port:                 8848,
			GrpcPort:             9848,
			ContextPath:          "/nacos",
			UserName:             "nacos",
			PassWord:             "123456",
			GroupName:            "DEFAULT_GROUP",
			NameSpaceId:          "nacos-namespace-id",
			DataId:               "nacos-dataId",
			LogDir:               filepath.Join(pathWd, "static", "nacos", "log"),
			CacheDir:             filepath.Join(pathWd, "static", "nacos", "cache"),
			NotLoadCacheAtStart:  true,
			RotateTime:           "24h",
			MaxAge:               30,
			MaxSize:              100,
			LogLevel:             "info",
			TimeoutMs:            500,
			UpdateCacheWhenEmpty: true,
			MaxBackups:           30,
			Compress:             true,
			LocalTime:            true,
		},
		Logger: types.LoggerConfig{
			Level:        "debug",
			MaxAge:       30 * 24 * time.Hour, // 保留30天
			Filepath:     filepath.Join(pathWd, "static", "logs", "app"),
			RotationTime: 24 * time.Hour, // 每天切割
		},
		Upload: types.UploadConfig{
			Filepath: filepath.Join(pathWd, "static", "uploads"),
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
