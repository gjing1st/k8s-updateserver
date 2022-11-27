// Path: pkg/utils
// FileName: errcode.go
// Author: GJing
// Date: 2022/11/1$ 20:06$

package errcode

// 错误代码100101，其中 10 代表发布平台服务；中间的 01 代表发布平台服务下的文章模块；最后的 01 代表模块下的错误码序号，每个模块可以注册 100 个错误
// 0代表成功
const (
	SuccessCode = 0 //成功返回错误码
	ErrCode     = 1
	ModuleCode  = 100 * ErrCode
	ServerCode  = ModuleCode * 100
)

// 服务代码
const (
	ServerPublishCode   = 25 * ServerCode //发布平台
	ServerStatisticCode = 26 * ServerCode //统计服务
	ServerAlertCode     = 27 * ServerCode //告警管理服务
)
