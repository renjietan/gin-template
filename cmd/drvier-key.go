package cmd

import "fmt"

const (
	Redis  = "redis"
	Sqlite = "sqlite"
	Mysql  = "mysql"
)

type DriverKey string

func (r *DriverKey) UnmarshalFlag(value string) error {
	switch value {
	case Redis, Sqlite, Mysql:
		*r = DriverKey(value)
		return nil
	default:
		return fmt.Errorf("错误的键值: %q, 必须是 %q | %q | %q",
			value, Redis, Sqlite, Mysql)
	}
}
