package types

import "time"

type AppConfig struct {
	Debug   bool   `toml:"debug"`
	AppName string `toml:"app_name"`
	// 忽略 Path 字段
	Path      string `toml:"-"`
	StaticDir string `toml:"static_dir"`
	StaticUrl string `toml:"static_url"`
	Listen    string `toml:"listen"`
	LogDir    string `toml:"log_dir"`
	MysqlConfig
	LoggerConfig
	UploadConfig
	NacosConfig
	RedisConfig
	SqliteConfig
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
	Nacos_ContextPath          string `toml:"/nacos"`                        // nacos 访问路径
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
