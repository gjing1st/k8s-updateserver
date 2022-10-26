// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/14$ 10:23$

package statistic

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/model/statistic"
	"upserver/internal/pkg/model/statistic/response"
	"upserver/internal/pkg/utils"
	"upserver/internal/pkg/utils/database"
)

// RankingByApp
// @description:
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/14 14:29
// @success:
func (ss *StatisticService) RankingByApp(startTime, endTime time.Time) (res []response.AppStatistic, err error) {
	cli := database.GetMgoCli()
	mgo := cli.Database(utils.K8sConfig.K8s.Statistic.MongoDatabase)
	collection := mgo.Collection(utils.K8sConfig.K8s.Statistic.Collection)

	//groups := mongo.Pipeline{bson.D{
	//	{"$match", bson.D{
	//		{"start_time", bson.D{
	//			{"$gte", startTime},
	//			{"$lt", endTime},
	//		}},
	//	}},
	//	{"$group", bson.D{
	//		{"_id", "$appid"},
	//		{"total", bson.D{
	//			{"$sum", "$total"},
	//		}},
	//	}},
	//}}

	matchTimes := statistic.MatchTime{
		statistic.StartTime{
			statistic.StartTimeRange{
				Gte: startTime.Unix(),
				Lt:  endTime.Unix(),
			},
		},
	}
	//要分组的字段
	var groupField []string
	groupField = append(groupField, "$appname")
	group := statistic.Group{
		statistic.GroupField{
			Field: groupField,
			Total: statistic.GroupTotal{
				"$total",
			},
		},
	}
	groups := []interface{}{}
	groups = append(groups, matchTimes)
	groups = append(groups, group)
	//delCond := DeleteCond{BeforeCond: TimeBeforeCond{BeforeTime: time.Now().Unix()}}
	//查询
	//s1, _ := json.Marshal(groups)
	//fmt.Println("-----------", string(s1))
	cursor, err := collection.Aggregate(context.TODO(), groups)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("查询mongodb失败", err)).Error(constant.Msg)
		return
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("len(results)========", len(results))
	for _, result := range results {
		fmt.Println(result)
	}
	//对结果进行解析
	resultsJson, _ := json.Marshal(results)
	var md []statistic.MarshalMgoData
	json.Unmarshal(resultsJson, &md)
	for _, v := range md {
		////类型断言
		//appName := ""
		//switch v.Id[0].(type) {
		//case string:
		//	appName = v.Id[0].(string)
		//default:
		//	appName = ""
		//}
		res = append(res, response.AppStatistic{
			utils.String(v.Id[0]),
			utils.Int(v.Total),
		})
	}
	return
}

// CipherStatistic
// @description: 密码服务统计
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/15 17:05
// @success:
func (ss *StatisticService) CipherStatistic(startTime, endTime time.Time) (res []response.CipherStatistic, err error) {
	cli := database.GetMgoCli()
	mgo := cli.Database(utils.K8sConfig.K8s.Statistic.MongoDatabase)
	collection := mgo.Collection(utils.K8sConfig.K8s.Statistic.Collection)
	//时间范围
	matchTimes := statistic.MatchTime{
		statistic.StartTime{
			statistic.StartTimeRange{
				Gte: startTime.Unix(),
				Lt:  endTime.Unix(),
			},
		},
	}
	//要分组的字段
	var groupField []string
	groupField = append(groupField, "$ciphertype", "$cipherserial")
	group := statistic.Group{
		statistic.GroupField{
			Field: groupField,
			Total: statistic.GroupTotal{
				"$total",
			},
		},
	}
	groups := []interface{}{}
	groups = append(groups, matchTimes)
	groups = append(groups, group)
	//delCond := DeleteCond{BeforeCond: TimeBeforeCond{BeforeTime: time.Now().Unix()}}
	//查询
	cursor, err := collection.Aggregate(context.TODO(), groups)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("查询mongodb失败", err)).Error(constant.Msg)
		return
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
	//对结果进行解析
	resultsJson, _ := json.Marshal(results)
	var md []statistic.MarshalMgoData
	json.Unmarshal(resultsJson, &md)
	for _, v := range md {

		//类型断言
		//var cipherType int
		//switch v.Id[0].(type) {
		//case float64:
		//	cipherType = int(v.Id[0].(float64))
		//case int:
		//	cipherType = v.Id[0].(int)
		//case int64:
		//	cipherType = int(v.Id[0].(int64))
		//case int32:
		//	cipherType = int(v.Id[0].(int32))
		//case string:
		//	cipherType, _ = strconv.Atoi(v.Id[0].(string))
		//default:
		//	cipherType = utils.Int(v.Id[0])
		//}
		cipherType := utils.Int(v.Id[0])
		//密码服务类型
		if cipherType > constant.TypeCA {
			//防止数组index越界
			cipherType = 0
		}
		typeStr := constant.CipherType[cipherType]

		res = append(res, response.CipherStatistic{
			typeStr + utils.String(v.Id[1]),
			utils.Int(v.Total),
		})
	}
	//fmt.Println("=======res", res)

	return
}

