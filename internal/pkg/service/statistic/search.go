// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/6$ 16:20$

package statistic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"statistic/internal/pkg/constant"
	"statistic/internal/pkg/tmpl"
	"statistic/internal/pkg/utils"
	"statistic/internal/pkg/utils/database"
	"strconv"
	"strings"
)

// Search
// @description: 去es查询数据
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/7 14:54
// @success:
func Search(query map[string]interface{}) ([]byte, error) {
	res, err := database.GetEsClient().Info()
	if err != nil {
		log.WithFields(utils.WriteDataLogs("数据查询失败", "Error getting response")).Error(constant.Msg)

		return nil, errors.New("数据查询失败")
	}
	//fmt.Println(res.String())
	// search - highlight
	var buf bytes.Buffer
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": title,
	//		},
	//	},
	//	"highlight": map[string]interface{}{
	//		"pre_tags":  []string{"<font color='red'>"},
	//		"post_tags": []string{"</font>"},
	//		"fields": map[string]interface{}{
	//			"title": map[string]interface{}{},
	//		},
	//	},
	//}
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		log.WithFields(utils.WriteDataLogs("数据查询失败", "Error encoding query")).Error(constant.Msg)

		return nil, errors.New("数据查询失败")
	}
	// Perform the search request.
	res, err = database.GetEsClient().Search(
		database.GetEsClient().Search.WithContext(context.Background()),
		//EsClient.Search.WithIndex(index),
		database.GetEsClient().Search.WithBody(&buf),
		database.GetEsClient().Search.WithTrackTotalHits(true),
		//EsClient.Search.WithFrom(0),
		//EsClient.Search.WithSize(10),
		//EsClient.Search.WithSort("time:desc"),
		database.GetEsClient().Search.WithPretty(),
	)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("数据查询失败", "Error getting response")).Error(constant.Msg)

		return nil, errors.New("数据查询失败")
	}
	if res.StatusCode != http.StatusOK {
		log.WithFields(utils.WriteDataLogs("数据查询失败", err)).Error(constant.Msg)
		return nil, errors.New("数据查询失败")
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("数据解析失败", err)).Error(constant.Msg)

		return nil, errors.New("数据查询失败")
	}
	//var r statistic.RealTimeResponse
	//err = json.Unmarshal(data, &r)
	//fmt.Println("========err", err)
	//fmt.Println("============", r)
	//fmt.Println("============")
	//fmt.Printf("%#v\n", r)
	//by, err := json.MarshalIndent(string(data), "", "")
	//fmt.Println("========err", err)
	//err = json.Unmarshal(by, &r)
	//fmt.Println("============", r)
	//fmt.Println("============")
	//aa, _ := json.Marshal(r)
	//fmt.Println("##########", string(aa))
	return data, nil
}

// RealTimeQuery
// @description: 实时调度查询
// @param: msgIndex string 索引id，每个项目一个固定的 csmp的为mmyypt_app_events
// @param: eventTid int 租户id
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/7 10:35
// @success:
func (ss *StatisticService) RealTimeQuery(msgIndex, nameSpace string, eventTid int) (res []byte, err error) {
	q := strings.Replace(tmpl.RealTime, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	q = strings.Replace(q, "{{nameSpaceField}}", nameSpace, 1)

	var query map[string]interface{}
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	//fmt.Printf("#########%#v\n", query)
	if eventTid == 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{})[0]
		//delete(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), "query")
	}
	//fmt.Printf("#########%#v\n", query)
	res, _ = Search(query)

	return
}

// RankingQuery
// @description: 排名查询
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/7 15:54
// @success:
func (ss *StatisticService) RankingQuery(msgIndex, nameSpace string, eventTid int, startTime, endTime string) (res []byte, err error) {
	q := strings.Replace(tmpl.Ranking, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//startByte, _ := json.Marshal(startTime)
	q = strings.Replace(q, "{{gtField}}", startTime, 1)
	//endByte, _ := json.Marshal(endTime)
	q = strings.Replace(q, "{{ltField}}", endTime, 1)
	q = strings.Replace(q, "{{nameSpaceField}}", nameSpace, 1)

	fmt.Println("#########", q)
	var query map[string]interface{}
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	if eventTid == 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{})[0]
	}

	res, _ = Search(query)

	return
}

// FlowQuery
// @description: 业务流量
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/7 16:39
// @success:
func (ss *StatisticService) FlowQuery(msgIndex, nameSpace string, eventTid int, startTime, endTime string) (res []byte, err error) {
	q := strings.Replace(tmpl.Flow, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//startByte, _ := json.Marshal(startTime)
	q = strings.Replace(q, "{{gtField}}", startTime, 1)
	//endByte, _ := json.Marshal(endTime)
	q = strings.Replace(q, "{{ltField}}", endTime, 1)
	q = strings.Replace(q, "{{nameSpaceField}}", nameSpace, 1)

	//fmt.Println("#########", q)
	var query map[string]interface{}
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	if eventTid == 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{})[0]
	}

	res, _ = Search(query)

	return
}

// StatisticsQuery
// @description: 使用统计
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/7 17:37
// @success:
func (ss *StatisticService) StatisticsQuery(msgIndex, nameSpace string, eventTid int, startTime, endTime string) (res []byte, err error) {
	q := strings.Replace(tmpl.Statistics, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//startByte, _ := json.Marshal(startTime)
	q = strings.Replace(q, "{{gtField}}", startTime, 1)
	//endByte, _ := json.Marshal(endTime)
	q = strings.Replace(q, "{{ltField}}", endTime, 1)
	q = strings.Replace(q, "{{nameSpaceField}}", nameSpace, 1)
	//fmt.Println("#########", q)
	var query map[string]interface{}
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	if eventTid == 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{})[0]
	}

	res, _ = Search(query)

	return
}

