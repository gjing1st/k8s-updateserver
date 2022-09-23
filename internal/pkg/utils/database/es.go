// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/13$ 20:28$

package database

import (
	"github.com/elastic/go-elasticsearch/v7"
	log "github.com/sirupsen/logrus"
	"upserver/internal/pkg/utils"
)

var esClient *elasticsearch.Client

// EsInit
// @description: 初始化es数据库
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/23 15:14
// @success:
func EsInit() {
	var err error
	//utils.K8sConfig.K8s.ElasticSearch.Address = "http://192.168.0.80:31199"
	cfg := elasticsearch.Config{
		Addresses: []string{
			utils.K8sConfig.K8s.ElasticSearch.Address,
		},
		// ...
	}
	esClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Panicf("Error creating the client: %s", err)
	}

}

// GetEsClient
// @description: 获取es客户端
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/23 15:15
// @success:
func GetEsClient() *elasticsearch.Client {
	if esClient == nil {
		EsInit()
	}
	return esClient
}
