// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 13:56$

package statistic

//RankingResponse 排名
type RankingResponse struct {
	Aggregations RankingAggregations `json:"aggregations"`
	ResponseBase
}
type RankingAggregations struct {
	EventAppName RankingEventAppName `json:"event_appname"`
}
type RankingEventAppName struct {
	Buckets       []Buckets `json:"buckets"`
	DocCount      int       `json:"doc_count_error_upper_bound"`
	SumOtherCount int       `json:"sum_other_doc_count"`
}
