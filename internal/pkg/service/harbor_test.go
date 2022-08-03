package service

import (
	"fmt"
	"testing"
)


func TestAddRepo(t *testing.T) {
	var has HarborService
	//has.AddHelmRepo()
	has.HelmChartPush("/home/hello-0.1.0.tgz")
}

func TestInit(t *testing.T) {
	Init()
}

func TestK8sService_UnzipAndPush(t *testing.T) {
	K8sService{}.UnzipAndPush()
}

func TestK8sService_GetAppAndVersion(t *testing.T) {
	res,_ := K8sService{}.GetAppAndVersion("dked","csmp")
	fmt.Printf("%#v\n",res.Version.Items[0].Name)
	fmt.Printf("%#v",res)
}