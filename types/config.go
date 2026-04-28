package types

import "time"

type AppConfig struct {
	Debug     bool   `toml:"debug"`
	AppName   string `toml:"app_name"`
	Path      string `toml:"-"`
	StaticDir string `toml:"static_dir"`
	StaticUrl string `toml:"static_url"`
	Listen    string `toml:"listen"`
	LogDir    string `toml:"log_dir"`
	/* 注意这里不可以用匿名嵌入子结构体，否则写入toml文件时, 无法形成树状结构 */
	Mysql  MysqlConfig
	Sqlite SqliteConfig
	Redis  RedisConfig
	Logger LoggerConfig
	Upload UploadConfig
	Ns     NsConfig
}
type MysqlConfig struct {
	Enable   bool   `toml:"enable"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"data_base"`
}

type SqliteConfig struct {
	Enable              bool   `toml:"enable"`
	BasePath            string `toml:"base_path"`
	SqliteEncryptionKey string `toml:"sqlite_encryption_key"`
}

type LoggerConfig struct {
	Level        string        `toml:"level"`
	Filepath     string        `toml:"file_path"`
	MaxAge       time.Duration `toml:"max_age"`
	RotationTime time.Duration `toml:"rotation_time"`
}

type UploadConfig struct {
	Filepath string `toml:"file_path"`
}

type NsConfig struct {
	Host                 string `toml:"host"`                    // Nacos连接的IP
	Port                 uint64 `toml:"port"`                    // Nacos连接的端口号
	GrpcPort             uint64 `toml:"grpc_port"`               // Nacos grpc端口号
	ContextPath          string `toml:"context_path"`            // nacos 访问路径
	UserName             string `toml:"username"`                // Nacos服务端的API鉴权Username
	PassWord             string `toml:"password"`                // Nacos服务端的API鉴权Password
	GroupName            string `toml:"group_name"`              // 组名
	NameSpaceId          string `toml:"namespaceId"`             // 命名空间
	DataId               string `toml:"dataId"`                  // 配置ID
	LogDir               string `toml:"log_dir"`                 // 日志存储路径
	CacheDir             string `toml:"cache_dir"`               // 缓存service信息的目录，默认是当前运行目录
	NotLoadCacheAtStart  bool   `toml:"not_load_cache_as_start"` // 在启动的时候不读取缓存在CacheDir的service信息
	RotateTime           string `toml:"rotate_time"`             // 日志轮转周期，比如：30m, 1h, 24h, 默认是24h
	MaxAge               int    `toml:"max_age"`                 // 日志最大保留的文件数量; 默认 3
	MaxSize              int    `toml:"max_size"`                // 单个日志文件最大尺寸 (MB)
	LogLevel             string `toml:"log_level"`               // 默认日志级别; 值必须是：debug,info,warn,error，默认值是info
	TimeoutMs            uint64 `toml:"timeout_ms"`              // 请求Nacos服务端的超时时间，默认是10000ms
	UpdateCacheWhenEmpty bool   `toml:"update_cache_when_empty"` // 当服务器返回的实例列表为空时，强制更新本地缓存
	MaxBackups           int    `toml:"max_backups"`             // 保留的旧日志文件最大数量
	Compress             bool   `toml:"compress"`                // 是否使用本地时间格式化备份文件名
	LocalTime            bool   `toml:"local_time"`              // 是否压缩旧日志文件
}

type RedisConfig struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Password    string `toml:"password"`
	DB          int    `toml:"db"`
	PoolSize    int    `toml:"pool_size"`
	PoolTimeout int    `toml:"pool_timeout"`
}
