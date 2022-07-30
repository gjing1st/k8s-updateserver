package constant

const (
	RequestParamErr   = "invalid request param"
	ServerErr         = "server error"
	RequestErrExt     = "上传tar包格式错误"
	UsernameHasExists = "该用户名已存在"
	TokenExpired      = "token is expired or error"
	UserNotExisted    = "user not existed"
	AdminPasswordErr  = "password error"
	AdminDisabledErr  = "admin account disabled"
)

const (
	DockerLoadErr  = "导入镜像失败"
	DockerTagErr   = "镜像标记失败"
	DockerPushErr  = "镜像推送失败"
	DockerLoginErr = "私有仓库登录失败"
	AddHelmRepoErr = "添加helm私有仓库失败"
	HelmPushErr    = "helm推送失败"
)
