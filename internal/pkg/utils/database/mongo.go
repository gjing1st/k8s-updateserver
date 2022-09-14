// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 17:10$

package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func InitMgo() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		InitMgo()
	}
	return mgoCli
}
