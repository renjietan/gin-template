package service

import (
	"fmt"
	"log"
	"net"

	"example.com/t/types"
	"example.com/t/utility"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/sirupsen/logrus"
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
	Content       []byte
}

func NewNacosSerivce(config *types.AppConfig, l *logrus.Logger) *NacosSerivce {
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
	s, _ := utility.JsonStrToMap(content)
	fmt.Println(s)
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

// 服务注册
func (ns *NacosSerivce) RegisterService() error {
	// 获取本机可用 IP（实际中可能需要配置或从环境变量获取）
	ip := getOutboundIP()
	_, err := ns.NamingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        ns.app_config.Nacos_Port,
		ServiceName: ns.app_config.AppName,
		Weight:      100,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"gin-version": "1.9.0"},
	})
	if err != nil {
		return fmt.Errorf("register service failed: %w", err)
	}
	log.Printf("✅ Service registered to Nacos: %s@%s:%d\n", ns.app_config.AppName, ip, ns.app_config.Nacos_Port)
	return nil
}

// deregisterService 优雅关闭时注销服务
func (ns *NacosSerivce) deregisterService(serviceName string) error {
	ip := getOutboundIP()
	_, err := ns.NamingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        ns.app_config.Nacos_Port,
		ServiceName: serviceName,
		Ephemeral:   true,
	})
	if err != nil {
		return err
	}
	log.Println("✅ Service deregistered from Nacos")
	return nil
}

// 获取本机出口 IP（简单实现）
func getOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