// AppFlow
// @description: 业务流量
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/15 18:06
// @success:
func (ss *StatisticService) AppFlow(startTime, endTime time.Time) (res []response.AppFlow, err error) {
	cli := database.GetMgoCli()
	mgo := cli.Database(utils.K8sConfig.K8s.Statistic.MongoDatabase)
	collection := mgo.Collection(utils.K8sConfig.K8s.Statistic.Collection)
	//查询
	//var matchs statistic.Match
	//var matchsIfc []interface{}
	//时间范围
	matchTimes := statistic.MatchTime{
		statistic.StartTime{
			statistic.StartTimeRange{
				Gte: startTime.Unix(),
				Lt:  endTime.Unix(),
			},
		},
	}
	//要分组的字段
	var groupField []string
	groupField = append(groupField, "$appname")
	group := statistic.Group{
		statistic.GroupField{
			Field: groupField,
			Total: statistic.GroupTotal{
				"$flow",
			},
		},
	}
	groups := []interface{}{}
	groups = append(groups, matchTimes)
	groups = append(groups, group)
	//delCond := DeleteCond{BeforeCond: TimeBeforeCond{BeforeTime: time.Now().Unix()}}
	//查询
	cursor, err := collection.Aggregate(context.TODO(), groups)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("查询mongodb失败", err)).Error(constant.Msg)
		return
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	//对结果进行解析
	resultsJson, _ := json.Marshal(results)
	var md []statistic.MarshalMgoData
	json.Unmarshal(resultsJson, &md)
	for _, v := range md {
		res = append(res, response.AppFlow{
			utils.String(v.Id[0]),
			utils.Int(v.Total),
		})
	}
	fmt.Println("=======res", res)

	return
}

// Realtime
// @description: 实时调度
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/16 15:04
// @success:
func (ss *StatisticService) Realtime(msg, nameSpace string, eventTid, size int) (res []response.Realtime, err error) {
	if msg == "" {
		msg = "mmyypt_app_events"
	}
	if nameSpace == "" {
		nameSpace = utils.K8sConfig.K8s.Namespace.Name
	}
	esRes, err := ss.RealTimeQuery(msg, nameSpace, eventTid)
	var realtimeRes statistic.RealTimeResponse
	json.Unmarshal(esRes, &realtimeRes)

	for _, v := range realtimeRes.Hits.Hits {
		cipherType := v.Source.EventDType
		if cipherType > constant.TypeCA {
			cipherType = 0
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
		})

	}

	return
}

// LastData
// @description: 查询最新一条数据
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/8 17:15
// @success:
func (ss StatisticService) LastData() (s *statistic.StatisticsTable) {
	collection := database.GetCollection()
	s = new(statistic.StatisticsTable)
	err := collection.FindOne(context.TODO(), bson.D{}, options.FindOne().SetSort(bson.D{{"_id", -1}})).Decode(s)
	if err != nil {
		log.WithField("mongo serach err", nil).Error()
		return
	}

	//for cuscor.Next(context.TODO()) {
	//	err = cuscor.Decode(s)
	//	if err != nil {
	//		//log.WithField("mongo error", err).Error()
	//	}
	//}
	return
}

