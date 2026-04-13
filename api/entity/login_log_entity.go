package entity

type LoginLog struct {
	BaseEntity
	Content string `gorm:"column:content;type:varchar(20);uniqueIndex;not null;comment:登录日志"`
	Time    string `gorm:"column:time;type:text;not null"`
}

func (m *LoginLog) TableName() string {
	return "login_log"
}
