package k8s

import (
	"encoding/json"
	"fmt"
	"testing"
	"upserver/internal/pkg/utils"
)

func TestClient(t *testing.T) {
	res, err := GetToken()
	fmt.Println("err=", err)
	fmt.Printf("%#v\n", res)
	fmt.Println("=========")
	res1, _ := GetToken()
	fmt.Println(res1)
	res2, _ := GetToken()
	fmt.Println(res2)
	//GenerateOauthToken()
}

func TestGetApp(t *testing.T) {
	GetApp("", "")
}
func TestGetVersions(t *testing.T) {
	apps, err := GetApp("", "")
	if err != nil {
		return
	}
	res, err := GetVersions(apps.Items[0].App.AppId)
	if err != nil {
		return
	}
	jon, _ := json.Marshal(res)
	fmt.Println(string(jon))
	if len(apps.Items) < 0 || len(res.Items) < 0 {
		return
	}
	//lastVersionId := res.Items[0].VersionId
	up := UpVersionReq{
		apps.Items[0].App.AppId,
		apps.Items[0].Cluster.ClusterId,
		"default",
		"",
		apps.Items[0].App.Name,
		utils.K8sConfig.K8s.Namespace,
		"admin",
		res.Items[0].VersionId,
		utils.K8sConfig.K8s.Workspace,
	}
	upJson, _ := json.Marshal(up)
	fmt.Println(string(upJson))
	res3, err := UpVersion(up)
	jsonRes, _ := json.Marshal(res3)
	fmt.Println(string(jsonRes))
}

func TestFiles(t *testing.T) {
	//v,_ := Files("app-4jylnmjkr95ow1","appv-wxnk29rv375zq9")
	//fmt.Println(v)
	file := "csmp-backend_V3.1.0.tar"
	f := utils.UnExt(file)
	fmt.Println(f)

}

func TestGetRepoList(t *testing.T) {
	res, _ := GetRepoList("dked")
	fmt.Println(res)
}

func TestUpdateRepo(t *testing.T) {
	res, _ := UpdateRepo("dked", "repo-xrk8vj9942vy2p")
	fmt.Println(res)
}
