package core

import (
	"bytes"
	"os"

	logger2 "example.com/t/logger"
	"example.com/t/types"
	"github.com/BurntSushi/toml"
)

var logger = logger2.GetLogger()

func NewDefaultConfig() *types.AppConfig {
	return &types.AppConfig{
		StaticDir: "./static",
		StaticUrl: "http://localhost:8080/static",
		Listen:    "0.0.0.0:3344",
		LogDir:    "./static/logs",
		MysqlConfig: types.MysqlConfig{
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "123456",
			DataBase: "springboot3",
		},
	}
}

func LoadConfig(configFile string) (*types.AppConfig, error) {
	var config *types.AppConfig
	_, err := os.Stat(configFile)
	if err != nil {
		logger.Info("创建配置文件: ", configFile)
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
