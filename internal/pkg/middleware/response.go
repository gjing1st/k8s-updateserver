package middleware


// Response 数据结构体
type Response struct {
	// StatusCode 业务状态码
	StatusCode int `json:"status_code"`

	// Message 提示信息
	Message string `json:"message"`

	// Data 数据，用interface{}的目的是可以用任意数据
	Data interface{} `json:"data"`

	// Meta 源数据,存储如请求ID,分页等信息
	Meta Meta `json:"meta"`

	// Errors 错误提示，如 xx字段不能为空等
	Errors []ErrorItem `json:"errors"`
}

// Meta 元数据
type Meta struct {
	RequestId string `json:"request_id"`
	// 还可以集成分页信息等
}

// ErrorItem 错误项
type ErrorItem struct {
	Key   string `json:"key"`
	Value string `json:"error"`
}


