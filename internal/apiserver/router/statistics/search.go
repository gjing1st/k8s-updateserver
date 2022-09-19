// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/7$ 18:13$

package statistics

import (
	"github.com/gin-gonic/gin"
	"upserver/internal/apiserver/controller/statistics"
)

func initStatisticsApi(apiV1 *gin.RouterGroup) {
	api := apiV1.Group("/statistics")
	statisticController := statistics.StatisticController{}
	api.POST("/appFlow", statisticController.AppFlow)
	api.POST("/cipherStatistic", statisticController.CipherStatistic)
	api.POST("/rankingByApp", statisticController.RankingByApp)
	api.POST("/realtime", statisticController.RealTime)

}
