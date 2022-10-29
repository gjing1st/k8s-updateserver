package harbor

import (
	"context"
	"errors"
	"fmt"
	"github.com/goharbor/go-client/pkg/harbor"
	"github.com/goharbor/go-client/pkg/sdk/assist/client/chart_repository"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/artifact"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/project"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/repository"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	log "github.com/sirupsen/logrus"
	"os"
	"upserver/internal/pkg/utils"
)

// Client
// @description: 获取harbor连接基本信息
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/25 14:24
// @success:
func Client() (*harbor.ClientSet, error) {
	c := &harbor.ClientSetConfig{
		URL:      utils.K8sConfig.Harbor.Address,
		Password: utils.K8sConfig.Harbor.Admin,
		Username: utils.K8sConfig.Harbor.Password,
		Insecure: true,
	}

	cs, err := harbor.NewClientSet(c)
	if err != nil {
		log.WithField("client error", err).Info("client set err")
		return nil, err
	}
	return cs, err
}

// UploadChart
// @description: helm chart上传至私有仓库
// @param: chartPath helmChart文件绝对路径
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/25 14:27
// @success:
func UploadChart(chartPath string) (bool, error) {
	c, err := Client()
	if err != nil {
		return false, err
	}
	clientV1 := c.Assist()
	file, err := os.Open(chartPath)
	defer file.Close()
	if err != nil {
		log.Println("file open err", err)
	}
	params := &chart_repository.PostChartrepoRepoChartsParams{
		Chart: file,
		//Repo:  constant.HarborProject,
		Repo: utils.K8sConfig.Harbor.Project,
	}

	ok, err := clientV1.ChartRepository.PostChartrepoRepoCharts(context.TODO(), params)
	if err != nil {
		log.WithField("err:", err).Info("helm 上传接口错误")
		return false, err
	}
	if ok.Error() != "" {
		log.Println("ok.err", ok.Error())
		return false, errors.New(ok.Error())
	}
	return true, nil
}

// ListRepositories
// @description: 镜像仓库列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/25 16:54
// @success:
func ListRepositories() (res *repository.ListRepositoriesOK, err error) {
	c, err := Client()
	if err != nil {
		return res, err
	}
	clientV2 := c.V2()
	params := &repository.ListRepositoriesParams{
		ProjectName: "csmp",
	}
	res, err = clientV2.Repository.ListRepositories(context.TODO(), params)
	if err != nil {
		log.WithField("err", err).Error("获取镜像仓库列表错误")
	}
	return
	//fmt.Println("err===", err)
	//fmt.Printf("%#v", ok.Payload)
	//for _, f := range ok.Payload {
	//	fmt.Printf("%#v\n", f)
	//}
	//return false, nil
}

// ListArtifacts
// @description: 镜像列表
// @param: 仓库名称
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/25 17:57
// @success:
func ListArtifacts(repoName string) (res *artifact.ListArtifactsOK, err error) {
	c, err := Client()
	if err != nil {
		return res, err
	}
	clientV2 := c.V2()
	params := &artifact.ListArtifactsParams{
		ProjectName:    "csmp",
		RepositoryName: repoName,
	}
	res, err = clientV2.Artifact.ListArtifacts(context.TODO(), params)
	if err != nil {
		log.WithField("err", err).Error("获取镜像列表错误")
	}
	return
	//fmt.Println("err===", err)
	//fmt.Printf("%#v", ok.Payload)
	//for _, f := range ok.Payload {
	//	for _, t := range f.Tags {
	//		fmt.Printf("%#v\n", t)
	//	}
	//
	//}
}

func CreateProject(projectName string) {
	c, err := Client()
	if err != nil {
		log.WithField("=============err", err)
	}
	clientV2 := c.V2()
	var public = true
	req := models.ProjectReq{
		ProjectName: projectName,
		Public:      &public,
	}
	params := &project.CreateProjectParams{
		Project: &req,
	}
	res, err := clientV2.Project.CreateProject(context.TODO(), params)
	fmt.Println("res", res,"err",err)
}
