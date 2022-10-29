package k8s

import (
	"encoding/base64"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"sync"
	"time"
	"upserver/internal/pkg/utils"
)

var tokenMap sync.Map

// Set
// @description: 内存变量过期 类redis
// @param: key 变量名
// @param: value 变量值
// @param: exp 过期时间
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/27 17:20
// @success:
func Set(key, value interface{}, exp time.Duration) {
	tokenMap.Store(key, value)
	time.AfterFunc(exp, func() {
		tokenMap.Delete(key)
	})
}

// GetToken
// @description: 获取k8s token
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/27 17:12
// @success:
func GetToken() (string, error) {
	//查看缓存中是否已有token
	token, ok := tokenMap.Load("access_token")
	if ok {
		return utils.String(token), nil
	}

	var res *TokenResponse
	reqData := url.Values{}
	reqData.Add("grant_type", "password")
	reqData.Add("username", utils.K8sConfig.K8s.Username)
	reqData.Add("password", utils.K8sConfig.K8s.Password)
	reqData.Add("client_id", "kubesphere")
	reqData.Add("client_secret", "kubesphere")

	reqUrl := utils.K8sConfig.K8s.Url + "/oauth/token"
	httpCode, err1 := TokenRequestTimeout("POST", reqUrl, reqData, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithFields(log.Fields{"err": err1, "reqData": reqData, "reqUrl": reqUrl}).Info("获取k8s token失败")
		return "", errors.New("token获取失败")
	}
	//log.WithField("res", res).Info("获取成功")
	//写入缓存，过期时间减半
	//此处已加入Bearer 获取后直接使用
	token = "Bearer " + res.AccessToken
	Set("access_token", token, time.Second*time.Duration(res.ExpiresIn/2))
	return utils.String(token), nil
}

// GetApp
// @description: 获取应用列表
// @param: workspace 企业空间
// @param: namespace 项目/命名空间
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/27 18:19
// @success:
func GetApp(workspace, namespace string) (res *AppListResponse, err error) {
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/workspaces/" +
		workspace + "/namespaces/" + namespace + "/applications"
	httpCode, err1 := JsonRequestTimeout("GET", reqUrl, nil, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithField("err", err1).Info("获取应用列表失败")
		return res, errors.New("token获取失败")
	}
	return
}

// GetVersions
// @description: 获取应用的所有版本信息
// @param: appid 应用id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/27 18:27
// @success:
func GetVersions(appid string) (res *VersionResponse, err error) {
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/apps/" + appid + "/versions" + "?orderBy=sequence&reverse=true"
	httpCode, err1 := JsonRequestTimeout("GET", reqUrl, nil, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithField("err", err1).Info("获取应用列表失败")
		return res, errors.New("token获取失败")
	}
	return
}

// Files
// @description: 获取版本文件，主要提取values.yaml文件数据
// @param: appid 应用id
// @param: versionId 要获取的版本id
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/28 10:39
// @success:
func Files(appid, versionId string) (valuesYaml string, err error) {
	var res *FilesResponse
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/apps/" + appid + "/versions/" + versionId + "/files"
	httpCode, err1 := JsonRequestTimeout("GET", reqUrl, nil, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithField("err", err1).Info("获取版本文件失败")
		return "", errors.New("token获取失败")
	}
	vaByte, err := base64.StdEncoding.DecodeString(res.Files.ValuesYaml)
	if err != nil {
		log.WithField("err", err).Error("版本文件values.yaml解析失败")
		return "", err
	}
	valuesYaml = string(vaByte)
	return
}

// UpVersion
// @description: 更新到指定版本
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/7/28 9:19
// @success:
func UpVersion(up UpVersionReq) (message string, err error) {
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/workspaces/" +
		up.Workspace + "/namespaces/" + up.Namespace + "/applications/" + up.ClusterId
	//reqData := url.Values{}
	//reqData.Add("app_id", up.AppId)
	//reqData.Add("cluster", up.Cluster)
	//reqData.Add("cluster_id", up.ClusterId)
	//reqData.Add("conf", up.Conf)
	//reqData.Add("name", up.Name)
	//reqData.Add("namespace", up.Namespace)
	//reqData.Add("owner", up.Owner)
	//reqData.Add("version_id", up.VersionId)
	//reqData.Add("workspace", up.Workspace)
	var res *MessageResponse
	httpCode, err1 := JsonRestRequestTimeout("POST", reqUrl, up, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithFields(log.Fields{"err": err1, "reqUrl": reqUrl, "res": res}).Info("升级失败")
		return "", errors.New("升级失败")
	}
	message = res.Message
	return
}

