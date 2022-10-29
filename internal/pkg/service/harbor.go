package service

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os/exec"
	"path"
	"strings"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/utils"
)

type HarborService struct {
}

var dockerService DockerService

func Init() {
	//docker登录
	dockerService.Login()
	//添加私有仓库
	HarborService{}.AddHelmRepo()
}

// DealFile
// @description: 处理解压后的文件
// @param: projectName harbor项目名称，如果不传则使用library
// @param: dirPath 镜像所在宿主机中的目录
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/28 15:12
// @success:
func (hs HarborService) DealFile(projectName, dirPath string) error {
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.WithFields(log.Fields{"dirPath": dirPath, "err": err}).Error("读取目录中的文件错误")
		return err
	}
	var has HarborService
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
			err = has.HelmChartPush(dirPath + "/" + fileName)
			if err != nil {
				return err
			}
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
// @email: gjing1st@gmail.com
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
		projectName = utils.K8sConfig.Harbor.Project
	}
	tagName := utils.K8sConfig.Harbor.Address + "/" + projectName + "/" + imageFullName
	err := dockerService.PushHarbor(dirPath+"/"+fileName, imageFullName, tagName)
	if err != nil {
		log.WithFields(log.Fields{"dirPath": dirPath, "fileName": fileName, "imageFullName": imageFullName, "tagName": tagName}).
			Error("镜像docker push失败")
	}
	return err
}

// AddHelmRepo
// @description: 添加helm私有仓库
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/29 17:24
// @success:
func (hs HarborService) AddHelmRepo() (err error) {
	//helm文件要上传到的仓库地址 http://admin:Harbor12345@core.harbor.dked:30002/chartrepo/csmp
	repoAddress := "http://" + utils.K8sConfig.Harbor.Admin + ":" + utils.K8sConfig.Harbor.Password + "@" +
		utils.K8sConfig.Harbor.Address + "/chartrepo/" + utils.K8sConfig.Harbor.Project
	var stderr bytes.Buffer
	cmd := exec.Command("helm", "repo", "add", constant.HelmRepoName, repoAddress)
	cmd.Stderr = &stderr
	err = cmd.Run()
	//fmt.Println("======",cmd.String())
	if err != nil {
		//镜像推送失败
		log.WithFields(log.Fields{
			"error": fmt.Sprint(err) + ": " + stderr.String(),
			"cmd":   cmd.String(),
		}).Error(constant.AddHelmRepoErr)
		return errors.New(constant.AddHelmRepoErr)
		//panic("添加helm仓库失败")
	}
	return

}

// HelmChartPush
// @description: 推送char包至私有仓库
// @param: charName helm chart 压缩包全路径名称 ex:/home/data/csmp-0.1.0.tgz
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/29 17:30
// @success:
func (hs HarborService) HelmChartPush(charName string) (err error) {
	var stderr bytes.Buffer
	cmd := exec.Command("helm", "cm-push", charName, constant.HelmRepoName)
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		//镜像推送失败
		log.WithFields(log.Fields{
			"error": fmt.Sprint(err) + ": " + stderr.String(),
			"cmd":   cmd.String(),
		}).Error(constant.HelmPushErr)
		panic("helm推送失败")
	}
	return
}
