package types

type AppConfig struct {
	// 忽略 Path 字段
	Path      string `toml:"-"`
	StaticDir string // 静态资源目录
	StaticUrl string
	Listen    string
	LogDir    string
}

type SystemConfig struct {
}
