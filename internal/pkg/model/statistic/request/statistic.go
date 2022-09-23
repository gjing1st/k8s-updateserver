// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/15$ 18:13$

package request

import "time"

type BaseRequest struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  int       `json:"duration"` //小时，需要前端换算好
	Msg       string    `json:"msg"`
	NameSpace string    `json:"name_space"`
	Tid       int       `json:"tid"` //租户id
}

//AppFlow 业务流量
type AppFlow struct {
	BaseRequest
}

//CipherStatistic 密码服务使用
type CipherStatistic struct {
	BaseRequest
}

//RankingByApp 业务排名
type RankingByApp struct {
	BaseRequest
}

type Realtime struct {
	BaseRequest
}
