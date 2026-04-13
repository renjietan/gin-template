package entity

// uniqueIndex;
type ConfigEntity struct {
	BaseEntity
	Name  string `gorm:"column:name;type:varchar(20);not null;comment:配置名称"`
	Value string `gorm:"column:value;type:text;not null"`
}

func (m *ConfigEntity) TableName() string {
	return "config"
}
