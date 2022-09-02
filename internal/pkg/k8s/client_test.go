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
	res, _ := GetRepoList("test")
	fmt.Println(res)
}

func TestUpdateRepo(t *testing.T) {
	res, _ := UpdateRepo("dked", "repo-xrk8vj9942vy2p")
	fmt.Println(res)
}

func TestCreateWorkspaces(t *testing.T) {
	err := CreateWorkspaces("test", "test","测试创建企业空间")
	fmt.Println(err)
}
func TestProject(t *testing.T)  {
	err := CreateProject("test-project", "test-project","测试创建项目","test")
	fmt.Println(err)
}
func TestExportMysqlServer(t *testing.T) {
	ExportMysqlSecret("/tmp/k8s/version/","test-project","mmyypt_db","zf12345678")
}

func TestCreateRepos(t *testing.T) {
	var project = "library1"
	CreateRepos("test",project,"csmp")


}

func TestGetRepoApps(t *testing.T) {
	res,err :=GetRepoApps("repo-95x39n214x9p1o")
	fmt.Println("res",res,"    err= ",err)
}

func TestFile(t *testing.T) {
	yaml,err := Files("app-2wpvoo7pz18qp8","appv-750pkkp1n63z1v")
	fmt.Println("yaml",yaml,"    err= ",err)
}

func TestCreateApp(t *testing.T) {
	err := CreateApp("ww","namespace","app-name2")
	fmt.Println("创建应用结果：",err)
}