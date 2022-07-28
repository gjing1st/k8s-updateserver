package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"upserver/internal/pkg/k8s"
	"upserver/internal/pkg/model"
)

type K8sService struct {
}

// GetAppAndVersion
// @description: 获取k8s命名空间下应用和版本信息
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/28 11:11
// @success:
func (ks K8sService) GetAppAndVersion(workspace, namespace string) (res model.K8sAppAndVersion, err error) {
	//项目对应的列表
	apps, err1 := k8s.GetApp(workspace, namespace)
	if err1 != nil {
		err = err1
		log.WithFields(log.Fields{"err": err, "namespace": namespace, "workspace": workspace}).
			Error("获取k8s应用列表错误")
		return
	}
	if len(apps.Items) < 1 {
		return res, errors.New("应用为空")
	}

	res.App = *apps
	appid := apps.Items[0].App.AppId
	//TODO 遍历每个应用获取其对应的版本号
	versionRes, err2 := k8s.GetVersions(appid)
	if err2 != nil {
		err = err2
		log.WithFields(log.Fields{"err": err, "namespace": namespace, "workspace": workspace, "appid": appid, "versionRes": versionRes}).
			Error("获取应用版本信息错误")
		return
	}
	if len(versionRes.Items) < 1 {
		return res, errors.New("版本为空")
	}
	res.Version = *versionRes
	return
}
