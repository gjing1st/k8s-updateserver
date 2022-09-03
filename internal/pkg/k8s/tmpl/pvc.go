// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/2$ 16:56$

package tmpl

//ConfPvc config pvc
var ConfPvc = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  namespace: {{projectName}}
  name: {{projectName}}-conf-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
  storageClassName: rook-cephfs
`
//FrontendPvc 前端nginx存储卷
var FrontendPvc = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  namespace: {{projectName}}
  name: {{projectName}}-frontend-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
  storageClassName: rook-cephfs
`
//KmcPvc kmc存储卷
var KmcPvc = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  namespace: {{projectName}}
  name: {{projectName}}-kmc-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
  storageClassName: rook-cephfs
`
//LibPvc lib存储卷
var LibPvc = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  namespace: {{projectName}}
  name: {{projectName}}-lib-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
  storageClassName: rook-cephfs
`
//MysqlPvc mysql存储卷
var MysqlPvc = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{projectName}}-mysql-pvc
  namespace: {{projectName}}
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 500Gi
  storageClassName: rook-cephfs
`
//UpdatePvc 升级服务pvc
var UpdatePvc = `
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{projectName}}-up-pvc
  namespace: {{projectName}}
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 10Gi
  storageClassName: rook-cephfs
`
