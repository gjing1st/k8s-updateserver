// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/15$ 18:13$

package request

import (
	"time"
)

type BaseRequest struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Duration  int       `json:"duration"` //小时，需要前端换算好
	Msg       string    `json:"msg"`
	NameSpace string    `json:"name_space"`
	TenantId  int       `json:"tenant_id"`  //租户id
	TimeRange int       `json:"time_range"` //时间范围 constant.TimeOneHour
}

//AppFlow 业务流量
type AppFlow struct {
	BaseRequest
	Appid      string `json:"appid"`       //业务id
	AppName    string `json:"app_name"`    //业务名称
	CipherType int    `json:"cipher_type"` //密码服务

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
type TransactionBase struct {
	Msg       string `json:"msg"`
	NameSpace string `json:"name_space"`
	TimeRange int    `json:"time_range"` //时间范围 constant.TimeOneHour
	TenantId  int    `json:"tenant_id"`  //租户id

}

//CipherServer 服务次数
type CipherServer struct {
	TransactionBase
	CipherSerial string `json:"cipher_serial" binding:"required"` //密码资源序列号
	CipherType   int    `json:"cipher_type" binding:"required"`   //密码服务

}

//RealTimeLog 交易日志
type RealTimeLog struct {
	TransactionBase
	CipherType string `json:"cipher_type"`                //密码服务
	StartTime  string `json:"start_time"`                 //开始时间
	EndTime    string `json:"end_time"`                   //结束时间
	Filter     string `json:"filter"`                     //搜索条件
	Page       int    `json:"page" form:"page"`           // 页码
	PageSize   int    `json:"page_size" form:"page_size"` // 每页大小
	Result     string `json:"result"`                     //结果
}

type CalculateAppFlow struct {
	TransactionBase
	CalculateType int `json:"calculate_type" binding:"required"` // 1调用次数 2使用流量
	CipherType    int `json:"cipher_type"`                       //密码服务
}

type Ranking struct {
	TransactionBase
	RankingType   int `json:"ranking_type" binding:"required"`   // 1业务排名 2租户排名
	CalculateType int `json:"calculate_type" binding:"required"` // 1调用次数 2使用流量

}

//CipherUsage 资源使用率
type CipherUsage struct {
	TransactionBase
	CipherSerial string `json:"cipher_serial"` //密码资源序列号
	CipherType   int    `json:"cipher_type"`   //密码服务
}