// LatestQuery
// @description: 某个时间段的数据
// @param: msgIndex string 每个系统独有的标识
// @param: nameSpace string 日志所在k8s中的命名空间，用来查询指定项目中的日志(比如:大数据局csmp和测试环境csmp)
// @param: eventTid int 租户id
// @param: startTime string 查询区间范围开始时间
// @param: endTime string 查询区间范围结束时间
// @param: fromField int 分页查询起始条数
// @param: sizeField int 分页条数
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/13 18:47
// @success:
func (ss *StatisticService) LatestQuery(msgIndex, nameSpace string, eventTid int, startTime, endTime string, fromField, sizeField int) (res []byte, err error) {
	q := strings.Replace(tmpl.Latest, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//startByte, _ := json.Marshal(startTime)
	q = strings.Replace(q, "{{gtField}}", startTime, 1)
	//endByte, _ := json.Marshal(endTime)
	q = strings.Replace(q, "{{ltField}}", endTime, 1)
	q = strings.Replace(q, "{{nameSpaceField}}", nameSpace, 1)
	q = strings.Replace(q, "{{fromField}}", strconv.Itoa(fromField), 1)
	q = strings.Replace(q, "{{sizeField}}", strconv.Itoa(sizeField), 1)
	fmt.Println("#########", q)
	var query map[string]interface{}
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	if eventTid == 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{})[:3]
	}
	fmt.Println("====================", query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"])

	res, _ = Search(query)

	return
}

// RealTimeLogQuery
// @description: 交易统计查询
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/17 16:43
// @success:
func (ss *StatisticService) RealTimeLogQuery(msgIndex, nameSpace, filter, cipherType string, eventTid, size, page int, startTime, endTime string) (res []byte, err error) {
	q := "" //去es查询条件
	//if filter == "" && startTime == "" && cipherType == 0 {
	//	//没有搜索条件,查全部
	//	q = strings.Replace(tmpl.RealTimeLogAll, "{{msgField}}", msgIndex, 1)
	//	eventTidStr := strconv.Itoa(eventTid)
	//	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//	q = strings.Replace(q, "{{nameSpaceField}}", nameSpace, 1)
	//}
	if q == "" {
		//说明有查询条件
		q = tmpl.RealTimeLogCipherFilter
	}
	if startTime != "" {
		//有时间范围查询
		timeRange := `,
                {
                    "range": {
                        "event_time": {
                            "gte": "{{startTime}}",
                            "lt": "{{endTime}}"
                        }
                    }
                }`
		timeRange = strings.Replace(timeRange, "{{startTime}}", startTime, 1)
		timeRange = strings.Replace(timeRange, "{{endTime}}", endTime, 1)
		q = strings.Replace(q, "{{rangeTime}}", timeRange, 1)
	} else {
		//没有时间查询

		q = strings.Replace(q, "{{rangeTime}}", "", 1)

	}
	//设置类型
	if cipherType != "" {
		cipherTypeQuery := `,
                {
                    "terms": {
                        "event_dtype": [
                            {{cipherType}}
                        ]
                    }
                }`

		cipherTypeQuery = strings.Replace(cipherTypeQuery, "{{cipherType}}", cipherType, 1)
		q = strings.Replace(q, "{{cipherType}}", cipherTypeQuery, 1)

	} else {
		q = strings.Replace(q, "{{cipherType}}", "", 1)
	}
	//搜索条件
	if filter != "" {
		var filterQuery string
		apiArr := utils.SearchApiType(filter)
		if len(apiArr) > 0 {
			//需要查询接口
			filterQuery = `,
                {
                    "terms": {
                        "event_type": [
                            {{apiTypeStr}}
                        ]
                    }
                }`
			var apiTypeStr string
			for _, i := range apiArr {
				apiTypeStr += "," + strconv.Itoa(i)
			}
			apiTypeStr = strings.Trim(apiTypeStr, ",")
			filterQuery = strings.Replace(filterQuery, "{{apiTypeStr}}", apiTypeStr, 1)
			//filterQuery = strings.Replace(filterQuery, "{{apiType}}", apiQuery, 1)

		} else {
			//没有对应接口需要查询
			//先拼接业务查询
			filterQuery = `,{
                    "match": {
                        "event_appname": "{{appName}}"
                    }
                }`
			filterQuery = strings.Replace(filterQuery, "{{appName}}", filter, 1)
		}

		q = strings.Replace(q, "{{filter}}", filterQuery, 1)

	} else {
		q = strings.Replace(q, "{{filter}}", "", 1)

	}

	//租户查询
	if eventTid != 0 {
		eventTidQuery := `,{
                    "match_phrase": {
                        "event_tid": {{eventTidField}}
                    }
                }`
		eventTidQuery = strings.Replace(eventTidQuery, "{{eventTidField}}", strconv.Itoa(eventTid), 1)
		q = strings.Replace(q, "{{eventTid}}", eventTidQuery, 1)
	} else {
		q = strings.Replace(q, "{{eventTid}}", "", 1)
	}

	if size == 0 {
		size = 10
	}
	var fromSize int
	if page != 0 {
		fromSize = size * (page - 1)
	}
	q = strings.Replace(q, "{{fromSize}}", strconv.Itoa(fromSize), 1)
	q = strings.Replace(q, "{{size}}", strconv.Itoa(size), 1)

	q = strings.Replace(q, "{{msgField}}", msgIndex, 1)

	q = strings.Replace(q, "{{nameSpaceField}}", nameSpace, 1)
	//fmt.Println(" ", q)

	var query map[string]interface{}
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	//fmt.Printf("#########%#v\n", query)
	//fmt.Printf("#########%#v\n", query)
	res, _ = Search(query)

	return
}
