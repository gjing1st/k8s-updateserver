package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/service"
	"upserver/internal/pkg/utils"
)

type UploadController struct {
	BaseController
}

var dockerService service.DockerService

// UploadTar
// @description: 上传镜像tar包
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/19 18:15
// @success:
func (uc UploadController) UploadTar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, constant.RequestParamErr)
		return
	}
	fileExt := path.Ext(file.Filename)
	if fileExt != "tar" {
		c.JSON(http.StatusBadRequest, constant.RequestErrExt)
		return
	}
	//日期存放路径
	dirName := "/" + time.Now().Format("2006_01_02") + "/"

	//创建存放目录
	os.MkdirAll(utils.Config.Path+dirName, 0777)
	//完整路径文件名
	fullName := utils.Config.Path + dirName + file.Filename
	//存放文件
	if err := c.SaveUploadedFile(file, fullName); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	files := strings.Split(file.Filename, "_")
	//镜像名称
	fileName := files[0]
	//镜像版本号
	fileVersion := "latest"
	if len(files) > 1 {
		versions := strings.Split(files[1],".")
		fileVersion = versions[0]
	}
	//镜像名称
	imageName := fileName + ":" + fileVersion
	//标记的镜像名称
	tagName := utils.K8sConfig.Harbor.Address+ "/" + "library" + "/" + imageName
	err = dockerService.PushHarbor(fullName, imageName, tagName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

}

func (uc UploadController) Test(c *gin.Context)  {
	//完整路径文件名
	Filename := "mckserver_latest.tar"
	fullName := "/home/app/mckserver_latest.tar"
	files := strings.Split(Filename, "_")
	//镜像名称
	fileName := files[0]
	//镜像版本号
	fileVersion := "latest"
	if len(files) > 1 {
		versions := strings.Split(files[1],".")
		fileVersion = versions[0]
	}
	//镜像名称
	imageName := fileName + ":" + fileVersion
	//标记的镜像名称
	tagName := utils.K8sConfig.Harbor.Address+ "/" + "library" + "/" + imageName
	err := dockerService.PushHarbor(fullName, imageName, tagName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
}
