package k8s

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
type Result struct {
	Code int
	Err  error
}

//AppListResponse k8s应用列表
type AppListResponse struct {
	Items []struct {
		Name    string `json:"name"`
		Cluster struct {
			AppId      string `json:"app_id"`
			ClusterId  string `json:"cluster_id"`
			CreateTime string `json:"create_time"`
			Env        string `json:"env"`
			Name       string `json:"name"`
			Owner      string `json:"owner"`
			RuntimeId  string `json:"runtime_id"`
			Status     string `json:"status"`
			StatusTime string `json:"status_time"`
			VersionId  string `json:"version_id"`
			Zone       string `json:"zone"`
		} `json:"cluster"`
		Version struct {
			AppId     string `json:"app_id"`
			Name      string `json:"name"`
			VersionId string `json:"version_id"`
		} `json:"version"`
		App struct {
			AppId       string `json:"app_id"`
			CategorySet string `json:"category_set"`
			ChartName   string `json:"chart_name"`
			Name        string `json:"name"`
		} `json:"app"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

//VersionResponse 版本信息
type VersionResponse struct {
	Items []struct {
		Active      bool   `json:"active"`
		AppId       string `json:"app_id"`
		CreateTime  string `json:"create_time"`
		Name        string `json:"name"`
		Owner       string `json:"owner"`
		PackageName string `json:"package_name"`
		Status      string `json:"status"`
		StatusTime  string `json:"status_time"`
		UpdateTime  string `json:"update_time"`
		VersionId   string `json:"version_id"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

type FilesResponse struct {
	Files struct {
		Helmignore string `json:".helmignore"`
		ValuesYaml string `json:"values.yaml"`
	} `json:"files"`
	VersionId string `json:"version_id"`
}

type UpVersionReq struct {
	AppId     string `json:"app_id"`
	ClusterId string `json:"cluster_id"`
	Cluster   string `json:"cluster"`
	Conf      string `json:"conf"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Owner     string `json:"owner"`
	VersionId string `json:"version_id"`
	Workspace string `json:"workspace"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type K8sAppAndVersion struct {
	App     *AppListResponse
	Version *VersionResponse
}

//AppRepoResponse 应用列表
type AppRepoResponse struct {
	Items []struct {
		RepoId string `json:"repo_id"`
		Url    string `json:"url"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

type UpdateRequest struct {
	Action string `json:"action"`
}
