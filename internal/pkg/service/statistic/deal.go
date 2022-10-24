// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/17$ 16:16$

package statistic

import (
	"encoding/json"
	"fmt"
	"statistic/internal/pkg/constant"
	"statistic/internal/pkg/model/statistic"
	"statistic/internal/pkg/model/statistic/response"
	"statistic/internal/pkg/utils"
	"time"
)

// RealTimeLog
// @description: 实时调度
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/16 15:04
// @success:
func (ss *StatisticService) RealTimeLog(msg, nameSpace, filter, cipherType string, eventTid, size, page int, startTime, endTime string) (res []response.Realtime, total int64, err error) {
	if msg == "" {
		msg = "mmyypt_app_events"
	}
	if nameSpace == "" {
		nameSpace = utils.Config.K8s.Namespace.Name
	}
	esRes, err := ss.RealTimeLogQuery(msg, nameSpace, filter, cipherType, eventTid, size, page, startTime, endTime)
	var realtimeRes statistic.RealTimeResponse
	json.Unmarshal(esRes, &realtimeRes)

	for _, v := range realtimeRes.Hits.Hits {
		cipherType := v.Source.EventDType
		if cipherType > constant.TypeCA {
			cipherType = 0
		}
		apiType := v.Source.EventType
		if apiType > constant.ApiTypeSm2SignVerify {
			apiType = 0
		}
		res = append(res, response.Realtime{
			EventAppid:   v.Source.EventAppid,
			EventAppName: v.Source.EventAppName,
			EventClient:  v.Source.EventClient,
			EventData:    v.Source.EventData,
			EventDType:   v.Source.EventDType,
			EventInfo:    v.Source.EventInfo,
			EventSerial:  v.Source.EventSerial,
			EventServer:  constant.CipherType[cipherType],
			EventTid:     v.Source.EventTid,
			EventTime:    v.Source.EventTime,
			EventType:    v.Source.EventType,
			EventApiName: constant.ApiType[apiType],
			EventResult:  true,
		})

	}
	total = int64(realtimeRes.Hits.Total)

	return
}

// CipherServerApi
// @description: 密码机服务次数查询
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/18 17:26
// @success:
func (ss StatisticService) CipherServerApi(timeRange, tenantId int, appid string, cipherType int, cipherSerial string) (list interface{}, err error) {
	var apiTotal [5][7]int
	var res response.CipherApiTotal
	var startTime, endTime time.Time
	//var res [constant.TimeLen]response.FlowAndTotal
	now := time.Now()
	switch timeRange {
	case constant.TimeOneHour, constant.TimeTwelveHour, constant.TimeOneDay:
		//一小时，12小时，1天。减掉小时
		for i := 0; i < constant.TimeLen; i++ {
			startTime = now.Add(time.Duration(constant.TimeRange[timeRange][i]))
			if i < constant.TimeLen-1 {
				endTime = now.Add(time.Duration(constant.TimeRange[timeRange][i+1]))
			} else {
				endTime = now
			}
			//去查询
			cipherApi, err1 := ss.SumCipherApi(startTime, endTime, tenantId, appid, cipherType, cipherSerial)

			if err1 != nil {
				err = err1
				return
			}
			fmt.Println("i=====", i, "======", cipherApi)
			FillCipherServer(&apiTotal, i, cipherApi)
			//放入结果集
			res.Time = append(res.Time, endTime)
			res.Total = apiTotal
		}

	case constant.TimeOneWeek, constant.TimeOneMonth, constant.TimeThreeMonth:
		//一周，一月，三个月。减掉天数
		for i := 0; i < constant.TimeLen; i++ {
			startTime = now.AddDate(0, 0, int(constant.TimeRange[timeRange][i]))
			if i < constant.TimeLen-1 {
				endTime = now.AddDate(0, 0, int(constant.TimeRange[timeRange][i+1]))
			} else {
				endTime = now
			}
			cipherApi, err1 := ss.SumCipherApi(startTime, endTime, tenantId, appid, cipherType, cipherSerial)
			if err1 != nil {
				err = err1
				return
			}
			FillCipherServer(&apiTotal, i, cipherApi)
			//放入结果集
			res.Time = append(res.Time, endTime)
			res.Total = apiTotal

		}
	case constant.TimeOneYear:
		//一年的。减掉月数
		for i := 0; i < constant.TimeLen; i++ {
			startTime = now.AddDate(0, int(constant.TimeRange[timeRange][i]), 0)
			if i < constant.TimeLen-1 {
				endTime = now.AddDate(0, int(constant.TimeRange[timeRange][i+1]), 0)
			} else {
				endTime = now
			}
			cipherApi, err1 := ss.SumCipherApi(startTime, endTime, tenantId, appid, cipherType, cipherSerial)
			if err1 != nil {
				err = err1
				return
			}
			for _, v := range cipherApi {
				if v.ApiType == constant.ApiTypeSm4Encrypt {

				}
			}
			FillCipherServer(&apiTotal, i, cipherApi)
			//放入结果集
			res.Time = append(res.Time, endTime)
			res.Total = apiTotal

		}
	}
	list = res
	return

}

