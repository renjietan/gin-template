package service

import (
	"example.com/t/types"
	"example.com/t/utility"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// NacosSerivce /**
/**
 * @FILE   nacosService
 * @AUTHOR TAN
 * @DESCRIPTION Nacos 服务
 * @DATE 16:00:32 CST 2026-04-21
 **/
type NacosSerivce struct {
	ServiceConfig []constant.ServerConfig
	ClientConfig  constant.ClientConfig
	ConfigClient  config_client.IConfigClient
	NamingClient  naming_client.INamingClient
	app_config    *types.AppConfig
	Contents      map[string]interface{}
}

type Nacos_err map[string]interface{}

func NewNacosSerivce(config *types.AppConfig) (*NacosSerivce, Nacos_err) {
	// Nacos 服务端配置
	sc := []constant.ServerConfig{
		{
			IpAddr:      config.NacosConfig.Nacos_Host,
			Port:        config.NacosConfig.Nacos_Port,
			ContextPath: config.NacosConfig.Nacos_ContextPath,
			GrpcPort:    config.NacosConfig.Nacos_GrpcPort,
		},
	}
	// 客户端配置
	cc := constant.ClientConfig{
		Username:             config.NacosConfig.Nacos_UserName,
		Password:             config.NacosConfig.Nacos_PassWord,
		NamespaceId:          config.NacosConfig.Nacos_NameSpaceId,
		TimeoutMs:            config.NacosConfig.Nacos_TimeoutMs,
		NotLoadCacheAtStart:  config.NacosConfig.Nacos_NotLoadCacheAtStart,
		LogDir:               config.NacosConfig.Nacos_LogDir,
		CacheDir:             config.NacosConfig.Nacos_CacheDir,
		LogLevel:             config.NacosConfig.Nacos_LogLevel,
		UpdateCacheWhenEmpty: config.NacosConfig.Nacos_UpdateCacheWhenEmpty,
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize:    config.NacosConfig.Nacos_MaxSize,    // 单个日志文件最大尺寸 (MB)
			MaxAge:     config.NacosConfig.Nacos_MaxAge,     // 保留的旧日志文件最大天数
			MaxBackups: config.NacosConfig.Nacos_MaxBackups, // 保留的旧日志文件最大数量
			LocalTime:  config.NacosConfig.Nacos_LocalTime,  // 是否使用本地时间格式化备份文件名
			Compress:   config.NacosConfig.Nacos_Compress,   // 是否压缩旧日志文件
		},
	}
	// 创建配置客户端
	configClient, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		return nil, parseInfo(config, "创建配置客户端失败: ", err.Error())
	}
	// 创建服务注册客户端
	namingClient, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err != nil {
		return nil, parseInfo(config, "创建服务注册客户端失败: ", err.Error())
	}
	return &NacosSerivce{
		ConfigClient: configClient,
		NamingClient: namingClient,
		app_config:   config,
	}, nil
}

// LoadAndWatchConfig 首次拉取配置，并启动监听协程
func (ns *NacosSerivce) LoadAndWatchConfig() Nacos_err {
	dataId := ns.app_config.NacosConfig.Nacos_DataId
	groupName := ns.app_config.NacosConfig.Nacos_GroupName
	// 首次获取配置
	content, err := ns.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  groupName,
	})
	if err != nil {
		return parseInfo(ns.app_config, "首次获取配置失败", err.Error())
	}
	s, parseErr := utility.JsonStrToMap(content)
	if parseErr != nil {
		return parseInfo(ns.app_config, "首次获取配置时，字符串转map失败", err.Error())
	}
	ns.Contents = s
	// 监听配置变更
	err = ns.ConfigClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  groupName,
		OnChange: func(namespace, group, dataId, data string) {
			s, parseErr := utility.JsonStrToMap(content)
			if parseErr != nil {
				//return parseInfo(ns.app_config, "首次获取配置时，字符串转map失败", err.Error())
			}
			ns.Contents = s
		},
	})
	if err != nil {
		return parseInfo(ns.app_config, "开启nacos监听器失败", err.Error())
	}
	return nil
}

// 服务注册
func (ns *NacosSerivce) RegisterService() Nacos_err {
	serviceName := ns.app_config.AppName
	host := ns.app_config.NacosConfig.Nacos_Host
	port := ns.app_config.NacosConfig.Nacos_Port
	// 获取本机可用 IP（实际中可能需要配置或从环境变量获取）
	_, err := ns.NamingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        port,
		ServiceName: serviceName,
		Weight:      100,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"gin-version": "1.9.0"},
	})
	if err != nil {
		return parseInfo(ns.app_config, "服务注册失败", err.Error())
	}
	return nil
}

// deregisterService 优雅关闭时注销服务
func (ns *NacosSerivce) deregisterService() error {
	serviceName := ns.app_config.AppName
	host := ns.app_config.NacosConfig.Nacos_Host
	port := ns.app_config.NacosConfig.Nacos_Port
	_, err := ns.NamingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          host,
		Port:        port,
		ServiceName: serviceName,
		Ephemeral:   true,
	})
	if err != nil {
		return err
	}
	return nil
}

// parseError /**
/**
 * @FILE   nacosService
 * @AUTHOR TAN
 * @DESCRIPTION
 * @DATE 10:58:01 CST 2026-04-22
 * @PARAM
 * @RETURN
 **/
func parseInfo(config *types.AppConfig, title string, err interface{}) Nacos_err {
	errors := map[string]interface{}{
		"namespace-id": config.NacosConfig.Nacos_NameSpaceId,
		"data-id":      config.NacosConfig.Nacos_DataId,
		"group-name":   config.NacosConfig.Nacos_GroupName,
		"title":        title,
		"errors":       err,
	}
	return errors
}
