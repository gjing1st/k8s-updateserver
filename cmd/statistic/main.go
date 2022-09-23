// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 17:05$

package main

import (
	"upserver/internal/apiserver/router/statistics"
	"upserver/internal/pkg/service/statistic"
	"upserver/internal/pkg/utils"
)

func main() {
	utils.InitConfig()
	statistic.AddCron()
	statistics.InitApi()
	select {}
}
