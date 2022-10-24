// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/12$ 17:36$

package statistic

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// PackageMatch
// @description: 组装where查询条件
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/12 17:39
// @success:
func PackageMatch(startTime, endTime time.Time, tenantId int, appid string, cipherType int, cipherSerial string) (filter bson.D) {
	//查询条件
	//时间范围
	var eventTime bson.E
	eventTime = bson.E{"event_time", bson.D{{"$gte", startTime.Unix()}, {"$lt", endTime.Unix()}}}
	filter = append(filter, eventTime)
	if tenantId > 0 {
		filter = append(filter, bson.E{"tenant_id", tenantId})
	}
	if appid != "" {
		filter = append(filter, bson.E{"appid", appid})
	}
	if cipherType > 0 {
		filter = append(filter, bson.E{"cipher_type", cipherType})
	}
	if cipherSerial != "" {
		filter = append(filter, bson.E{"cipher_serial", cipherSerial})
	}
	return filter
}
