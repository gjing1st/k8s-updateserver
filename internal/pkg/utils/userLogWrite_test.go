package utils

import(
	"time"
	"testing"
	"upserver/internal/pkg/model"
)

func TestLogWrite(t *testing.T) {
	

	// req := map[string]interface{}{
	// 	"type":1,
	// 	"adminname":"w",
	// 	"info":"test========updateServer",
	// 	"tenantId":0,
	// 	"ip":"127.0.0.1",
	// }

	req := model.LogWriteReq{
		Type : 1,
		OpName: "w",
		Info : "test--UserWrite BASE",
		TenantId: 0,
		Ip: "localhost",
	}

	sc ,err := UserLogWrite(req,time.Second*5)

	t.Log(sc,err)
}