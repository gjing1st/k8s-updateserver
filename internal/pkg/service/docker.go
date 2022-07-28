package service

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/utils"
)

type DockerService struct {
}

// PushHarbor
// @description: 推送镜像到私有仓库
// @param: fullName 镜像的全路径名称
// @param: imageName 镜像名称（带版本号）
// @param: tagName 标记名称（全称）
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/19 19:31
// @success:
func (ds DockerService) PushHarbor(fullName, imageName, tagName string) (err error) {
	err = ds.Login()
	if err != nil {
		return
	}
	err = ds.Load(fullName)
	if err != nil {
		return
	}
	err = ds.Tag(imageName, tagName)
	if err != nil {
		return
	}
	err = ds.Push(tagName)
	return
}

// Login
// @description: 登录私有仓库
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/19 20:15
// @success:
func (ds DockerService) Login() (err error) {
	var stderr bytes.Buffer
	cmd := exec.Command("docker", "login", "-u", utils.K8sConfig.Harbor.Admin, "-p", utils.K8sConfig.Harbor.Password, utils.K8sConfig.Harbor.Address)
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		//镜像推送失败
		log.WithFields(log.Fields{
			"error": fmt.Sprint(err) + ": " + stderr.String(),
			"cmd":   cmd.String(),
		}).Error(constant.DockerLoginErr)
		err = errors.New(constant.DockerLoginErr)
	}
	return
}

// Push
// @description: 镜像推送
// @param: tagName 标记的镜像名称
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/19 19:31
// @success:
func (ds DockerService) Push(tagName string) (err error) {

	var stderr bytes.Buffer
	cmd := exec.Command("docker", "push", tagName)
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		//镜像推送失败
		log.WithFields(log.Fields{
			"error": fmt.Sprint(err) + ": " + stderr.String(),
			"cmd":   cmd.String(),
		}).Error(constant.DockerPushErr)
		err = errors.New(constant.DockerPushErr)
	}
	return
}

// Load
// @description: 镜像导入
// @param: fullName 镜像的全路径名称
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/19 19:12
// @success:
func (ds DockerService) Load(fullName string) (err error) {
	var stderr bytes.Buffer
	cmd := exec.Command("docker", "load", "-i", fullName)
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		//镜像导入失败
		log.WithFields(log.Fields{
			"error": fmt.Sprint(err) + ": " + stderr.String(),
			"cmd":   cmd.String(),
		}).Error(constant.DockerLoadErr)
		err = errors.New(constant.DockerLoadErr)
	}
	return
}

// Tag
// @description: 镜像标记
// @param: imageName 镜像名称（带版本号）
// @param: tagName 标记名称（全称）
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/19 19:15
// @success:
func (ds DockerService) Tag(imageName, tagName string) (err error) {
	var stderr bytes.Buffer
	cmd := exec.Command("docker", "tag", imageName, tagName)
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		//镜像标记失败
		log.WithFields(log.Fields{
			"error": fmt.Sprint(err) + ": " + stderr.String(),
			"cmd":   cmd.String(),
		}).Error(constant.DockerTagErr)
		err = errors.New(constant.DockerTagErr)
	}
	return
}
