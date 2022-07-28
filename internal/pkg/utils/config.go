package utils

import (
	configor "github.com/jinzhu/configor"
)

/*
	存储全局参数，供其他模块使用
*/
var Config = struct {
	Log struct {
		Output string `default:"std"`       //日志输出，标准输出或文件
		Level  string `default:"InfoLevel"` //日志等级
		Caller bool   `default:"true"`      //是否打印调用者信息
		Dir    string `default:"."`         //存放目录
	}
	Web struct {
		Port string `default:"9008"`
		Cors bool   `default:"true"`
	}
	Path   string `default:"/tmp"`
	Mysql struct {
		Host     string `default:"127.0.0.1"`
		UserName string `default:"root"`
		Password string `default:"root"`
		DBName   string `default:"mck"`
		Port     string `default:"3306"`
		MinConns int    `default:"90"`  //连接池最小连接数量 不要太小
		MaxConns int    `default:"120"` //连接池最大连接数量 两者相差不要太大
	}
}{}

var K8sConfig = struct {
	K8s struct {
		//Url       string `default:"http://ks-apiserver.kubesphere-system.svc"`
		Url        string `default:"http://192.168.0.80:30532"`
		Namespace  string `default:"test"`
		Workspace string `default:"dked"`
		Username   string `default:"csmp"`
		Password   string `default:"Dked@213"`
	}
	Harbor struct {
		//Address  string `default:"dockerhub.dked.local:30002"`
		Address  string `default:"192.168.0.80:30002"`
		Admin    string `default:"admin"`
		Password string `default:"Harbor12345"`
		Project string `default:"csmp"`
	}
}{}

//读取用户的配置文件
func InitConfig() {
	err := configor.Load(&Config, "./config/config.yml")
	if err != nil {
		panic("config load error" + err.Error())
	}
	err = configor.Load(&K8sConfig, "/config/csmp-k8s.yml")
}
