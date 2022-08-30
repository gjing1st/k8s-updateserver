package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"upserver/internal/apiserver/controller"
	"upserver/internal/pkg/middleware"
	"upserver/internal/pkg/utils"
)

func InitApi() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//是否跨域
	if utils.Config.Web.Cors {
		router.Use(middleware.CORS)
	}

	apiV1 := router.Group("/v1")
	//ping服务检测接口
	apiV1.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})
	//上传接口
	var uploadController controller.UploadController
	//镜像上传
	apiV1.POST("/tar/upload",uploadController.UploadTar)//改到harbor，上传整体压缩包
	apiV1.GET("test",uploadController.Test)
	//调用harbor相关接口
	initHarborApi(apiV1)
	initK8sVersionApi(apiV1)

	//启动gin路由服务
	err := router.Run(fmt.Sprintf(":%s", utils.Config.Web.Port))
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic("http服务启动失败")
	}
}
