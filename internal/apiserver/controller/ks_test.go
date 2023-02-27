// Path: internal/apiserver/controller
// FileName: ks_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/2$ 17:34$

package controller

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"testing"
	"upserver/internal/pkg/harbor"
)

func TestVersioon(t *testing.T) {
	f, _ := ioutil.ReadFile("D:\\project\\upserver\\cmd\\upserver\\upVersionExample\\version_info.json")

	res := &harbor.VersionInfo{}
	json.Unmarshal(f, res)
	var req IpmReq
	var content = "<p>系统已升级至" + "V3.5" + "版本。</p>"
	req.Title = content
	if len(res.Content) > 0 {
		for i, v := range res.Content[0].Detail {
			content += "<p>" + strconv.Itoa(i) + ". " + v + ";</p>"
		}
	}
	req.Content = content
	PushUpdateVersionMsg(req)
}
