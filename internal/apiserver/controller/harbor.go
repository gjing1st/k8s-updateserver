package controller

import (
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/harbor"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

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
func (hc HarborController) ListRepositories(c *gin.Context)  {

	res,err := harbor.ListRepositories()
	if err != nil{
		c.JSON(http.StatusInternalServerError,"获取失败")
		return
	}
	c.JSON(http.StatusOK,res)

}

func (hc HarborController) ListArtifacts(c *gin.Context)  {
	var params harbor.ArtList
	if err := c.ShouldBindJSON(&params);err != nil{
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Warn(constant.RequestParamErr)
		c.JSON(http.StatusBadRequest, constant.RequestParamErr)
		return
	}
	res,err := harbor.ListArtifacts(params.RepositoryName)
	if err != nil{
		c.JSON(http.StatusInternalServerError,"获取失败")
		return
	}
	c.JSON(http.StatusOK,res)

}