package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/harbor"
	"upserver/internal/pkg/k8s"
	"upserver/internal/pkg/service"
	"upserver/internal/pkg/utils"
)

var harborService service.HarborService

type HarborController struct {
	BaseController
}

// ListRepositories
// @description: 镜像仓库列表
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/25 18:05
// @success:
func (hc HarborController) ListRepositories(c *gin.Context) {

	res, err := harbor.ListRepositories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "获取失败")
		return
	}
	c.JSON(http.StatusOK, res)

}

func (hc HarborController) ListArtifacts(c *gin.Context) {
	var params harbor.ArtList
	if err := c.ShouldBindJSON(&params); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Warn(constant.RequestParamErr)
		c.JSON(http.StatusBadRequest, constant.RequestParamErr)
		return
	}
	res, err := harbor.ListArtifacts(params.RepositoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "获取失败")
		return
	}
	c.JSON(http.StatusOK, res)

}

// Upload
// @description: 上传升级包接口
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/28 14:04
// @success:
func (hc HarborController) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, constant.RequestParamErr)
		return
	}
	//文件名为harbor项目名_版本号格式 ex:csmp_V3.1.0.zip
	fileExt := path.Ext(file.Filename)
	if fileExt != ".zip" {
		log.WithField("fileExt", fileExt).Error(constant.RequestErrExt)
		c.JSON(http.StatusGone, constant.RequestErrExt)
		return
	}
	//日期存放路径
	dirName := "/" + time.Now().Format("2006-01-02_15_04") + "/"

	//创建存放目录
	os.MkdirAll(utils.Config.Path+dirName, 0777)
	//完整路径文件名
	fullName := utils.Config.Path + dirName + file.Filename
	//存放文件
	if err := c.SaveUploadedFile(file, fullName); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	constant.HarborPushed = 0
	c.JSON(http.StatusOK, nil)
	//上传成功后先返回成功，再解压缩和推送私有仓库
	//return
	//解压缩升级包
	utils.UnzipDir(fullName, utils.Config.Path+dirName)
	files := strings.Split(file.Filename, "_")
	//要上传到的harbor项目名称
	projectName := files[0]
	//解压后的路径
	dirPath := utils.Config.Path + dirName + utils.UnExt(file.Filename)
	//解压后处理解压后的文件
	err = harborService.DealFile(projectName, dirPath)
	if err == nil {
		constant.HarborPushed = 1
		//c.JSON(http.StatusInternalServerError, err.Error())
		//return
	}
	//更新应用仓库
	_ = k8s.GetAndUpdateRepo("")

}
