// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/7$ 18:13$

package statistics

import (
	"github.com/gin-gonic/gin"
	"upserver/internal/apiserver/controller"
)

func initStatisticsApi(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/statistics")
	k8sController := controller.K8sVersionController{}
	api.POST("/broadcast", k8sController.UpdateVersion)

}
