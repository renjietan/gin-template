package types

type BizCode int

type Response struct {
	Code     BizCode     `json:"code"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Total    int         `json:"total,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type PaginationResponse struct {
	List       interface{} `json:"list"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

const (
	Success       = BizCode(0)
	Failed        = BizCode(1)
	NotAuthorized = BizCode(401) // 未授权
)

const (
	OkMsg       = "操作成功"
	FailedMsg   = "操作失败"
	ErrorMsg    = "系统开小差了"
	InvalidArgs = "非法参数或参数解析失败"
	UploadFaild = "文件上传失败"
)
