package model

import "upserver/internal/pkg/k8s"

//GetVersionRequest 获取应用版本信息请求参数
type GetVersionRequest struct {
	Workspace string `json:"workspace"`
	Namespace string `json:"namespace"`
}

//GetVersionResponse 获取版本返回数据
type GetVersionResponse struct {
	Current int `json:"current"`
	Data    struct {
		Appid           string `json:"appid"`
		NowVersion      string `json:"current"`
		LatestVersion   string `json:"latest"`
		LatestVersionId string `json:"latest_version_id"`
	} `json:"data"`
	Last int `json:"last"`
}


type GetVersionListResponse struct {
	Appid           string `json:"appid"`
	NowVersion      string `json:"current"`
	k8s.VersionResponse
}

//UpdateVersionRequest 更新应用版本请求参数
type UpdateVersionRequest struct {
	Appid string `json:"appid" form:"appid"`
	//Appid     string `json:"appid" form:"appid" binding:"required"`
	VersionId string `json:"version_id" form:"version_id"`
	//VersionId string `json:"version_id" form:"version_id" binding:"required"`
	Workspace string `json:"workspace"`
	Namespace string `json:"namespace"`
}

type K8sAppAndVersion struct {
	App     k8s.AppListResponse
	Version k8s.VersionResponse
}
