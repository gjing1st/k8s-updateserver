#
## 🌱1基本介绍
### 1.1 项目目的
业务部署在kubesphere中，为方便业务系统升级和业务系统部署。使用go语言开发此升级服务和saas化接口
### 1.2 项目简介
基于以上目的，升级服务和k8s部署工具单独打包

## 2 使用说明
- 升级服务可通过Dockerfile打包成镜像，随业务系统一起部署至k8s中
- 部署工具至cmd/kubetool/kubetool.go打包
- 针对首次部署k8s集群，增加kubetool功能，部署完ks后，可以将所有需要的镜像以及helm打包成一个zip压缩包，使用kubetool上传镜像至harbor私有仓库以及部署应用

# Thanks
<a href="https://www.jetbrains.com/?from=k8s-updateserver"><img src="https://gitee.com/gjing1st/images/raw/master/JetBrains.png" width="100" alt="JetBrains"/></a>