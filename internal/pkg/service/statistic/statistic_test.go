// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/9$ 17:48$

package statistic

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"statistic/internal/pkg/constant"
	"statistic/internal/pkg/model/statistic/response"
	"statistic/internal/pkg/utils"
	"statistic/internal/pkg/utils/database"
	"testing"
	"time"
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
	flowTotal := ss.SumFlowAndTotal(time.Now().Add(time.Hour*-2000), time.Now(), 0, "30r998w6jtmpruoe", 2, "")
	fmt.Println("res==", flowTotal)
}
func TestGetFlowAndTotal(t *testing.T) {
	res := ss.GetFlowAndTotal(5, 0, "", 0, "")
	fmt.Println(res)
}

func TestStatisticService_SumFlowGroupCipherType(t *testing.T) {
	res := ss.SumTotalGroupCipherType(time.Now().Add(time.Hour*-2000), time.Now(), 0, "", 0, "")
	fmt.Println(res)
}

func TestApi(t *testing.T) {
	res := ss.SumFlowGroupFiled(time.Now().Add(time.Hour*-2000), time.Now(), 0, "", 0, "", "")
	res1 := ss.SumTotalGroupFiled(time.Now().Add(time.Hour*-2000), time.Now(), 0, "", 0, "", "")
	fmt.Println(res, "total=", res1)
}

func TestSumGroupField(t *testing.T) {
	res := ss.SumGroupField(time.Now().Add(time.Hour*-2000), time.Now(), 0, "", 0, "", "$tenant_id", "$total")
	fmt.Println(res)
}

func TestCalculateFlowOrTotal(t *testing.T) {
	res, _ := ss.CalculateFlowOrTotal(1, 2, 0, "", 0, "")
	s, _ := json.Marshal(res)
	fmt.Println(res)
	fmt.Println(string(s))
}
