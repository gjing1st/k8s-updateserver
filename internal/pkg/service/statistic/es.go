// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/6$ 16:42$

package statistic

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"log"
	"upserver/internal/pkg/utils"
)

type StatisticService struct {
}

type esType map[string]interface{}

var EsClient *elasticsearch.Client

func Init() {
	var err error
	utils.K8sConfig.K8s.ElasticSearch.Address = "http://192.168.0.80:31199"
	cfg := elasticsearch.Config{
		Addresses: []string{
			utils.K8sConfig.K8s.ElasticSearch.Address,
		},
		// ...
	}
	EsClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Panicf("Error creating the client: %s", err)
	}

	//res, err := es.Info()
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//fmt.Println("====res", res)
	//defer res.Body.Close()
	//// Check response status
	//if res.IsError() {
	//	log.Fatalf("Error: %s", res.String())
	//
	//}
	//if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	//	log.Fatalf("Error parsing the response body: %s", err)
	//}
	//// Print client and server version numbers.
	//log.Printf("Client: %s", elasticsearch.Version)
	//log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	//log.Println(strings.Repeat("~", 37))
}
