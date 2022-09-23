// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/6$ 16:42$

package statistic

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/model/statistic"
	"upserver/internal/pkg/utils"
	"upserver/internal/pkg/utils/crontab"
	"upserver/internal/pkg/utils/database"
)

type StatisticService struct {
}

var startTimeStr []byte
var startTime time.Time
var rollback bool //上次定时任务是否已执行成功
// Cron
// @description: 定时任务读取es日志写入mongodb
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/13 20:17
// @success:
func Cron() {
	endTime := time.Now()
	endTimeStr, _ := endTime.MarshalJSON()
	if !rollback {
		//上次执行成功，开始时间使用2分钟前的时间
		//startTime = endTime.Add(time.Second * time.Duration(utils.K8sConfig.K8s.Statistic.CrontabTime) * -1)
		//FIXME 此处为了测试修改时间
		startTime = endTime.Add(100 * time.Hour * time.Duration(utils.K8sConfig.K8s.Statistic.CrontabTime) * -1)
		startTimeStr, _ = startTime.MarshalJSON()
	}
	var ss StatisticService
	var from, num int
	//用于存放临时数据，来验证定时任务期间的数据 租户  业务  密码服务  密码资源
	statistics := make(map[string]statistic.StatisticsTable)
	var size = 1000
LOOP:
	res, err := ss.LatestQuery("mmyypt_app_events", "csmp", 0, string(startTimeStr), string(endTimeStr), from, size)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("定时任务查询es出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("es返回数据=============", string(res))
	var rr statistic.RealTimeResponse
	//FIXME 能解析，但是err有值？？
	//err = json.Unmarshal(res, &rr)
	json.Unmarshal(res, &rr)
	fmt.Println("len", len(rr.Hits.Hits))
	//sasa, _ := json.MarshalIndent(rr, "  ", "  ")
	//fmt.Println("解析后数据============", string(sasa))

	if err != nil {
		log.WithFields(utils.WriteDataLogs("定时任务查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("--", rr)
	if len(rr.Hits.Hits) == 0 {
		log.WithFields(utils.WriteDataLogs("定时任务查询查询的数据为空", err)).Warn(constant.Msg)
		return
	}
	//var statistics []statistic.StatisticsTable
	//var statistics map[string]statistic.StatisticsTable

	for _, v := range rr.Hits.Hits {
		md5 := utils.Md5(strconv.Itoa(v.Source.EventTid) + v.Source.EventAppid + strconv.Itoa(v.Source.EventDType) + v.Source.EventSerial)
		if _, ok := statistics[md5]; ok {
			sts := statistics[md5]
			sts.Total++
			sts.Flow += v.Source.EventData
			statistics[md5] = sts
		} else {
			sts := statistic.StatisticsTable{
				TenantId:     v.Source.EventTid,
				Appid:        v.Source.EventAppid,
				AppName:      v.Source.EventAppName,
				CipherType:   v.Source.EventDType,
				CipherSerial: v.Source.EventSerial,
				StartTime:    startTime,
				EndTime:      endTime,
				Total:        1,
				Flow:         v.Source.EventData,
				CreateTime:   time.Now(),
				EventTime:    v.Source.EventTime.Unix(),
			}
			statistics[md5] = sts
		}
	}
	if rr.Hits.Total > size {
		//查询结果大于1000条，需要分页查询计算
		//num = int(rr.Hits.Total/1000)+1
		num++
		if rr.Hits.Total > num*size {
			//继续分页查询
			from = num * size
			goto LOOP
		}
	}
	var statisticArr []interface{}
	for _, v := range statistics {
		statisticArr = append(statisticArr, v)
	}
	cli := database.GetMgoCli()
	fmt.Println("==========", utils.K8sConfig.K8s.Statistic.MongoDatabase)
	mgo := cli.Database(utils.K8sConfig.K8s.Statistic.MongoDatabase)
	collection := mgo.Collection(utils.K8sConfig.K8s.Statistic.Collection)
	_, err = collection.InsertMany(context.TODO(), statisticArr)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("数据写入mongodb失败", err)).Error(constant.Msg)
		rollback = true
		return
	} else {
		rollback = false
	}
	fmt.Println("cron success num=", len(statisticArr))
}

// AddCron
// @description: 添加定时任务
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/13 20:27
// @success:
func AddCron() {
	seconds := utils.K8sConfig.K8s.Statistic.CrontabTime
	//seconds = 120
	fmt.Println("定时任务设置时间seconds=", seconds)
	crontab.AddSecondFunc(seconds, Cron)
}

func EsDataMarshal(startTime, endTime string, from int) {

}

// InitCheckMongoData
// @description: 每次启动服务检测mongodb中是否有数据，是否重启后很久没更新
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/23 15:02
// @success:
func InitCheckMongoData() {
	//首次运行，全量获取数据

	//服务重启，只获取缺少数据

}
