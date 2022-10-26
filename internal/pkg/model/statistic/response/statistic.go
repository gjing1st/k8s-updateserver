// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/15$ 18:29$

package response

import "time"

//AppFlow 业务流量
type AppFlow struct {
	AppName string `json:"app_name"`
	Flow    int    `json:"flow" bson:"flow"`
}

type FlowAndTotal struct {
	Time  []time.Time `json:"time"` //x轴为各时间节点
	Flow  []int       `json:"flow"`
	Total []int       `json:"total"`
}

//AppStatistic 业务调用服务排名
type AppStatistic struct {
	AppName string `json:"app_name"`
	Total   int    `json:"total"`
}

//CipherStatistic 密码服务使用统计
type CipherStatistic struct {
	Server string `json:"server"`
	Total  int    `json:"total"`
}

type Realtime struct {
	EventAppid   string    `json:"event_appid"`
	EventAppName string    `json:"event_appname"`
	EventClient  string    `json:"event_client"`
	EventData    int       `json:"event_data"`
	EventDType   int       `json:"event_dtype"`
	EventInfo    string    `json:"event_info"`
	EventSerial  string    `json:"event_serial"`
	EventServer  string    `json:"event_server"`
	EventTid     int       `json:"event_tid"`
	EventTime    time.Time `json:"event_time"`
	EventType    int       `json:"event_type"`
	EventApiName string    `json:"event_api_name"`
	EventResult  bool      `json:"event_result"`
}

type Ranking struct {
	X []int         `json:"x"`
	Y []interface{} `json:"y"`
}

type Calculate struct {
	Name string     `json:"name"`
	Data []ApiValue `json:"data"`
}
type ApiValue struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

//CipherApiTotal 服务次数
type CipherApiTotal struct {
	Time  []time.Time `json:"time"` //x轴为各时间节点
	Total interface{} `json:"total"`
}

type CipherUsage struct {
	Key      string  `json:"key"`
	DocCount float64 `json:"doc_count"`
}
