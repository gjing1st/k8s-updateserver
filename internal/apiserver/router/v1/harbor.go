package v1

import (
	"github.com/gin-gonic/gin"
	"upserver/internal/apiserver/controller"
)

func initHarborApi(apiV1 *gin.RouterGroup)  {
	api :=apiV1.Group("/repo")
	harborController :=controller.HarborController{}
	//镜像仓库列表
	api.GET("/list",harborController.ListRepositories)
	api.GET("/list-art",harborController.ListArtifacts)
}
