// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/6$ 17:16$

package statistic

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var ss StatisticService

func TestInit(t *testing.T) {
	Init()
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
	t1, _ := json.Marshal(time.Now())
	fmt.Println(string(t1))

	//res, _ := ss.RankingQuery("mmyypt_app_events", 3, "2022-09-06T07:30:50.444Z", "2022-09-07T07:18:48.712Z")
	//res, _ := ss.FlowQuery("mmyypt_app_events", 3, "2022-08-06T07:30:50.444Z", "2022-09-07T07:18:48.712Z")
	res, _ := ss.StatisticsQuery("mmyypt_app_events", 3, "2022-08-06T07:30:50.444Z", "2022-09-07T07:18:48.712Z")
	//res, _ := ss.RankingQuery("mmyypt_app_events", 3, time.Now().Add(time.Duration(-1*time.Hour*33)), time.Now())
	fmt.Println("res", res)

}

func TestTime(t *testing.T) {

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			fmt.Println("执行的业务逻辑")
		}

	}()

	select {}
}
