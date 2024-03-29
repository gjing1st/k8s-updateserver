package utils

import (
	"fmt"
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
		Port string `default:"9680"`
		Cors bool   `default:"true"`
	}
	Path                      string `default:"/tmp"`
	VersionInfoPath           string `default:"/home/versionInfo"`
	VersionInfoFlleNamePrefix string `default:"version_info"`
	Mysql                     struct {
		Host     string `default:"127.0.0.1"`
		UserName string `default:"root"`
		Password string `default:"root"`
		DBName   string `default:"es"`
		Port     string `default:"3306"`
		MinConns int    `default:"90"`  //连接池最小连接数量 不要太小
		MaxConns int    `default:"120"` //连接池最大连接数量 两者相差不要太大
	}
	IPM struct {
		Addr string `default:"http://114.115.134.131:9681"`
	}
	VersionInfo struct {
		Manufacturer string `default:"笛卡尔盾"`
		Serial       string `default:"1564835VAF348CF"`
		Algorithm    string `default:"SM2、SM3、SM4"`
	}
}{}

var K8sConfig = struct {
	K8s struct {
		Addr string `default:"http://ks-apiserver.kubesphere-system.svc"`
		//Namespace string `default:"csmp"`
		//Workspace string `default:"dked"`
		Username  string `default:"csmp"`
		Password  string `default:"Dked@213"`
		Workspace struct {
			Name      string `default:"dsjj"`
			Aliasname string `default:"dked-space"`
			Desc      string `default:"企业空间"`
		}
		Namespace struct {
			Name      string `default:"csmp"`
			Aliasname string `default:"csmp-space"`
			Desc      string `default:"命名空间"`
		}
		Repo struct {
			Name        string `default:"harbor-repo"`
			Projectname string `default:"library"`
		}
		Appname string `default:"csmp"`
		Mysql   struct {
			Database string `default:"mmyypt_db"`
			Password string `default:"123456"`
			Image    string `default:"mysql:5.7.35"`
		}
	}
	Harbor struct {
		Address string `default:"core.harbor.dked:30002"`
		//Address  string `default:"http://192.168.0.80:30002"`
		Admin    string `default:"admin"`
		Password string `default:"Harbor12345"`
		Project  string `default:"csmp"`
	}
	Backend struct {
		Addr string `default:"http://csmp-backend.csmp:8100"`
	}
}{}

var KubeToolConfig = struct {
	K8s struct {
		Workspace struct {
			Name      string `default:"dked"`
			Aliasname string `default:"dked-space"`
			Desc      string `default:"企业空间"`
		}
		//Url string `default:"http://ks-apiserver.kubesphere-system.svc"`
		Url       string `default:"http://192.168.0.80:31177"`
		Namespace struct {
			Name      string `default:"csmp"`
			Aliasname string `default:"csmp-space"`
			Desc      string `default:"命名空间"`
		}
		Repo struct {
			Name        string `default:"harbor-repo"`
			Projectname string `default:"library"`
		}
		Appname string `default:"csmp"`
		Mysql   struct {
			Database string `default:"mmyypt_db"`
			Password string `default:"123456"`
		}
	}
	Harbor struct {
		//Address  string `default:"core.harbor.dked:30002"`
		Address  string `default:"http://192.168.0.80:30002"`
		Admin    string `default:"admin"`
		Password string `default:"Harbor12345"`
		Project  string `default:"csmp"`
	}
}{}

// 读取用户的配置文件
func InitConfig() {
	err := configor.Load(&Config, "./config/config.yml")
	if err != nil {
		panic("config load error" + err.Error())
	}
	err = configor.Load(&K8sConfig, "./config/k8s.yml")
	fmt.Println(K8sConfig)
}
