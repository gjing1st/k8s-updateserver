// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/9$ 17:48$

package statistic

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
	"upserver/internal/pkg/constant"
	"upserver/internal/pkg/model/statistic/response"
	"upserver/internal/pkg/utils"
	"upserver/internal/pkg/utils/database"
)

func TestName(t *testing.T) {
	collection := database.GetCollection()
	filter := bson.D{{"event_time", 1663229791}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.WithFields(utils.WriteDataLogs("查询mongodb失败", err)).Error(constant.Msg)
		return
	}
	var res []response.AppFlow
	if err = cursor.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println("res", res)

}

func TestGetAppFlow(t *testing.T) {
	flowTotal := ss.SumFlowAndTotal(time.Now().Add(time.Hour*-2000), time.Now(), 0, "30r998w6jtmpruoe", 2)
	fmt.Println("res==", flowTotal)
}
func TestGetFlowAndTotal(t *testing.T) {
	res := ss.GetFlowAndTotal(5, 0, "", 0)
	fmt.Println(res)
}
