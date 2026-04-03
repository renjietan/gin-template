package types

type AppConfig struct {
	Path      string `toml:"-"`
	StaticDir string // 静态资源目录
	StaticUrl string
}

type SystemConfig struct {
}
