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

const (
	Success       = BizCode(0)
	Failed        = BizCode(1)
	NotAuthorized = BizCode(401) // 未授权
)

const (
	OkMsg       = "Success"
	ErrorMsg    = "系统开小差了"
	InvalidArgs = "非法参数或参数解析失败"
	UploadFaild = "文件上传失败"
)
