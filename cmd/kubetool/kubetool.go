// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/8/31$ 15:24$

package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/k8s"
	"upserver/internal/pkg/k8s/csmp"
	"upserver/internal/pkg/service"
	"upserver/internal/pkg/utils"
)

var harborService service.HarborService
var (
	zipName        string
	configFileName string
	projectName    string
	tempPath       string
	initOs         bool
	create         string
)

func init() {
	flag.StringVar(&zipName, "z", "", "zip 指定zip压缩包名称用于推送至harbor私有仓库")
	flag.StringVar(&configFileName, "f", constant.KubeConfigName, "configFileName 配置文件名称用来k8s部署初始化")
	//flag.StringVar(&tempPath, "t", "./kubetool", "tmpPath 解压升级包的目录")
	flag.BoolVar(&initOs, "i", false, "initOS 初始化操作系统")
	flag.StringVar(&create, "c", "", "create config|app 创建k8s应用或配置文件")

}

//go run main.go -z E:/project/csmp/version/csmp_v3.2.0.1.zip
func main() {
	flag.Parse()
	var err error
	//fmt.Println("tempPath+configFileName", tempPath+configFileName)
	//err = configor.Load(&utils.K8sConfig, tempPath+configFileName)
	//config := utils.K8sConfig
	//fmt.Printf("%#v", config.K8s.Url)
	//fmt.Println("Workspace.Name", config.K8s.Workspace.Name)
	//fmt.Println("Workspace.Desc", config.K8s.Workspace.Desc)
	//fmt.Println("Namespace.Name", config.K8s.Namespace.Name)
	//fmt.Println("Username.Name", config.K8s.Username)
	//fmt.Println("Password.Name", config.K8s.Password)
	//fmt.Println("Repo.Name", config.K8s.Repo.Name)
	//fmt.Println("Repo.Projectname", config.K8s.Repo.Projectname)
	//fmt.Println("Appname", config.K8s.Appname)
	//fmt.Println("Mysql", config.K8s.Mysql.Database)
	//fmt.Println("Password", config.K8s.Mysql.Password)
	//fmt.Println("Harbor.Address", config.Harbor.Address)
	//fmt.Println("Project", config.Harbor.Project)
	//fmt.Println("Project", config.Harbor.Project)
	//fmt.Println("Password", config.Harbor.Password)
	//fmt.Println("Admin", config.Harbor.Admin)
	//return
	tempPath = "./kubetools/"
	tempVersion := tempPath + "version/"
	fmt.Println("临时目录路径", tempVersion)
	err = os.MkdirAll(tempVersion, os.ModePerm)
	if err != nil {
		fmt.Println("创建临时目录失败", err)
		return
	}
	//创建k8s中应用
	if configFileName == "" {
		configFileName = constant.KubeConfigName
	}
	err = configor.Load(&utils.K8sConfig, configFileName)
	if err != nil {
		log.WithField("err", err).Error("读取配置文件出错")
		return
	}

	if create == "config" {
		//导入配置模板
		err = k8s.ExportKubeConfig(configFileName)
		if err != nil {
			log.WithField("err", err).Error("创建配置文件失败")
			return
		}
	} else if create == "app" {

		config := utils.K8sConfig

		//导出文件
		k8s.Export(tempPath, config.K8s.Namespace.Name, config.K8s.Mysql.Database, config.K8s.Mysql.Password, config.K8s.Mysql.Image)

		//创建k8s中的配置项
		c := &k8s.Create{
			config.K8s.Workspace.Name,
			config.K8s.Namespace.Name,
			config.K8s.Repo.Name,
			config.K8s.Repo.Projectname,
			config.K8s.Workspace.Aliasname,
			config.K8s.Workspace.Desc,
			config.K8s.Namespace.Aliasname,
			config.K8s.Namespace.Desc,
			config.K8s.Appname,
			config.Harbor.Address,
		}
		csmp.KubeApply(tempPath, c)
	}

	fmt.Println("create=", create)
	fmt.Println("fileName=", configFileName)
	//return

	if initOs {
		//初始化操作系统
		InitOS()
	}
	fmt.Println("zipName", zipName)
	if zipName != "" {
		//推送升级包
		PushTar(tempVersion)
	}

	fmt.Println("finish")

}

// PushTar
// @description: 推送镜像和helm至私有仓库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 18:28
// @success:
func PushTar(tempVersion string) {
	//登录harbor仓库和添加helm仓库
	service.Init()
	utils.RunCommand("mkdir", "-p", "/root/.local/share/helm/plugins/helm-push")
	utils.RunCommand("tar", "-zxvf", "helm-push_0.10.3_linux_amd64.tar.gz", "-C", "/root/.local/share/helm/plugins/helm-push")
	//解压缩升级包
	utils.UnzipDir(zipName, tempVersion)
	zipPath, zipFileName := path.Split(zipName)

	fmt.Println("zipPath:", zipPath, "zipName", zipFileName)
	//解压后的路径
	err := harborService.DealFile(projectName, tempVersion+utils.UnExt(zipFileName))
	if err != nil {
		fmt.Println("=======err", err)
		return
	}
	err = utils.RunCommand("rm", "-rf", tempVersion+utils.UnExt(zipFileName))
	if err != nil {
		fmt.Println("清楚升级包文件失败", err)
		return
	}
}

// InitOS
// @description: 初始化操作系统
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 18:29
// @success:
func InitOS() {
	err := utils.RunCommand("chmod", "+x", tempPath+constant.InitOSFileName)
	if err != nil {
		fmt.Println("initOS.sh授权失败，请检查文件", err)
		return
	}
	err = utils.RunCommand(tempPath + constant.InitOSFileName)
	if err != nil {
		fmt.Println("配置操作系统初始化失败", err)
		return
	}
}
