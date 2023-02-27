package v1

import (
	"github.com/gin-gonic/gin"
	"upserver/internal/apiserver/controller"
)

// @description: k8s升级相关接口
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/28 9:31
// @success:
func initK8sVersionApi(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/versions")
	k8sController := controller.K8sVersionController{}
	api.GET("/latest", k8sController.GetVersion)
	api.GET("/list", k8sController.GetVersionList)
	api.POST("/update", k8sController.UpdateVersion)
	api.GET("/info", k8sController.VersionInfo)

}
