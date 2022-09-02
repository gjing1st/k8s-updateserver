// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/8/31$ 15:24$

package main

import (
	"flag"
	"fmt"
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
	ConfigName  string
	projectName string
	tempPath    string
	initOs      bool
	namespace   string
	createApp   bool
	workspace string
	appName string
)

func init() {
	flag.StringVar(&zipName, "z", "", "zip 指定zip压缩包名称用于推送至harbor私有仓库")
	flag.StringVar(&ConfigName, "f", "", "configFileName 配置文件名称用来k8s部署初始化")
	flag.StringVar(&projectName, "p", "test1", "projectName harbor私有仓库项目名称，镜像要上传至的项目")
	flag.StringVar(&tempPath, "t", "/tmp/k8s/", "tmpPath 解压升级包的目录")
	flag.BoolVar(&initOs, "i", false, "initOS 初始化操作系统")
	flag.StringVar(&namespace, "n", "csmp", "命名空间，即企业空间中的项目")
	flag.BoolVar(&createApp, "c", false, "是否开始部署k8s中的实际项目应用")
	flag.StringVar(&workspace, "w", "dked", "企业空间")
	flag.StringVar(&appName, "a", "csmp", "实际要部署项目应用名称")

}

//go run main.go -z E:/project/csmp/version/csmp_v3.2.0.1.zip - i true
func main() {
	flag.Parse()
	tempVersion := tempPath + "version/"
	_ = os.MkdirAll(tempVersion, os.ModePerm)
	//导出文件
	k8s.Export(tempPath,projectName,"mmyypd_db","zf12345678")


	if initOs {
		//初始化操作系统
		InitOS()
	}
	if zipName != "" {
		//推送升级包
		PushTar(tempVersion)
	}
	if createApp{
		//创建k8s中的配置项
		c := &k8s.Create{
			workspace,
			namespace,
			projectName,
			"harbor",
			"library",
			"",
			"",
			"",
			"",
			appName,
		}
		csmp.KubeApply(tempPath,c)
		//k8s.CreateAll(c)
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
func PushTar(tempVersion string)  {
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
	err = utils.RunCommand("rm","-rf",tempVersion+utils.UnExt(zipFileName))
	if err != nil {
		fmt.Println("清楚升级包文件失败",err)
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
func InitOS()  {
	err := utils.RunCommand("chmod","+x",tempPath + constant.InitOSFileName)
	if err != nil {
		fmt.Println("initOS.sh授权失败，请检查文件",err)
		return
	}
	err = utils.RunCommand(tempPath + constant.InitOSFileName)
	if err != nil {
		fmt.Println("配置操作系统初始化失败",err)
		return
	}
}

