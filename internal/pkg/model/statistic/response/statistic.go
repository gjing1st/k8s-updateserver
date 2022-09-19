// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/15$ 18:29$

package response

import "time"

//AppFlow 业务流量
type AppFlow struct {
	AppName string `json:"app_name"`
	Flow    int    `json:"flow"`
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
}
