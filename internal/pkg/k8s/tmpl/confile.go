// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/2$ 23:29$

package tmpl

var ConfigFile = `
k8s:
  #企业空间名称
  workspace: 
    name: dked
    #别名
    aliasname: dked
    desc: 笛卡尔盾
  #企业空间中的项目名称对应k8s命名空间
  namespace: 
    name: csmp
    aliasname: dked-csmp
    desc: 笛卡尔盾x-csmp项目
  #应用仓库
  repo: 
    name: harbor
    #harbor私有仓库中的项目名称
    projectname: library
  #要部署的应用名称
  appname: csmp
  mysql:
    database: mmyypt_db
    password: zf12345678

harbor:
  address: core.harbor.dked:30002
  admin: admin
  password: Harbor12345
  project: csmp
`
