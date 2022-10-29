// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/8/31$ 16:33$

package k8s

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/k8s/tmpl"
	"upserver/internal/pkg/utils"
)

// Export
// @description: 导出文件
// @param: path string 要导出至的路径
// @param: projectName string k8s中的namespace对应ks企业空间中的项目
// @param: database string 要创建的数据库
// @param: password string 数据库root密码明文
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/1 17:28
// @success:
func Export(path, projectName, database, password, image string) {
	ExportInitOS(path)
	{
		//导出mysql数据库相关yaml文件
		ExportMysqlServer(path, projectName, image)
		ExportMysqlConf(path, projectName)
		ExportMysqlSecret(path, projectName, database, password)
	}
	{
		//导出存储卷yaml文件
		ExportConfPvc(path, projectName)
		ExportLibPvc(path, projectName)
		ExportFrontendPvc(path, projectName)
		ExportKmcPvc(path, projectName)
		ExportMysqlPvc(path, projectName)
		ExportUpdatePvc(path, projectName)
	}

}

//ExportInitOS
// @description: 导入系统初始化文件
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/31 16:39
// @success:
func ExportInitOS(tempPath string) {
	initOSFile := tempPath + constant.InitOSFileName
	file, err := os.OpenFile(initOSFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(tmpl.InitOsTmpl)
	write.Flush()
}

// ExportMysqlServer
// @description: mysql服务和工作负载
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/31 16:38
// @success:
func ExportMysqlServer(tempPath, projectName, image string) {
	fileName := tempPath + constant.MysqlName
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	mysqlYaml := strings.Replace(tmpl.MysqlServer, "{{projectName}}", projectName, 2)
	mysqlYaml = strings.Replace(mysqlYaml, "{{image}}", image, 1)
	write.WriteString(mysqlYaml)
	write.Flush()
}

// ExportMysqlConf
// @description: configMap配置字典
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/31 16:39
// @success:
func ExportMysqlConf(tempPath, projectName string) {
	fileName := tempPath + constant.MysqlConfName
	mysqlYaml := strings.Replace(tmpl.MysqlConf, "{{projectName}}", projectName, 2)
	utils.WriteFile(fileName, mysqlYaml)
}

// ExportMysqlSecret
// @description: mysql保密字典
// @param: path string 要导出至的路径
// @param: projectName string k8s中的namespace对应ks企业空间中的项目
// @param: database string 要创建的数据库明文
// @param: password string 数据库root密码明文
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/31 16:41
// @success:
func ExportMysqlSecret(tempPath, projectName, database, password string) {
	fileName := tempPath + constant.MysqlSecretName
	mysqlYaml := strings.Replace(tmpl.MysqlSecret, "{{projectName}}", projectName, 2)
	database = base64.StdEncoding.EncodeToString([]byte(database))
	mysqlYaml = strings.Replace(mysqlYaml, "{{database}}", database, 1)
	password = base64.StdEncoding.EncodeToString([]byte(password))
	mysqlYaml = strings.Replace(mysqlYaml, "{{password}}", password, 1)
	utils.WriteFile(fileName, mysqlYaml)
}

// ExportConfPvc
// @description: 导出conf 存储卷
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 17:53
// @success:
func ExportConfPvc(tempPath, projectName string) {
	fileName := tempPath + constant.ConfPvc
	pvcYaml := strings.Replace(tmpl.ConfPvc, "{{projectName}}", projectName, 2)
	utils.WriteFile(fileName, pvcYaml)
}

func ExportLibPvc(tempPath, projectName string) {
	fileName := tempPath + constant.LibPvc
	pvcYaml := strings.Replace(tmpl.LibPvc, "{{projectName}}", projectName, 2)
	utils.WriteFile(fileName, pvcYaml)
}

func ExportFrontendPvc(tempPath, projectName string) {
	fileName := tempPath + constant.FrontendPvc
	pvcYaml := strings.Replace(tmpl.FrontendPvc, "{{projectName}}", projectName, 2)
	utils.WriteFile(fileName, pvcYaml)
}

func ExportKmcPvc(tempPath, projectName string) {
	fileName := tempPath + constant.KmcPvc
	pvcYaml := strings.Replace(tmpl.KmcPvc, "{{projectName}}", projectName, 2)
	utils.WriteFile(fileName, pvcYaml)
}

func ExportMysqlPvc(tempPath, projectName string) {
	fileName := tempPath + constant.MysqlPvc
	pvcYaml := strings.Replace(tmpl.MysqlPvc, "{{projectName}}", projectName, 2)
	utils.WriteFile(fileName, pvcYaml)
}

func ExportUpdatePvc(tempPath, projectName string) {
	fileName := tempPath + constant.UpdatePvc
	pvcYaml := strings.Replace(tmpl.UpdatePvc, "{{projectName}}", projectName, 2)
	utils.WriteFile(fileName, pvcYaml)
}

// ExportKubeConfig
// @description: 导出kubeTool配置文件
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/3 21:12
// @success:
func ExportKubeConfig(fileName string) error {
	if fileName == "" {
		fileName = constant.KubeConfigName
	}
	filePath := "./" + fileName
	return utils.WriteFile(filePath, tmpl.ConfigFile)
}
