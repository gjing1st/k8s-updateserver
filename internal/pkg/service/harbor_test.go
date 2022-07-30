package service

import "testing"


func TestAddRepo(t *testing.T) {
	var has HarborService
	//has.AddHelmRepo()
	has.HelmChartPush("/home/hello-0.1.0.tgz")
}

func TestInit(t *testing.T) {
	Init()
}