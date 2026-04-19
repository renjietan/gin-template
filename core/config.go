package core

import (
	"bytes"
	"os"
	"time"

	"example.com/t/types"
	"github.com/BurntSushi/toml"
)

func NewDefaultConfig(debug bool) *types.AppConfig {
	return &types.AppConfig{
		Debug:     debug,
		StaticDir: "./static",
		StaticUrl: "http://localhost:8080/static",
		Listen:    "0.0.0.0:3344",
		MysqlConfig: types.MysqlConfig{
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "123456",
			DataBase: "springboot3",
		},
		LoggerConfig: types.LoggerConfig{
			Level:  "debug",
			MaxAge: 30 * 24 * time.Hour, // 保留30天
			//FilePath: "static/logs/app.log",
			FilePath:     "static/logs/app",
			RotationTime: 24 * time.Hour, // 每天切割
		},
		UploadConfig: types.UploadConfig{
			UploadPath: "static/uploads",
		},
	}
}

func LoadConfig(configFile string, debug bool) (*types.AppConfig, error) {
	var config *types.AppConfig
	_, err := os.Stat(configFile)
	if err != nil {
		//logger.Info("创建配置文件: ", configFile)
		config = NewDefaultConfig(debug)
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
