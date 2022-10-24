// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/6$ 17:16$

package statistic

import (
	"encoding/json"
	"fmt"
	"statistic/internal/pkg/model/statistic"
	"statistic/internal/pkg/utils"
	"testing"
	"time"
)

var ss StatisticService

func init() {
	utils.InitConfig()
}
func TestInit(t *testing.T) {
	//query := esT{
	//	"query": esT{
	//		"match": esT{
	//			"msg": "mmyypt_app_events",
	//		},
	//	},
	//}
	//q := strings.Replace(tmpl.RealTime, "{{msgIndex}}", "mmyypt_app_events", 1)
	//var query map[string]interface{}
	//err := json.Unmarshal([]byte(q), &query)
	//fmt.Println("err", err)
	//fmt.Println(query)
	//Search("ks-logstash-log-2022.09.06", query)

	res, _ := ss.RealTimeQuery("mmyypt_app_events", "csmp", 0)
	//endTime, _ := time.Now().MarshalJSON()
	//tt := time.Now()
	//timeByte, _ := json.Marshal(tt)
	//res, _ := ss.RankingQuery("mmyypt_app_events", "csmp", 0, "2022-09-05T07:30:50.444Z", string(timeByte))
	//res, _ := ss.FlowQuery("mmyypt_app_events", "csmp", 0, "2022-08-06T07:30:50.444Z", "2022-09-07T07:18:48.712Z")
	//res, _ := ss.StatisticsQuery("mmyypt_app_events", "csmp", 0, "2022-08-06T07:30:50.444Z", string(endTime))
	//res, _ := ss.RankingQuery("mmyypt_app_events", "csmp",0, time.Now().Add(time.Duration(-1*time.Hour*33)), time.Now())
	var r statistic.RealTimeResponse
	json.Unmarshal(res, &r)
	sasa, _ := json.MarshalIndent(r, "  ", "  ")
	fmt.Println("============", string(sasa))
}

func TestTime(t *testing.T) {

	tt := time.Now()
	timeByte, _ := tt.MarshalJSON()
	fmt.Println(string(timeByte))
}

func TestLatest(t *testing.T) {
	now := time.Now()
	endTime, _ := now.MarshalJSON()
	startTime, _ := now.Add(time.Second * time.Duration(utils.Config.K8s.Statistic.CrontabTime) * -1).MarshalJSON()
	fmt.Println(string(endTime))
	fmt.Println(string(startTime))
	res, _ := ss.LatestQuery("mmyypt_app_events", "csmp", 0, string(startTime), string(endTime), 0, 10)
	var r statistic.RealTimeResponse
	json.Unmarshal(res, &r)
	sasa, _ := json.MarshalIndent(r, "  ", "  ")
	fmt.Println("============", string(sasa))
}

func TestRanking(t *testing.T) {

	ss.RankingByApp(time.Now().Add(time.Hour*-200), time.Now())
}

func TestStatisticService_CipherStatistic(t *testing.T) {
	ss.CipherStatistic(time.Now().Add(time.Hour*-200), time.Now())
}

func TestAppFlow(t *testing.T) {
	ss.AppFlow(time.Now().Add(time.Hour*-2000), time.Now())
}

func TestLastQuery(t *testing.T) {
	s := ss.LastData()
	fmt.Println(s)
}

type Ta struct {
	Name   string  `json:"name,omitempty"`
	Age    string  `json:"age,omitempty"`
	Person *Person `json:"person,omitempty"`
}

type Person struct {
	Sex   string `json:"sex,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func TestTa(t *testing.T) {
	var ta Ta
	ta = Ta{Name: "qq"}
	js, _ := json.Marshal(ta)
	fmt.Println(string(js))
}