// TimeStampApi
// @description: 交易概览-服务次数，时间戳
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/18 17:45
// @success:
func (ss StatisticService) TimeStampApi(timeRange, tenantId int, appid string, cipherType int, cipherSerial string) (list interface{}, err error) {
	var apiTotal [3][7]int
	var res response.CipherApiTotal
	var startTime, endTime time.Time
	//var res [constant.TimeLen]response.FlowAndTotal
	now := time.Now()
	switch timeRange {
	case constant.TimeOneHour, constant.TimeTwelveHour, constant.TimeOneDay:
		//一小时，12小时，1天。减掉小时
		for i := 0; i < constant.TimeLen; i++ {
			startTime = now.Add(time.Duration(constant.TimeRange[timeRange][i]))
			if i < constant.TimeLen-1 {
				endTime = now.Add(time.Duration(constant.TimeRange[timeRange][i+1]))
			} else {
				endTime = now
			}
			//去查询
			cipherApi, err1 := ss.SumCipherApi(startTime, endTime, tenantId, appid, cipherType, cipherSerial)

			if err1 != nil {
				err = err1
				return
			}
			FillTimeStampApi(&apiTotal, i, cipherApi)
			//放入结果集
			res.Time = append(res.Time, endTime)
		}

	case constant.TimeOneWeek, constant.TimeOneMonth, constant.TimeThreeMonth:
		//一周，一月，三个月。减掉天数
		for i := 0; i < constant.TimeLen; i++ {
			startTime = now.AddDate(0, 0, int(constant.TimeRange[timeRange][i]))
			if i < constant.TimeLen-1 {
				endTime = now.AddDate(0, 0, int(constant.TimeRange[timeRange][i+1]))
			} else {
				endTime = now
			}
			cipherApi, err1 := ss.SumCipherApi(startTime, endTime, tenantId, appid, cipherType, cipherSerial)
			if err1 != nil {
				err = err1
				return
			}
			FillTimeStampApi(&apiTotal, i, cipherApi)
			//放入结果集
			res.Time = append(res.Time, endTime)
		}
	case constant.TimeOneYear:
		//一年的。减掉月数
		for i := 0; i < constant.TimeLen; i++ {
			startTime = now.AddDate(0, int(constant.TimeRange[timeRange][i]), 0)
			if i < constant.TimeLen-1 {
				endTime = now.AddDate(0, int(constant.TimeRange[timeRange][i+1]), 0)
			} else {
				endTime = now
			}
			cipherApi, err1 := ss.SumCipherApi(startTime, endTime, tenantId, appid, cipherType, cipherSerial)
			if err1 != nil {
				err = err1
				return
			}
			for _, v := range cipherApi {
				if v.ApiType == constant.ApiTypeSm4Encrypt {

				}
			}
			FillTimeStampApi(&apiTotal, i, cipherApi)
			//放入结果集
			res.Time = append(res.Time, endTime)

		}
	}
	res.Total = apiTotal
	list = res
	return

}

// FillCipherServer
// @description: 交易概览-服务次数，密码机服务填充Y轴数据
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/18 17:32
// @success:
func FillCipherServer(apiTotal *[5][7]int, i int, cipherApi []statistic.CipherApi) {
	for _, v := range cipherApi {
		switch v.ApiType {
		case constant.ApiTypeSm2Sign:
			apiTotal[0][i] = v.Total
		case constant.ApiTypeSm4Encrypt:
			apiTotal[1][i] = v.Total
		case constant.ApiTypeSm4Decrypt:
			apiTotal[2][i] = v.Total
		case constant.ApiTypeSm3HMAC:
			apiTotal[3][i] = v.Total
		case constant.ApiTypeSm3HMACVerify:
			apiTotal[4][i] = v.Total

		}
	}
	fmt.Println("########", apiTotal)
}

// FillTimeStampApi
// @description: 交易概览-服务次数，时间戳填充Y轴数据
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/18 17:40
// @success:
func FillTimeStampApi(apiTotal *[3][7]int, index int, cipherApi []statistic.CipherApi) {
	for _, v := range cipherApi {
		switch v.ApiType {
		case constant.ApiTypeTimeStampCreate:
			apiTotal[0][index] = v.Total
		case constant.ApiTypeTimeStampVerify:
			apiTotal[1][index] = v.Total
		case constant.ApiTypeTimeStampParse:
			apiTotal[2][index] = v.Total
		}
	}
}
