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
		Port string `default:"9680"`
		Cors bool   `default:"true"`
	}
	K8s struct {
		Namespace struct {
			Name      string `default:"csmp"`
			Aliasname string `default:"csmp-space"`
			Desc      string `default:"命名空间"`
		}
		Mysql   struct {
			Database string `default:"mmyypt_db"`
			Password string `default:"123456"`
			Image    string `default:"mysql:5.7.35"`
		}
		ElasticSearch struct {
			Address string `default:"http://192.168.0.80:31199"`
			//Address string `default:"elasticsearch-logging-data.kubesphere-logging-system.svc:9200"`
		}
		Statistic struct {
			CrontabTime   int    `default:"120"`
			MongoHost     string `default:"localhost:27017"`
			MongoDatabase string `default:"csmp"`
			Collection    string `default:"statistic"`
		}
	}
}{}


//读取用户的配置文件
func InitConfig() {
	err := configor.Load(&Config, "./config/config.yml")
	if err != nil {
		panic("config load error" + err.Error())
	}
}
