package service

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"strings"
	"upserver/internal/pkg/utils"
)

type HarborService struct {
}

var dockerService DockerService

// DealFile
// @description: 处理解压后的文件
// @param: projectName harbor项目名称，如果不传则使用library
// @param: dirPath 镜像所在宿主机中的目录
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/28 15:12
// @success:
func (hs HarborService) DealFile(projectName, dirPath string) error {
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.WithFields(log.Fields{"dirPath": dirPath, "err": err}).Error("读取目录中的文件错误")
		return err
	}

	//遍历解压后目录中的所有文件
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		fileExt := path.Ext(fileName)
		if fileExt == ".tar" {
			//镜像
			err = hs.DockerPush(projectName, dirPath, fileName)
			if err != nil {
				return err
			}
		} else if fileExt == ".tgz" {
			//helm包

		}

	}
	return nil
}

// DockerPush
// @description: docker镜像推送
// @param: projectName harbor项目名称，如果不传则使用library
// @param: dirPath 镜像所在宿主机中的目录
// @param: fileName 镜像名称 ex:csmp-backend_V3.1.0.tar
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/28 15:06
// @success:
func (hs HarborService) DockerPush(projectName, dirPath, fileName string) error {
	unExtFilename := utils.UnExt(fileName)
	files := strings.Split(unExtFilename, "_")
	//镜像名称
	imageName := files[0] //csmp-backend
	//镜像版本号
	imageVersion := "latest"
	if len(files) > 1 {
		imageVersion = files[1]
	}
	//镜像名称（带版本号）
	imageFullName := imageName + ":" + imageVersion
	//标记的镜像名称
	if projectName == "" {
		projectName = "library"
	}
	tagName := utils.K8sConfig.Harbor.Address + "/" + projectName + "/" + imageFullName
	err := dockerService.PushHarbor(dirPath+"/"+fileName, imageFullName, tagName)
	if err != nil {
		log.WithFields(log.Fields{"dirPath": dirPath, "fileName": fileName, "imageFullName": imageFullName, "tagName": tagName}).
			Error("镜像docker push失败")
	}
	return err
}