// GetRepoList
// @description: 获取应用仓库列表
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/4 9:18
// @success:
func GetRepoList(workspaces string) (res *AppRepoResponse, err error) {
	if workspaces == "" {
		workspaces = utils.K8sConfig.K8s.Workspace.Name
	}
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/workspaces/" + workspaces + "/repos"
	httpCode, err1 := JsonRequestTimeout("GET", reqUrl, nil, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithField("err", err1).Info("获取应用仓库列表失败")
		return res, errors.New("获取应用仓库列表失败")
	}

	return
}

// UpdateRepo
// @description: 更新仓库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/4 9:48
// @success:
func UpdateRepo(workspaces, repoId string) (message string, err error) {
	if workspaces == "" {
		workspaces = utils.K8sConfig.K8s.Workspace.Name
	}
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/workspaces/" + workspaces + "/repos/" + repoId + "/action"
	var reqData UpdateRequest
	reqData.Action = "index"
	var res *MessageResponse
	httpCode, err1 := JsonRestRequestTimeout("POST", reqUrl, reqData, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithFields(log.Fields{"err": err1, "reqUrl": reqUrl, "res": res}).Info("升级失败")
		return "", errors.New("升级失败")
	}
	message = res.Message
	return
}

// GetAndUpdateRepo
// @description: 获取应用仓库列表并更新
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/4 9:58
// @success:
func GetAndUpdateRepo(workspaces string) error {
	if workspaces == "" {
		workspaces = utils.K8sConfig.K8s.Workspace.Name
	}
	res, err := GetRepoList(workspaces)
	if err != nil {
		return err
	}
	for _, v := range res.Items {
		_, err = UpdateRepo(workspaces, v.RepoId)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateWorkspaces
// @description: 创建企业空间
// @param: name string 要创建的企业空间名称
// @param: desc string 企业空间描述
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/8/30 20:16
// @success:l
func CreateWorkspaces(name, aliasName, desc string) error {
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/tenant.kubesphere.io/v1alpha2/workspaces"
	reqData := NewCreateWorkspacesRequest(name, aliasName, desc)

	httpCode, err := JsonRestRequestTimeout("POST", reqUrl, reqData, nil, time.Second*10)
	if err != nil || httpCode != http.StatusOK {
		log.WithFields(log.Fields{"err": err, "reqUrl": reqUrl, "res": nil}).Info("创建企业空间失败")
		return err
	}
	return nil
}

// CreateProject
// @description: 创建企业空间下的项目
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/1 10:41
// @success:
func CreateProject(name, workspace, aliasName, desc string) error {
	reqData := NewCreateProjectRequest(name, aliasName, desc, workspace)
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/tenant.kubesphere.io/v1alpha2/workspaces/" + workspace + "/namespaces"

	httpCode, err := JsonRestRequestTimeout("POST", reqUrl, reqData, nil, time.Second*10)
	if err != nil || httpCode != http.StatusOK {
		log.WithFields(log.Fields{"httpCode": httpCode, "err": err, "reqUrl": reqUrl, "res": nil}).Error("创建项目失败")
		return err
	}
	return nil
}

// CreateRepos
// @description: 添加应用仓库
// @param: workspace string 要添加到的企业空间
// @param: repoName string 应用仓库名称
// @param: projectName string 私有仓库中的项目名称
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/1 18:03
// @success:
func CreateRepos(harborUrl, workspace, repoName, projectName string) error {
	if harborUrl == "" {
		harborUrl = "http://core.harbor.dked:30002"
	}
	//TODO 创建harbor中的项目
	reqData := NewCreateRepoRequest(harborUrl, repoName, projectName)
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/workspaces/" + workspace + "/repos"
	httpCode, err := JsonRestRequestTimeout("POST", reqUrl, reqData, nil, time.Second*10)
	if err != nil || httpCode != http.StatusOK {
		log.WithFields(log.Fields{"httpCode": httpCode, "err": err, "reqUrl": reqUrl, "res": nil}).Info("创建应用仓库失败")
		return err
	}
	return nil
}

// GetRepoApps
// @description: 获取应用仓库下的应用
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 11:39
// @success:
func GetRepoApps(repoId string) (res *ReposAppsResponse, err error) {
	//reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/apps?orderBy=create_time&paging=limit%3D12%2Cpage%3D1&conditions=status%3Dactive%2Crepo_id%3Drepo-95x39n214x9p1o&reverse=true"
	reqUrl := utils.K8sConfig.K8s.Url + "/kapis/openpitrix.io/v1/apps?orderBy=create_time&conditions=repo_id=" + repoId

	httpCode, err1 := JsonRequestTimeout("GET", reqUrl, nil, &res, time.Second*10)
	if err1 != nil || httpCode != http.StatusOK {
		log.WithField("err", err1).Info("获取应用仓库应用失败")
		return res, errors.New("获取应用仓库应用失败")
	}

	return
}

// CreateApp
// @description: 创建项目中的应用
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 14:39
// @success:
func CreateApp(workspace, namespace, appName string) error {
	fmt.Println("-----------", workspace, "---", namespace, "------", appName)
	//获取仓库列表
	repos, err := GetRepoList(workspace)
	if err != nil {
		log.WithField("err", err).Error("获取应用仓库列表失败")
		return err
	}
	if len(repos.Items) <= 0 {
		errs := "当前项目未添加应用仓库"
		log.WithField("err", errs).Error(errs)
		return errors.New(errs)
	}
	repoId := repos.Items[0].RepoId
	//更新应用仓库
	UpdateRepo(workspace, repoId)
	//获取应用仓库中的应用
	apps, err := GetRepoApps(repoId)
	if err != nil {
		log.WithField("err", err).Error("获取应用仓库中的应用失败")
		return err
	}
	if len(apps.Items) <= 0 {
		errs := "harbor中的该项目没有对应的helm应用"
		log.WithField("err", errs).Error(errs)
		return errors.New(errs)
	}

	appid := apps.Items[0].Appid
	versionId := apps.Items[0].LatestAppVersion.VersionId
	//获取version helm yaml内容
	conf, err := Files(appid, versionId)
	if err != nil {
		log.WithField("err", err).Error("获取应用对应的helm文件信息错误")
		return err
	}
	//开始创建应用，部署实际项目
	err = CreateProjectApp(workspace, namespace, appid, versionId, conf, appName)
	if err != nil {
		log.WithField("err", err).Error("创建实际应用出错")
		return err
	}
	return err
}

// CreateProjectApp
// @description: 创建开发的实际项目应用。企业空间-项目-应用
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 14:36
// @success:
func CreateProjectApp(workspace, namespace, appid, versionId, conf, name string) error {
	reqData := CreateProjectAppRequest{
		appid,
		conf,
		name,
		versionId,
	}
	reqUrl := utils.K8sConfig.K8s.Url + fmt.Sprintf("/kapis/openpitrix.io/v1/workspaces/%s/namespaces/%s/applications", workspace, namespace)
	httpCode, err := JsonRestRequestTimeout("POST", reqUrl, reqData, nil, time.Second*10)
	if err != nil || httpCode != http.StatusOK {
		log.WithFields(log.Fields{"httpCode": httpCode, "err": err, "reqUrl": reqUrl, "res": nil}).Info("创建项目失败")
		return err
	}
	return nil
}

type Create struct {
	Workspace          string
	Namespace          string
	RepoName           string
	RepoProjectName    string
	WorkspaceAliasName string
	WorkspaceDesc      string
	ProjectAliasName   string
	ProjectDesc        string
	AppName            string
	HarborUrl          string
}

// CreateAll
// @description: 创建企业空间，项目，应用仓库，应用
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/9/2 18:43
// @success:
func CreateAll(c *Create) (err error) {
	err = CreateWorkspaces(c.Workspace, c.WorkspaceAliasName, c.WorkspaceDesc)
	if err != nil {
		log.WithField("err", err).Error("创建企业空间失败")
		//return
	}
	time.Sleep(time.Second)
	err = CreateProject(c.Namespace, c.Workspace, c.ProjectAliasName, c.ProjectDesc)
	if err != nil {
		log.WithField("err", err).Error("创建企业空间中的项目")
		//return
	}
	time.Sleep(time.Second)
	err = CreateRepos(c.HarborUrl, c.Workspace, c.RepoName, c.RepoProjectName)
	if err != nil {
		log.WithField("err", err).Error("创建企业空间中的应用仓库失败")
		return
	}
	time.Sleep(time.Second)
	err = CreateApp(c.Workspace, c.Namespace, c.AppName)
	if err != nil {
		log.WithField("err", err).Error("创建企业空间中的应用失败")
		return
	}
	return
}
