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
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/tmpl"
	"upserver/internal/pkg/utils"
)

// Search
// @description: 去es查询数据
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/7 14:54
// @success:
func Search(query esType) (string, error) {
	res, err := EsClient.Info()
	if err != nil {
		log.WithFields(utils.WriteDataLogs("数据查询失败", "Error getting response")).Error(constant.Msg)

		return "", errors.New("数据查询失败")
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

		return "", errors.New("数据查询失败")
	}
	// Perform the search request.
	res, err = EsClient.Search(
		EsClient.Search.WithContext(context.Background()),
		//EsClient.Search.WithIndex(index),
		EsClient.Search.WithBody(&buf),
		EsClient.Search.WithTrackTotalHits(true),
		//EsClient.Search.WithFrom(0),
		//EsClient.Search.WithSize(10),
		//EsClient.Search.WithSort("time:desc"),
		EsClient.Search.WithPretty(),
	)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("数据查询失败", "Error getting response")).Error(constant.Msg)

		return "", errors.New("数据查询失败")
	}
	if res.StatusCode != http.StatusOK {
		log.WithFields(utils.WriteDataLogs("数据查询失败", err)).Error(constant.Msg)
		return "", errors.New("数据查询失败")
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("数据解析失败", err)).Error(constant.Msg)

		return "", errors.New("数据查询失败")
	}
	return string(data), nil
}

// RealTimeQuery
// @description: 实时调度查询
// @param: msgIndex string 索引id，每个项目一个固定的 csmp的为mmyypt_app_events
// @param: eventTid int 租户id
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/7 10:35
// @success:
func (ss *StatisticService) RealTimeQuery(msgIndex string, eventTid int) (res string, err error) {
	q := strings.Replace(tmpl.RealTime, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)

	var query esType
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	//fmt.Printf("#########%#v\n", query)
	if eventTid == 0 {
		query["query"].(esType)["bool"].(esType)["must"] = query["query"].(esType)["bool"].(esType)["must"].([]interface{})[0]
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
func (ss *StatisticService) RankingQuery(msgIndex string, eventTid int, startTime, endTime string) (res string, err error) {
	q := strings.Replace(tmpl.Ranking, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//startByte, _ := json.Marshal(startTime)
	q = strings.Replace(q, "{{gtField}}", startTime, 1)
	//endByte, _ := json.Marshal(endTime)
	q = strings.Replace(q, "{{ltField}}", endTime, 1)
	//fmt.Println("#########", q)
	var query esType
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	if eventTid == 0 {
		query["query"].(esType)["bool"].(esType)["must"] = query["query"].(esType)["bool"].(esType)["must"].([]interface{})[0]
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
func (ss *StatisticService) FlowQuery(msgIndex string, eventTid int, startTime, endTime string) (res string, err error) {
	q := strings.Replace(tmpl.Flow, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//startByte, _ := json.Marshal(startTime)
	q = strings.Replace(q, "{{gtField}}", startTime, 1)
	//endByte, _ := json.Marshal(endTime)
	q = strings.Replace(q, "{{ltField}}", endTime, 1)
	//fmt.Println("#########", q)
	var query esType
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	if eventTid == 0 {
		query["query"].(esType)["bool"].(esType)["must"] = query["query"].(esType)["bool"].(esType)["must"].([]interface{})[0]
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
func (ss *StatisticService) StatisticsQuery(msgIndex string, eventTid int, startTime, endTime string) (res string, err error) {
	q := strings.Replace(tmpl.Statistics, "{{msgField}}", msgIndex, 1)
	eventTidStr := strconv.Itoa(eventTid)
	q = strings.Replace(q, "{{eventTidField}}", eventTidStr, 1)
	//startByte, _ := json.Marshal(startTime)
	q = strings.Replace(q, "{{gtField}}", startTime, 1)
	//endByte, _ := json.Marshal(endTime)
	q = strings.Replace(q, "{{ltField}}", endTime, 1)
	//fmt.Println("#########", q)
	var query esType
	err = json.Unmarshal([]byte(q), &query)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("实时查询请求数据解析json出错", err)).Error(constant.Msg)
		return
	}
	//fmt.Println("====================")
	if eventTid == 0 {
		query["query"].(esType)["bool"].(esType)["must"] = query["query"].(esType)["bool"].(esType)["must"].([]interface{})[0]
	}

	res, _ = Search(query)

	return
}
