package model

type LoginLog struct {
	Id      uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Content string `gorm:"column:content;type:varchar(20);uniqueIndex;not null;comment:登录日志"`
	Time    string `gorm:"column:time;type:text;not null"`
}

func (m *LoginLog) TableName() string {
	return "login_log"
}
