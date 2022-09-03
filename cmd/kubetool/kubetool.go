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
	zipName     string
	configFileName  string
	projectName string
	tempPath    string
	initOs      bool
	create string
)

func init() {
	flag.StringVar(&zipName, "z", "", "zip 指定zip压缩包名称用于推送至harbor私有仓库")
	flag.StringVar(&configFileName, "f", "", "configFileName 配置文件名称用来k8s部署初始化")
	flag.StringVar(&tempPath, "t", "./", "tmpPath 解压升级包的目录")
	flag.BoolVar(&initOs, "i", false, "initOS 初始化操作系统")
	flag.StringVar(&create, "c", "", "create config|app 创建k8s应用或配置文件")

}



//go run main.go -z E:/project/csmp/version/csmp_v3.2.0.1.zip
func main() {
	flag.Parse()
	var err error
	if configFileName == ""{
		//配置文件名称
		configFileName = constant.KubeConfigName
	}
	if create == "config"{
		//导入配置模板
		err = k8s.ExportKubeConfig(configFileName)
		if err != nil {
			log.WithField("err",err).Error("创建配置文件失败")
			return
		}
	}else if create == "app"{
		//创建k8s中应用
		err = configor.Load(&utils.KubeToolConfig, tempPath+configFileName)
		if err != nil {
			log.WithField("err",err).Error("读取配置文件出错")
			return
		}
		config := utils.KubeToolConfig

		//导出文件
		k8s.Export(tempPath, config.K8s.Namespace.Name, config.K8s.Mysql.Database, config.K8s.Mysql.Password)

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



	fmt.Println("create=",create)
	fmt.Println("fileName=",configFileName)
	fmt.Printf("%#v",utils.KubeToolConfig.K8s.Namespace.Desc)
	return


	tempVersion := tempPath + "version/"
	_ = os.MkdirAll(tempVersion, os.ModePerm)


	if initOs {
		//初始化操作系统
		InitOS()
	}
	if zipName != "" {
		//推送升级包
		PushTar(tempVersion)
	}

	fmt.Println("success")

}

// PushTar
// @description: 推送镜像和helm至私有仓库
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/2 18:28
// @success:
func PushTar(tempVersion string) {
	//登录harbor仓库和添加helm仓库
	service.Init()

	//解压缩升级包
	utils.UnzipDir(zipName, tempVersion)
	zipPath, zipFileName := path.Split(zipName)

	fmt.Println("zipPath:", zipPath, "zipName", zipFileName)
	//解压后的路径
	err := harborService.DealFile(projectName, tempVersion+utils.UnExt(zipFileName))
	if err != nil {
		fmt.Println("=======err", err)
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
// @email: guojing@tna.cn
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
