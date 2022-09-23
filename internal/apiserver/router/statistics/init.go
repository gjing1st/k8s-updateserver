package statistics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	initStatisticsApi(apiV1)

	//启动gin路由服务
	err := router.Run(fmt.Sprintf(":%s", utils.Config.Web.Port))
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic("http服务启动失败")
	}
}