// SumFlowAndTotal
// @description: 获取流量调用次数
// @param: tenantId int 租户id
// @param: appName string 业务名称
// @param: cipherType int 密码资源类型
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/9 18:24
// @success:
func (ss StatisticService) SumFlowAndTotal(startTime, endTime time.Time, tenantId int, appid string, cipherType int) (res statistic.FlowAndTotal) {
	collection := database.GetCollection()
	//查询条件
	filter := bson.D{}
	//时间范围
	var eventTime bson.E
	eventTime = bson.E{"event_time", bson.D{{"$gte", startTime.Unix()}, {"$lt", endTime.Unix()}}}
	filter = append(filter, eventTime)
	if tenantId > 0 {
		filter = append(filter, bson.E{"tenant_id", tenantId})
	}
	if appid != "" {
		filter = append(filter, bson.E{"appid", appid})
	}
	if cipherType > 0 {
		filter = append(filter, bson.E{"cipher_type", cipherType})
	}
	//分组
	//group := bson.E{"total", bson.E{"$sum", "$flow"}}
	//filter = append(filter, group)

	fmt.Println("filter", filter)
	pipeline := bson.A{
		bson.D{{"$match", filter}},
		bson.D{
			{"$group", bson.D{
				{"_id", "null"},
				{"flow", bson.D{
					{"$sum", "$flow"},
				}},
				{"total", bson.D{
					{"$sum", "$total"},
				}},
			}},
		},
		//bson.D{{"$sort", bson.D{{"_id", 1}}}},
	}
	//groupStage := bson.D{
	//	{"$group", bson.D{
	//		{"_id", "null"},
	//		{"count", bson.D{
	//			{"$sum", "$flow"},
	//		}},
	//	}},
	//}
	//sss := mongo.Pipeline{pipeline}
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("查询mongodb失败", err)).Error(constant.Msg)
		return
	}
	var resArr []statistic.FlowAndTotal
	if err = cursor.All(context.Background(), &resArr); err != nil {
		log.Fatal(err)
	}
	fmt.Println("res", res)
	if len(resArr) == 0 {
		return
	} else {
		return resArr[0]
	}
}

// GetFlowAndTotal
// @description: 获取流量和调用次数
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/10 18:47
// @success:
func (ss StatisticService) GetFlowAndTotal(timeRange, tenantId int, appid string, cipherType int) (res response.FlowAndTotal) {
	var startTime, endTime time.Time
	//var res [constant.TimeLen]response.FlowAndTotal
	now := time.Now()
	//if value, ok := constant.TimeRange[timeRange]; ok {
	//	for i := 0; i < constant.TimeLen; i++ {
	//		startTime = endTime.Add(time.Duration(constant.TimeRange[timeRange][0]))
	//		flow := ss.SumFlow(startTime, endTime, tenantId, appid, cipherType)
	//	}
	//
	//}
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
			flowTotal := ss.SumFlowAndTotal(startTime, endTime, tenantId, appid, cipherType)
			//放入结果集
			res.Time = append(res.Time, endTime)
			res.Flow = append(res.Flow, flowTotal.Flow)
			res.Total = append(res.Total, flowTotal.Total)
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
			flowTotal := ss.SumFlowAndTotal(startTime, endTime, tenantId, appid, cipherType)
			//放入结果集
			res.Time = append(res.Time, endTime)
			res.Flow = append(res.Flow, flowTotal.Flow)
			res.Total = append(res.Total, flowTotal.Total)

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
			flowTotal := ss.SumFlowAndTotal(startTime, endTime, tenantId, appid, cipherType)
			//放入结果集
			res.Time = append(res.Time, endTime)
			res.Flow = append(res.Flow, flowTotal.Flow)
			res.Total = append(res.Total, flowTotal.Total)

		}
	}

	//fmt.Println(res)
	return res
}
