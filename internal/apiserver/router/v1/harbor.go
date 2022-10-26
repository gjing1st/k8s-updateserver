package v1

import (
	"upserver/internal/apiserver/controller"

	"github.com/gin-gonic/gin"
)

func initHarborApi(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/repo")
	harborController := controller.HarborController{}
	//镜像仓库列表
	api.GET("/list", harborController.ListRepositories)
	api.GET("/list-art", harborController.ListArtifacts)
	api.POST("/upload", harborController.Upload)
	api.POST("/upload/info/sum", harborController.UploadInfoSummary)
	api.POST("/upload/info/:version", harborController.UploadInfo)
}
