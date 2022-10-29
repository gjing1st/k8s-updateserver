// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/13$ 17:01$

package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/utils"
)

var c *cron.Cron

// AddSecondFunc
// @description: 添加分钟定时任务
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/13 17:11
// @success:
func AddSecondFunc(s int, cmd func()) {

	spec := fmt.Sprintf("@every %ds", s)
	_, err := GetCron().AddFunc(spec, cmd)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("定时任务添加失败", err)).Error(constant.Msg)
	}
}

func InitCron() {
	c = cron.New(cron.WithSeconds())
}

// GetCron
// @description: 获取定时任务
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/13 17:06
// @success:
func GetCron() *cron.Cron {
	if c == nil {
		InitCron()
	}
	c.Start()
	return c
}
