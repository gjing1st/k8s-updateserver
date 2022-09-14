// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 14:01$

package statistic

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
	TenantId     int    `json:"tenant_id"`     //租户id
	Appid        string `json:"appid"`         //业务id
	CipherType   int    `json:"cipher_type"`   //密码服务
	CipherSerial string `json:"cipher_serial"` //密码资源
	Total        int64  `json:"total"`         //次数
	Flow         int    `json:"flow"`          //流量
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}
