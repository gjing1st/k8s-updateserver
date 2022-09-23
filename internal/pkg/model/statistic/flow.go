// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 14:57$

package statistic

//FlowResponse 业务流量
type FlowResponse struct {
	FlowAggregations FlowAggregations `json:"aggregations"`
	ResponseBase
}

type FlowAggregations struct {
	EventAppName EventAppName `json:"event_appname"`
}
type EventAppName struct {
	Buckets       []FlowBuckets `json:"buckets"`
	DocCount      int           `json:"doc_count_error_upper_bound"`
	SumOtherCount int           `json:"sum_other_doc_count"`
}
type FlowBuckets struct {
	Key      string `json:"key"`
	DocCount int    `json:"doc_count"`
	Flow     struct {
		Value int `json:"value"`
	} `json:"flow"`
}
