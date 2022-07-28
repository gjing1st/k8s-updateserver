package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/k8s"
	"upserver/internal/pkg/model"
	"upserver/internal/pkg/service"
	"upserver/internal/pkg/utils"
)

type K8sVersionController struct {
	BaseController
}

// GetVersion
// @description: 获取k8s中应用的版本信息，使用项目中的第一个应用
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/28 9:59
// @success:
func (kv K8sVersionController) GetVersion(c *gin.Context) {
	var reqData model.GetVersionRequest
	if err := c.ShouldBindJSON(&reqData); err != nil {

	}
	//前端未传企业空间和项目则使用配置文件中的
	if reqData.Workspace == "" {
		reqData.Workspace = utils.K8sConfig.K8s.Workspace
	}
	if reqData.Namespace == "" {
		reqData.Namespace = utils.K8sConfig.K8s.Namespace
	}
	var k8sService service.K8sService
	appAndVersion, err := k8sService.GetAppAndVersion(reqData.Workspace, reqData.Namespace)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "namespace": reqData.Namespace, "workspace": reqData.Workspace}).
			Error("获取k8s应用列表错误")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//获取版本信息
	var res model.GetVersionResponse
	res.Appid = appAndVersion.App.Items[0].App.AppId
	res.NowVersion = appAndVersion.App.Items[0].Version.Name
	//默认最新版为当前版本
	res.LatestVersion = appAndVersion.App.Items[0].Version.Name
	if len(appAndVersion.Version.Items) > 0 {
		res.LatestVersion = appAndVersion.Version.Items[0].Name
		res.LatestVersionId = appAndVersion.Version.Items[0].VersionId
	}

	c.JSON(http.StatusOK, res)

}

// UpdateVersion
// @description: 更新应用版本信息
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/28 10:31
// @success:
func (kv K8sVersionController) UpdateVersion(c *gin.Context) {
	var reqData model.UpdateVersionRequest
	if err := c.ShouldBindJSON(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, constant.RequestParamErr)
		return
	}
	//前端未传企业空间和项目则使用配置文件中的
	if reqData.Workspace == "" {
		reqData.Workspace = utils.K8sConfig.K8s.Workspace
	}
	if reqData.Namespace == "" {
		reqData.Namespace = utils.K8sConfig.K8s.Namespace
	}
	var k8sService service.K8sService
	appAndVersion, err := k8sService.GetAppAndVersion(reqData.Workspace, reqData.Namespace)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "namespace": reqData.Namespace, "workspace": reqData.Workspace}).
			Error("获取k8s应用列表错误")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// 前端未传应用id和版本id则使用第一个应用和最新版本
	if reqData.Appid == "" {
		reqData.Appid = appAndVersion.App.Items[0].App.AppId
	}
	if reqData.VersionId == "" {
		reqData.VersionId = appAndVersion.Version.Items[0].VersionId

	}
	//获取版本id的values.yaml
	valuesYaml, err := k8s.Files(reqData.Appid, reqData.VersionId)

	if err != nil || valuesYaml == ""{
		log.WithFields(log.Fields{"err": err, "appid": reqData.Appid, "versionId": reqData.VersionId}).
			Error("获取信息错误")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//组装更新应用需要的数据
	upReq := k8s.UpVersionReq{
		reqData.Appid,
		appAndVersion.App.Items[0].Cluster.ClusterId,
		"default",
		valuesYaml,
		appAndVersion.App.Items[0].App.Name,
		reqData.Namespace,
		"admin",
		reqData.VersionId,
		reqData.Workspace,
	}

	msg, err := k8s.UpVersion(upReq)
	if err != nil || msg != "success" {
		log.WithFields(log.Fields{"upReq": upReq, "err": err}).Error("k8s应用更新失败")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
	//TODO 升级失败
	//cluster.status = failed

}
