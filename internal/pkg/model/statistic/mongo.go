// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/14$ 15:14$

package statistic

type TimeSectionFilter struct {
	StartTime interface{} `bson:"start_time,omitempty"` //开始时间
	EndTime   interface{} `bson:"end_time,omitempty"`   //结束时间
}

type StatisticsFilter struct {
	ID           interface{} `bson:"_id,omitempty"`
	TenantId     interface{} `bson:"tenant_id,omitempty" json:"tenant_id"`
	Appid        interface{} `bson:"appid,omitempty" json:"appid"`
	CipherType   interface{} `bson:"cipher_type,omitempty" json:"cipher_type"`
	CipherSerial interface{} `bson:"cipher_serial,omitempty" json:"cipher_serial"`
	Total        interface{} `bson:"total,omitempty" json:"total" `
	Flow         interface{} `bson:"flow,omitempty" json:"flow"`
	StartTime    interface{} `bson:"start_time,omitempty" json:"start_time"` //开始时间
	EndTime      interface{} `bson:"end_time,omitempty" json:"end_time"`     //结束时间
}

type Lt struct {
	Lt string `bson:"$lt"` //小于
}

type Gte struct {
	Gte string `bson:"$gte"` //小于
}

type Group struct {
	Group GroupField `bson:"$group" json:"$group"` //分组
}
type GroupField struct {
	Field []string   `bson:"_id" json:"_id"`
	Total GroupTotal `bson:"total" json:"total"`
}
type GroupTotal struct {
	Total string `bson:"$sum" json:"$sum"`
}

type Sum struct {
	Sum interface{} `bson:"$sum" json:"$sum"` //求和
}

type Match struct {
	Match interface{} `bson:"$match,omitempty" json:"$match"` //匹配项
}

//MatchTime 时间范围
type MatchTime struct {
	Match StartTime `bson:"$match,omitempty" json:"$match"` //开始时间

}
type StartTime struct {
	EventTime StartTimeRange `bson:"event_time,omitempty" json:"event_time"` //开始时间
}
type StartTimeRange struct {
	Gte int64 `bson:"$gte" json:"$gte"` //大于等于
	Lt  int64 `bson:"$lt" json:"$lt"`   //小于
}

type MatchAndGroup struct {
	Match StartTime  `bson:"$match,omitempty" json:"$match"` //开始时间
	Group GroupField `bson:"$group" json:"$group"`           //分组
}

//MarshalMgoData 解析mongodb分组返回的数据
type MarshalMgoData struct {
	Id    []interface{} `json:"_id" bson:"_id"`
	Total int           `json:"total" bson:"total"`
}
