// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 17:10$

package database

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"upserver/internal/pkg/utils"
)

var mgoCli *mongo.Client

// InitMgo
// @description: 初始化mongodb
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/23 15:11
// @success:
func InitMgo() {
	var err error
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", utils.K8sConfig.K8s.Statistic.MongoHost))

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMongo + "数据库连接失败")
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMongo + "连接失败")
	}
}

// GetMgoCli
// @description: 获取mongodb
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/23 15:12
// @success:
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		InitMgo()
	}
	return mgoCli
}

// GetCollection
// @description: 返回要操作的表
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/8 10:38
// @success:
func GetCollection() *mongo.Collection {
	cli := GetMgoCli()
	mgo := cli.Database(utils.K8sConfig.K8s.Statistic.MongoDatabase)
	collection := mgo.Collection(utils.K8sConfig.K8s.Statistic.Collection)
	return collection
}
