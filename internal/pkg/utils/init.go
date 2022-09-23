package utils

func init() {
	//InitConfig()
	InitLogger(Config.Log.Level, Config.Log.Output, Config.Log.Dir, Config.Log.Caller)
}
