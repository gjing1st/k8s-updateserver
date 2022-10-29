package service

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"sort"
	"strings"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/k8s"
	"upserver/internal/pkg/model"
	"upserver/internal/pkg/utils"
)

type K8sService struct {
}

// GetAppAndVersion
// @description: 获取k8s命名空间下应用和版本信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
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

// UnzipAndPush
// @description: 解压升级包并推送到私有仓库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/3 10:08
// @success:
func (ks K8sService) UnzipAndPush() error {
	if constant.HarborPushed == 1 {
		return nil
	}
	return nil
	rootPath := utils.Config.Path
	fileInfos, err := ioutil.ReadDir(rootPath)
	if err != nil {
		log.WithField("err", err).Error("读取升级包目录错误")
		return errors.New("读取升级包目录错误")
	}
	dirNames := []string{}
	for _, fileInfo := range fileInfos {
		dirName := fileInfo.Name()
		dirNames = append(dirNames, dirName)
	}
	sort.Strings(dirNames)
	//最后上传应用升级包的目录
	latestDir := rootPath + "/" + dirNames[len(dirNames)-1] + "/"
	//获取压缩包名称
	zipInfos, err := ioutil.ReadDir(latestDir)
	if err != nil {
		log.WithField("err", err).Error("读取压缩包目录错误")
		return errors.New("读取压缩包目录错误")
	}
	zipName := zipInfos[0].Name()
	//解压缩升级包
	fmt.Println("zipName", zipName)
	fileExt := path.Ext(zipName)
	if fileExt == ".zip" {
		utils.UnzipDir(latestDir+zipName, latestDir)
	}

	files := strings.Split(zipName, "_")
	//要上传到的harbor项目名称
	projectName := files[0]
	//解压后的路径
	dirPath := latestDir + zipName
	if fileExt == ".zip" {
		dirPath = latestDir + utils.UnExt(zipName)
	}

	fmt.Println("projectName", projectName)
	fmt.Println("dirPath", dirPath)
	//解压后处理解压后的文件
	err = HarborService{}.DealFile(projectName, dirPath)
	if err != nil {
		constant.HarborPushed = 0
	} else {
		constant.HarborPushed = 1
	}

	return err
}
