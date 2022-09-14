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

func EsInit() {
	var err error
	utils.K8sConfig.K8s.ElasticSearch.Address = "http://192.168.8.129:30637"
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

func GetEsClient() *elasticsearch.Client {
	if esClient == nil {
		EsInit()
	}
	return esClient
}
