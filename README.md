#
# 🌱基本介绍
### 1 项目目的
为方便部署在k8s-kubesphere上的应用在自己的管理后台进行升级，特开发此系统。

## 2 项目简介
业务系统第一次通过kubesphere部署应用。以后升级各个应用服务只需将所有的镜像以及应用helm包打成一个大的压缩包，通过自己管理平台上传压缩包。
升级服务进行解压缩并遍历所有的镜像tar包和helm包上传至指定仓库。
业务系统根据实际项目和应用请求升级服务接口获取应用版本信息，并升级到指定版本。
此服务使用户只需关心自己业务系统即可，不需去kubesphere平台进行操作。
## 3 使用说明
```bash
git clone https://github.com/gjing1st/k8s-updateserver.git
```

## 4 限制
```bash
golang版本 >= 1.17
```

## 4. 捐赠
如果你觉得这个项目对你有帮助，你可以请作者喝饮料 :
<div align="center">
<table><tr>
<td><img src=https://gitee.com/gjing1st/bill_admin/raw/master/billApi/public/image/alipay.jpg border=0></td>
<td><img src=https://gitee.com/gjing1st/bill_admin/raw/master/billApi/public/image/wechat.jpg border=0></td>
</tr></table>
</div>
