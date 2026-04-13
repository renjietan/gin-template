package dto

type PagerDTO struct {
	PageNum  int    `json:"page_num"  example:"1"`
	PageSize int    `json:"page_size"  example:"10"`
	Order    string `json:"order" example:"ASC" enums:"ASC|DESC"`
}
