// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/15$ 18:10$

package statistics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"upserver/internal/pkg/model/statistic/request"
	"upserver/internal/pkg/service/statistic"
)

type StatisticController struct{}

var statisticService statistic.StatisticService

// AppFlow
// @description: 业务流量
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/15 18:11
// @success:
func (sc *StatisticController) AppFlow(c *gin.Context) {
	var req request.AppFlow
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("===", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if req.Duration > 0 {
		//前端传了持续时间，按此参数查询
		req.EndTime = time.Now()
		req.StartTime = req.EndTime.Add(time.Hour * time.Duration(req.Duration) * -1)
	}
	res, err := statisticService.AppFlow(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "查询失败")
	} else {
		c.JSON(http.StatusOK, res)
	}

}

// CipherStatistic
// @description: 密码服务使用统计
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/16 10:08
// @success:
func (sc *StatisticController) CipherStatistic(c *gin.Context) {
	var req request.CipherStatistic
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if req.Duration > 0 {
		//前端传了持续时间，按此参数查询
		req.EndTime = time.Now()
		req.StartTime = req.EndTime.Add(time.Hour * time.Duration(req.Duration) * -1)
	}
	res, err := statisticService.CipherStatistic(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "查询失败")
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// RankingByApp
// @description: 业务调用排名
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/16 14:34
// @success:
func (sc *StatisticController) RankingByApp(c *gin.Context) {
	var req request.RankingByApp
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if req.Duration > 0 {
		//前端传了持续时间，按此参数查询
		req.EndTime = time.Now()
		req.StartTime = req.EndTime.Add(time.Hour * time.Duration(req.Duration) * -1)
	}
	res, err := statisticService.RankingByApp(req.StartTime, req.EndTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "查询失败")
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// RealTime
// @description: 实时调度
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/16 15:02
// @success:
func (sc *StatisticController) RealTime(c *gin.Context) {
	var req request.Realtime
	if err := c.ShouldBindJSON(&req); err != nil {
		//c.JSON(http.StatusBadRequest, err)
		//return
	}
	//return
	if req.Duration > 0 {
		//前端传了持续时间，按此参数查询
		req.EndTime = time.Now()
		req.StartTime = req.EndTime.Add(time.Hour * time.Duration(req.Duration) * -1)
	}
	res, err := statisticService.Realtime(req.Msg, req.NameSpace, req.Tid, 50)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "查询失败")
	} else {
		c.JSON(http.StatusOK, res)
	}
}
