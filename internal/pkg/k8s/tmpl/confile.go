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
  url: http://192.168.0.80:31601
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
    password: 123456
    image: core.harbor.dked:30002/csmp/mysql:5.7.35

harbor:
  address: http://core.harbor.dked:30002
  admin: admin
  password: Harbor12345
  project: csmp
`
