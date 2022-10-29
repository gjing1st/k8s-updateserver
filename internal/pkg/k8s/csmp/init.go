// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/8/31$ 17:59$

package csmp

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/k8s"
	"upserver/internal/pkg/utils"
)

// KubeApply
// @description: ks安装后部署项目
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 18:44
// @success:
func KubeApply(path string, c *k8s.Create) {
	{
		//创建k8s应用
		k8s.CreateAll(c)
	}

	{
		//mysql相关配置应用到k8s中
		ApplyMysqlConf(path)
		ApplyMysqlSecret(path)
		ApplyMysql(path)
	}
	{
		//存储卷添加到k8s
		ApplyConfPvc(path)
		ApplyLibPvc(path)
		ApplyFrontendPvc(path)
		ApplyKmcPvc(path)
		ApplyMysqlPvc(path)
		ApplyUpdatePvc(path)
	}
}

// ApplyMysqlConf
// @description: 创建mysql配置字典
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/31 18:51
// @success:
func ApplyMysqlConf(path string) {
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.MysqlConfName)
	if err != nil {
		fmt.Println("创建mysql配置字典失败", err)
		return
	}
}

// ApplyMysqlSecret
// @description: 创建数据库保密字典
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 17:30
// @success:
func ApplyMysqlSecret(path string) {
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.MysqlSecretName)
	if err != nil {
		fmt.Println("创建mysql保密字典失败", err)
		return
	}
}

// ApplyMysql
// @description: 创建数据库配置字典
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 17:30
// @success:
func ApplyMysql(path string) {
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.MysqlName)
	if err != nil {
		fmt.Println("创建mysql工作负载失败", err)
		return
	}
}

// ApplyConfPvc
// @description: 将conf存储卷添加到k8s
// @param: path 路径
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 18:06
// @success:
func ApplyConfPvc(path string) {
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.ConfPvc)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("pvc 添加到k8s失败")
		return
	}
}

func ApplyLibPvc(path string) {
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.LibPvc)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("pvc 添加到k8s失败")
		return
	}
}

func ApplyFrontendPvc(path string) {
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.FrontendPvc)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("pvc 添加到k8s失败")
		return
	}
}

func ApplyKmcPvc(path string) {
	cmd := "/usr/local/bin/kubectl apply -f " + path + constant.KmcPvc
	err := utils.RunCommand("/usr/local/bin/kubectl", "apply", "-f", path+constant.KmcPvc)
	if err != nil {
		log.WithFields(log.Fields{"cmd": cmd, "err": err}).Error("pvc 添加到k8s失败")
		return
	}
}

func ApplyMysqlPvc(path string) {
	cmd := "/usr/local/bin/kubectl apply -f " + path + constant.MysqlPvc
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.MysqlPvc)
	if err != nil {
		log.WithFields(log.Fields{"cmd": cmd, "err": err}).Error("pvc 添加到k8s失败")
		return
	}
}

func ApplyUpdatePvc(path string) {
	cmd := "/usr/local/bin/kubectl apply -f " + path + constant.UpdatePvc
	err := utils.RunCommand("kubectl", "apply", "-f", path+constant.UpdatePvc)
	if err != nil {
		log.WithFields(log.Fields{"cmd": cmd, "err": err}).Error("pvc 添加到k8s失败")
		return
	}
}
