// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 14:01$

package statistic

import "time"

//StatisticsResponse 统计数据
type StatisticsResponse struct {
	Aggregations StatisticsAggregations `json:"aggregations"`
	Hits         StatisticsHits         `json:"hits"`
	ResponseBase
}

type StatisticsHits struct {
	HitsBase
	Hits []map[string]interface{} `json:"hits"`
}

type StatisticsAggregations struct {
	EventSerial   EventSerial `json:"event_serial,event_dtype"`
	DocCount      int         `json:"doc_count_error_upper_bound"`
	SumOtherCount int         `json:"sum_other_doc_count"`
}
type EventSerial struct {
	Buckets Buckets `json:"buckets"`
}
type Buckets struct {
	DocCount int    `json:"doc_count"`
	Key      string `json:"key"`
}

type StatisticsTable struct {
	TenantId     int       `json:"tenant_id" bson:"tenant_id"`         //租户id
	Appid        string    `json:"appid" bson:"appid"`                 //业务名称
	AppName      string    `json:"appname" bson:"app_name"`            //业务id
	CipherType   int       `json:"cipher_type" bson:"cipher_type"`     //密码服务
	CipherSerial string    `json:"cipher_serial" bson:"cipher_serial"` //密码资源
	Total        int64     `json:"total" bson:"total"`                 //调用次数
	Flow         int       `json:"flow" bson:"flow"`                   //流量
	StartTime    time.Time `json:"start_time" bson:"start_time"`
	EndTime      time.Time `json:"end_time" bson:"end_time"`
	CreateTime   time.Time `json:"create_time" bson:"create_time"`
	EventTime    int64     `json:"event_time" bson:"event_time"`
}
