// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 11:10$

package statistic

import "time"

//Kubernetes k8s相关信息
type Kubernetes struct {
	ContainerImage string `json:"container_image"`
	ContainerName  string `json:"container_name"`
	DockerId       string `json:"docker_id"`
	NamespaceName  string `json:"namespace_name"`
	PodName        string `json:"pod_name"`
}

type HitsBase struct {
	Total    int     `json:"total"`
	MaxScore float64 `json:"max_score"`
}

type ResponseBase struct {
	Took    int  `json:"took"`
	TimeOut bool `json:"time_out"`
}

//RealTimeResponse 实时调度查询返回数据
type RealTimeResponse struct {
	ResponseBase
	Hits RealTimeOutHits `json:"hits"`
}
type RealTimeOutHits struct {
	HitsBase
	Hits []RealTimeHits `json:"hits"`
}
type RealTimeHits struct {
	Id     string         `json:"_id"`
	Index  string         `json:"_index"`
	Type   string         `json:"_type"`
	Score  string         `json:"_score"`
	Source RealTimeSource `json:"_source"`
}
type RealTimeSource struct {
	Timestamp    string     `json:"@timestamp"`
	EventAppid   string     `json:"event_appid"`
	EventAppName string     `json:"event_appname"`
	EventClient  string     `json:"event_client"`
	EventData    int        `json:"event_data"`
	EventDType   int        `json:"event_dtype"`
	EventInfo    string     `json:"event_info"`
	EventSerial  string     `json:"event_serial"`
	EventTid     int        `json:"event_tid"`
	EventTime    time.Time  `json:"event_time"`
	EventType    int        `json:"event_type"` //对应接口
	Level        string     `json:"level"`
	Log          string     `json:"log"`
	Msg          string     `json:"msg"`
	Time         string     `json:"time"`
	Kubernetes   Kubernetes `json:"kubernetes"`
}
