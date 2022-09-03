// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/8/31$ 15:27$

package tmpl

//MysqlServer k8s中的服务和工作负载
var MysqlServer = `apiVersion: apps/v1
kind: StatefulSet
metadata:
  namespace: {{projectName}}
  labels:
    version: v1
    app: mysql
  name: mysql-v1
spec:
  replicas: 1
  selector:
    matchLabels:
      version: v1
      app: mysql
  template:
    metadata:
      labels:
        version: v1
        app: mysql
      annotations:
        logging.kubesphere.io/logsidecar-config: '{}'
    spec:
      containers:
        - name: container-58s792
          imagePullPolicy: IfNotPresent
          image: {{image}}
          ports:
            - name: tcp-3306
              protocol: TCP
              containerPort: 3306
              servicePort: 3306
            - name: tcp-33060
              protocol: TCP
              containerPort: 33060
              servicePort: 33060
          env:
            - name: MYSQL_DATABASE
              valueFrom:
                secretKeyRef:
                  name: csmp-mysql-secret
                  key: MYSQL_DATABASE
            - name: MYSQL_ROOT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: csmp-mysql-secret
                  key: MYSQL_ROOT_PASSWORD
          volumeMounts:
            - name: host-time
              mountPath: /etc/localtime
              readOnly: true
            - name: volume-cujdzd
              readOnly: true
              mountPath: /etc/mysql/conf.d
            - name: volume-ec7lm2
              readOnly: false
              mountPath: /var/lib/mysql
      serviceAccount: default
      initContainers: []
      volumes:
        - hostPath:
            path: /etc/localtime
            type: ''
          name: host-time
        - name: volume-cujdzd
          configMap:
            name: csmp-mysql-conf
            items:
              - key: my.cnf
                path: my.cnf
        - name: volume-ec7lm2
          persistentVolumeClaim:
            claimName: csmp-mysql-pvc
      imagePullSecrets: null
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      partition: 0
  serviceName: mysql
---
apiVersion: v1
kind: Service
metadata:
  namespace: {{projectName}}
  labels:
    version: v1
    app: mysql
  annotations:
    kubesphere.io/serviceType: statefulservice
    kubesphere.io/description: csmp数据库
  name: mysql
spec:
  sessionAffinity: None
  selector:
    app: mysql
  ports:
    - name: tcp-3306
      protocol: TCP
      port: 3306
      targetPort: 3306
    - name: tcp-33060
      protocol: TCP
      port: 33060
      targetPort: 33060
  clusterIP: None
`

//MysqlSecret mysql保密字典
var MysqlSecret = `kind: Secret
apiVersion: v1
metadata:
  name: {{projectName}}-mysql-secret
  namespace: {{projectName}}
  annotations:
    kubesphere.io/creator: admin
    kubesphere.io/description: csmp数据库保密字典
data:
  MYSQL_DATABASE: {{database}}
  MYSQL_ROOT_PASSWORD: {{password}}
type: Opaque
`

//MysqlConf mysql配置信息
var MysqlConf = `kind: ConfigMap
apiVersion: v1
metadata:
  name: {{projectName}}-mysql-conf
  namespace: {{projectName}}
  annotations:
    kubesphere.io/creator: admin
    kubesphere.io/description: csmp数据库配置
data:
  my.cnf: |+
    [client]
    default-character-set=utf8mb4

    [mysql]
    default-character-set=utf8mb4

    [mysqld]
    init_connect='SET collation_connection = utf8mb4_unicode_ci'
    init_connect='SET NAMES utf8mb4'
    character-set-server=utf8mb4
    collation-server=utf8mb4_unicode_ci
    sql_mode="NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"

    [client]
    default-character-set = utf8mb4
    [mysql]
    default-character-set = utf8mb4
    !includedir /etc/my.cnf.d
    [mysqld]
    server-id=1
    log-bin=mysql-bin
    #log-bin-index=master-bin.index
    binlog-ignore-db=mysql
    relay-log=mysql-relay
    #datadir=/var/lib/mysql
    character_set_server = utf8mb4
    socket=/var/lib/mysql/mysql.sock
    sql_mode="NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"
    # Disabling symbolic-links is recommended to prevent assorted security risks
    symbolic-links=0
    skip-name-resolve
`
