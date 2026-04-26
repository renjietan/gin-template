//D:/work/go/template/sever/gin-template/example.toml

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

// 定义嵌套结构体 AppConfig，并为所有需要序列化的字段添加 toml 标签
type AppConfig struct {
	Debug   bool   `toml:"debug"`
	AppName string `toml:"app_name"`
	// 忽略 Path 字段
	Path         string `toml:"-"`
	StaticDir    string `toml:"static_dir"`
	StaticUrl    string `toml:"static_url"`
	Listen       string `toml:"listen"`
	LogDir       string `toml:"log_dir"`
	MysqlConfig  MysqlConfig
	LoggerConfig LoggerConfig
	UploadConfig UploadConfig
	NacosConfig  NacosConfig
	RedisConfig  RedisConfig
	SqliteConfig SqliteConfig
}
type MysqlConfig struct {
	MysqlHost     string `toml:"mysql_host"`
	MysqlPort     string `toml:"mysql_port"`
	MysqlUsername string `toml:"mysql_username"`
	MysqlPassword string `toml:"mysql_password"`
	MysqlDataBase string `toml:"mysql_data_base"`
}

type SqliteConfig struct {
	SqlitePath          string `toml:"sqlite_path"`
	SqliteEncryptionKey string `toml:"sqlite_encryption_key"`
}

type LoggerConfig struct {
	LoggerLevel        string        `toml:"logger_level"`
	LoggerFilePath     string        `toml:"logger_file_path"`
	LoggerMaxAge       time.Duration `toml:"logger_max_age"`
	LoggerRotationTime time.Duration `toml:"logger_rotation_time"`
}

type UploadConfig struct {
	UploadPath string `toml:"upload_path"`
}

type NacosConfig struct {
	Nacos_Host                 string `toml:"nacos_host"`                    // Nacos连接的IP
	Nacos_Port                 uint64 `toml:"nacos_port"`                    // Nacos连接的端口号
	Nacos_GrpcPort             uint64 `toml:"nacos_grpc_port"`               // Nacos grpc端口号
	Nacos_ContextPath          string `toml:"nacos_context_path"`            // nacos 访问路径
	Nacos_UserName             string `toml:"nacos_username"`                // Nacos服务端的API鉴权Username
	Nacos_PassWord             string `toml:"nacos_password"`                // Nacos服务端的API鉴权Password
	Nacos_GroupName            string `toml:"nacos_group_name"`              // 组名
	Nacos_NameSpaceId          string `toml:"nacos_namespaceId"`             // 命名空间
	Nacos_DataId               string `toml:"nacos_dataId"`                  // 配置ID
	Nacos_LogDir               string `toml:"nacos_log_dir"`                 // 日志存储路径
	Nacos_CacheDir             string `toml:"nacos_cache_dir"`               // 缓存service信息的目录，默认是当前运行目录
	Nacos_NotLoadCacheAtStart  bool   `toml:"nacos_not_load_cache_as_start"` // 在启动的时候不读取缓存在CacheDir的service信息
	Nacos_RotateTime           string `toml:"nacos_rotate_time"`             // 日志轮转周期，比如：30m, 1h, 24h, 默认是24h
	Nacos_MaxAge               int    `toml:"nacos_max_age"`                 // 日志最大保留的文件数量; 默认 3
	Nacos_MaxSize              int    `toml:"nacos_max_size"`                // 单个日志文件最大尺寸 (MB)
	Nacos_LogLevel             string `toml:"nacos_log_level"`               // 默认日志级别; 值必须是：debug,info,warn,error，默认值是info
	Nacos_TimeoutMs            uint64 `toml:"nacos_timeout_ms"`              // 请求Nacos服务端的超时时间，默认是10000ms
	Nacos_UpdateCacheWhenEmpty bool   `toml:"nacos_update_cache_when_empty"` // 当服务器返回的实例列表为空时，强制更新本地缓存
	Nacos_MaxBackups           int    `toml:"nacos_max_backups"`             // 保留的旧日志文件最大数量
	Nacos_Compress             bool   `toml:"nacos_compress"`                // 是否使用本地时间格式化备份文件名
	Nacos_LocalTime            bool   `toml:"nacos_local_time"`              // 是否压缩旧日志文件
}

