// Path: pkg/utils/errcode
// FileName: ipm.go
// Author: GJing
// Date: 2022/11/17$ 15:24$

package errcode

// 发布平台相关模块
const (
	PublishArticleCode      = 1 * ModuleCode //文章模块
	PublishFileCode         = 2 * ModuleCode //文件模块
	PublishNotificationCode = 3 * ModuleCode //通知模块
)

// 发布平台文章模块错误码
const (
	PublishArticlePublishFailed          = 1 * ErrCode //发布失败
	PublishArticleSaveFailed             = 2 * ErrCode //保存失败
	PublishArticleUploadFailed           = 3 * ErrCode //上传文件失败
	PublishArticleDeleteFailed           = iota + 1    //文章删除失败			4
	PublishArticleQueryFailed                          //查询失败			5
	PublishArticlePushNoticeCreateFailed               //推送通知创建失败		6
	PublishArticleListStandard                         //查询行业规范列表失败	7
	PublishArticleDeleteArticleBatch                   //批量删除文章失败		8
	PublishArticleDeleteArticle                        //删除文章失败			9
	PublishArticleEditArticle                          //编辑文章失败			10
)
