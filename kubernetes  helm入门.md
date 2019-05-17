kubernetes  helm入门

chart：是Helm管理的安装包，里面包含需要部署的安装包资源。类似于yum中的rpm文件。每个Chart包含下面两部分：包的基本描述文件Chart.yaml放在templates目录中的一个或多个Kubernetes manifest文件模板。

**Release:**在Kubernetes集群上运行的一个Chart实例。在同一个集群上，一个Chart可以安装多次。例如一个MySQL
Chart，如果想在服务器上运行两个MySQL数据库，就可以基本这个Chart安装两次。每次安装都会生成新的Release,会有独立的Release名称。

**Repository:** 用于存放和共享Chart的仓库。

helm（客户端）-->tiller（服务端）

chart：一个helm程序包

Repository：Chart仓库，http/https服务器

Release：特定Chart部署于目标集群上的一个实例

Chart—>config-->Release

git:https://github.com/helm/charts/blob/master/stable

程序架构：

​         helm:客户端，管理本地的Chart仓库，管理chart，于Tiller服务器交互，发送Chart，实例安装，查询，卸载等操作

​         Tiller：服务端（k8s node上安装），接收helm发送过来的

下载：wget https://storage.googleapis.com/kubernetes-helm/helm-v2.13.0-rc.1-linux-amd64.tar.gz

tar xf helm-v2.13.0-rc.1-linux-amd64.tar.gz

tiller:安装：https://github.com/helm/helm/blob/master/docs/rbac.md

vim rbac-config.yaml

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: tiller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: tiller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: tiller
    namespace: kube-system
```

```shell
helm init --upgrade --service-account tiller --tiller-image registry.cn-hangzhou.aliyuncs.com/google_containers/tiller:v2.14.0  #初始化 Helm 并安装 Tiller 服务
helm version   #查看版本信息
helm repo update #更新仓库到本地
helm search    #查看helm仓库
helm inspect stable/jenkins  #查看chart详细信息
helm install --name mem1 stable/tomcat #部署安装
helm upgrade --set mysqlRootPassword=passwd db-mysql stable/mysql  #升级
helm rollback db-mysql 1 #回滚
helm listh	
helm ls #查看Release列表
helm delete #删除Release
helm upgrade
helm rollback
helm delete --purge redis
helm install --name els --namespace=efk -f values.yaml incubator/elasticsearch
cd /root/.helm/cache/archive
# 添加 incubator repo
helm repo add incubator https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts-incubator/
# 查询 repo 列表
helm repo list
# 生成 repo 索引（用于搭建 helm repository）
helm repo index
# 创建一个新的 chart
helm create hello-chart
# validate chart
helm lint
# 打包 chart 到 tgz
helm package hello-chart

#生成yaml
helm template install/kubernetes/helm/istio-init --name istio-init --namespace istio-system > istio.yaml
```

可用helm仓库：https://hub.helm.sh/   https://hub.kubeapps.com

helm常用命令

​    re

创建自定义chart

reids数据同步:https://github.com/vipshop/redis-migrate-tool