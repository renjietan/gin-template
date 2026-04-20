package core

import (
	"fmt"

	"example.com/t/types"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/sirupsen/logrus"
)

type NacosSerivce struct {
	ConfigClient config_client.IConfigClient
	NamingClient naming_client.INamingClient
	app_config   *types.AppConfig
}

func NewNacosSerivce(config *types.AppConfig, l *logrus.Logger) *NacosSerivce {
	// Nacos 服务端配置
	sc := []constant.ServerConfig{
		{
			IpAddr:      config.NacosConfig.Nacos_Host,
			Port:        config.NacosConfig.Nacos_Port,
			ContextPath: config.NacosConfig.Nacos_ContextPath,
		},
	}
	// 客户端配置
	cc := constant.ClientConfig{
		Username:            config.NacosConfig.Nacos_UserName,
		Password:            config.NacosConfig.Nacos_PassWord,
		NamespaceId:         config.NacosConfig.Nacos_NameSpaceId,
		TimeoutMs:           config.NacosConfig.Nacos_TimeoutMs,
		NotLoadCacheAtStart: config.NacosConfig.Nacos_NotLoadCacheAtStart,
		LogDir:              config.NacosConfig.Nacos_LogDir,
		CacheDir:            config.NacosConfig.Nacos_CacheDir,
		LogLevel:            config.NacosConfig.Nacos_LogLevel,
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize:    100,  // 单个日志文件最大尺寸 (MB)
			MaxAge:     30,   // 保留的旧日志文件最大天数
			MaxBackups: 7,    // 保留的旧日志文件最大数量
			LocalTime:  true, // 是否使用本地时间格式化备份文件名
			Compress:   true, // 是否压缩旧日志文件
		},
	}
	// 创建配置客户端
	configClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		panic(fmt.Errorf("创建nacos客户端失败: %w", err))
		return nil
	}
	// 创建服务注册客户端
	namingClient, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		panic(fmt.Errorf("创建naming客户端失败: %w", err))
	}
	return &NacosSerivce{
		ConfigClient: configClient,
		NamingClient: namingClient,
		app_config:   config,
	}
}

// LoadAndWatchConfig 首次拉取配置，并启动监听协程
func (ns *NacosSerivce) LoadAndWatchConfig() error {
	dataId := ns.app_config.NacosConfig.Nacos_DataId
	groupName := ns.app_config.NacosConfig.Nacos_GroupName
	// 首次获取配置
	content, err := ns.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  groupName,
	})
	if err != nil {
		return fmt.Errorf("获取配置失败: %w", err)
	}
	fmt.Println(content)
	//// 解析并更新全局配置
	//if err := updateConfig(content); err != nil {
	//	return fmt.Errorf("parse initial config failed: %w", err)
	//}

	// 监听配置变更
	err = ns.ConfigClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  groupName,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("=======================")
		},
	})
	if err != nil {
		return fmt.Errorf("listen config failed: %w", err)
	}
	return nil
}