type RedisConfig struct {
	Redis_Host        string `toml:"redis_host"`
	Redis_Port        int    `toml:"redis_port"`
	Redis_Password    string `toml:"redis_password"`
	Redis_DB          int    `toml:"redis_db"`
	Redis_PoolSize    int    `toml:"redis_pool_size"`
	Redis_PoolTimeout int    `toml:"redis_pool_timeout"`
}

func main() {
	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 1. 创建一个包含嵌套数据的 AppConfig 实例
	config := &AppConfig{
		Debug:     false,
		AppName:   "gin-template",
		StaticDir: "./static",
		StaticUrl: "http://localhost:8080/static",
		Listen:    "0.0.0.0:3344",
		MysqlConfig: MysqlConfig{
			MysqlHost:     "127.0.0.1",
			MysqlPort:     "3306",
			MysqlUsername: "root",
			MysqlPassword: "123456",
			MysqlDataBase: "springboot3",
		},
		LoggerConfig: LoggerConfig{
			LoggerLevel:        "debug",
			LoggerMaxAge:       30 * 24 * time.Hour, // 保留30天
			LoggerFilePath:     "static/logs/app",
			LoggerRotationTime: 24 * time.Hour, // 每天切割
		},
		UploadConfig: UploadConfig{
			UploadPath: "static/uploads",
		},
		NacosConfig: NacosConfig{
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
		RedisConfig: RedisConfig{
			Redis_Host:        "0.0.0.0",
			Redis_Port:        6379,
			Redis_Password:    "",
			Redis_DB:          0,
			Redis_PoolSize:    20,
			Redis_PoolTimeout: 5,
		},
		SqliteConfig: SqliteConfig{
			SqlitePath:          filepath.Join(basePath, "static", "db"),
			SqliteEncryptionKey: "123456",
		},
	}

	// 2. 将结构体转换为TOML格式 (美化输出)
	fmt.Println("--- 美化后的 TOML 输出 ---")
	if err := marshalToTOML(config); err != nil {
		log.Fatalf("序列化失败: %v", err)
	}

	// 3. 保存为文件以便下一步读取
	const filename = "config.toml"
	if err := saveTOMLToFile(config, filename); err != nil {
		log.Fatalf("保存文件失败: %v", err)
	}
	fmt.Printf("\n已成功将配置保存到 %s\n", filename)

	// 4. 从 TOML 文件中读取并解析回 AppConfig 结构体
	fmt.Println("\n--- 从 TOML 文件读取并解析回 AppConfig ---")
	loadedConfig, err := loadTOMLFromFile(filename)
	if err != nil {
		log.Fatalf("读取文件失败: %v", err)
	}
	fmt.Printf("读取的值：%s", loadedConfig)
}

// marshalToTOML 使用 Encoder 将 AppConfig 序列化为格式化的 TOML 并输出到标准输出。
// 注意：设置 Encoder.Indent 字段来实现美化输出。
func marshalToTOML(config *AppConfig) error {
	encoder := toml.NewEncoder(os.Stdout)
	encoder.Indent = "  " // 设置缩进为两个空格，可根据需求调整
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("无法序列化配置: %w", err)
	}
	return nil
}

// saveTOMLToFile 将 AppConfig 序列化并保存到指定文件。
func saveTOMLToFile(config *AppConfig, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("无法创建文件: %w", err)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	encoder.Indent = "  " // 保持与 marshalToTOML 一致的缩进
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("无法编码配置: %w", err)
	}
	return nil
}

// loadTOMLFromFile 从指定的 TOML 文件解析配置到 AppConfig 结构体中。
func loadTOMLFromFile(filename string) (AppConfig, error) {
	var config AppConfig
	// toml.DecodeFile 可直接将文件内容解析到结构体指针中
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return config, fmt.Errorf("无法解码 TOML 文件: %w", err)
	}
	return config, nil
}
