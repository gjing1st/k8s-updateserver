package apiserver

import v1 "upserver/internal/apiserver/router/v1"

func HttpStart() {
	run()
}

// @description: 启动http服务
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/12 18:56
// @success:
func run() {
	v1.InitApi()
}
