package dto

type ConfigDTO struct {
	Name  string `json:"name" binding:"required" example:"name-张三"` // 使用 example 提供示例值
	Value string `json:"value" binding:"required" example:"value-张三"`
}

type ConfigsDTO struct {
	Items []ConfigDTO `json:"items"`
}

type ConfigListDTO struct {
	PagerDTO
	Name string `json:"name" binding:"required" example:"name"`
}
