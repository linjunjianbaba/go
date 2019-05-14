

学习博客：http://blog.itpub.net/28624388/viewspace-2153546/

kubernetes（库巴乃梯丝）​         舵手，飞行员

scheduler					调度器

controller manager                       控制器

selector                                          选择器（标签）

kubelet						接受任务

Pod 

label						标签

label selector                                标签选择器

containers					容器

Pod：自主式pod

​            控制器管理的pod

​                     ReplicationController    副本控制器

​                      Replicaset                       副本集控制器

​                      Deployment

​                      StatefulSet

​                      DaemonSet

​                      Job，Cronjob

HPA：

​       HorizontalPodAutoscaler



service

​          iptables 规则

AddOns：附件

Overlay Network：叠加网络

CNI插件

​        flannel：网络配置

​         calico：网络配置，网络策略bgp协议路由直通

​         canel:网络策略

![1544160316790](C:\Users\bill\AppData\Roaming\Typora\typora-user-images\1544160316790.png)

​          kubeadm

​                  1.master，nodes：安装kubelet，kubeamd，docker

​                  2.master：kubeamd  init

​                   3.nodes:

docker info

docker image inspect  镜像名称

1.配置docker Unit Fil中的Environment变量，定义其HTTPS_PROXY,或者先导入所需要的镜像文件；

Environment="HTTPS_PROXY=http://www.ik8s.io:10080"

Envrionment="NO_PROXY=127.0.0.0/8,172.16.0.0/16"

2.编辑kubelet的配置文件/etc/sysconfig/kubelet，设置其忽略Swap启用的状态错误，内容如下：

KUBELET_EXTRA_ARGS="--fail-swap-on=false"

KUBE_PROXY_MODE=ipvs

3.设定docker kubelet开机启动

systemclt enable docker kubelet

4.初始化master节点：

  kubeadm init --kubernetes-version=v1.13.1 --pod-network-cidr=10.244.0.0/16

--service-cidr=10.96.0.0/12 --ignore-predlight-errors=Swap

5.初始化kubectl

mkdir ~/.kube

cp /etc/kubernetes/admin.conf ~/.kube/

echo "source <(kubectl completion bash)" >> ~/.bashrc

6.添加flannel网络附件

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

7.验证









kubeadm --kubernetes-version= --pod-network-cidr= --service-cidr=

vim /etc/sysconfig/kubelet        KUBELET_EXTRA_ARGS='"--fail-swap-on=flase"

kubeadm config images pull

kubectl run nginx-deploy --image=nginx:1.14-alpine --replicas=1

kubectl expose deployment nginx-deploy --name=nginx

rpm --import

kubectl get cs

kubectl get ns

kubect get svc

kubectl get nodes

kubectl get pods -n kube-system -o wide

 kubectl get deployment

kubectl get pods

kubectl describe svc nginx

kubectl get deployment -w

kubectl get pods -o wide  

kubectl scale --replicas=4 deployment nginx-deploy    pod扩容

kubectl delete pods    删除pod

kubectl set image deployment nginx-deploy nginx-deploy=nginx  滚动升级

 kubectl rollout status deployment nginx-deploy            查看滚动升级

kubectl get pod *** -o yaml

kubectl api-versions

kubectl get pods --show-labels

RESTful

​          GET,PUT,DELETE,POST.....

​          kubectl run ，get ，edit

资源：对象

​     workload：Pod, ReplicaSet,Deploument,StatefulSet,DaemonSet,Job,Cronjob...

​     服务发现及均衡：Service，Ingress

​     配置与存储：Volume，CSI

​            ConfigMap，Secret

​             DownwardAPI

​      集群资源：

​          Namespace，Node，Role，ClusterRole，RoleBinding，ClusterRoleBinding

​       元数据型资源：

​             HPA,PodTemplate，LimitRange





​    





资源创建的方法：

​           apiversion仅接受JSON格式的资源定义

​            yaml格式提供配置清单，apiservice可自动将其转为JSON格式，而后再提交

大部分资源的配置清单：

​      apiVersion：group/version

​              $kubectl api-versions

​       kind：资源类别

​       metadata：元数据

​              name：

​              namespace：

​              labels：

​               annotations      ：资源注解

​             每个资源的 应用API：/api/GROUP/VERSION/namespace/NAMESPACE/TYPE/NAME

​      spec：期望状态，disired state

​       status：当前状态，本字段由集群维护

​               $kubuctl explain pods.spec.containers.images             yaml文件帮助

​                 kubectl exec -it pod -- /bin/bash



资源配置清单：

​       自主式Pod资源

​       资源清单格式:

​               一级字段：apiVersion（group/version）kind，metadata（name，namespace，labels，annotaions），spec（containers），status（只读）

​        pod资源：12

​               spec.containers

​                   name

​                   image

​                   imagaPullPolicy

​                            Always,Never,IfNotPresent,

​                   ports：

​                                        



​       kubectl get pods -L app

​        kubectl get pods -l app --show-labels  标签选择

​       kubectl labels pods  pods-name labelsname=zhi  --overwrite（已存在修改标签）

​                                 类型    名称             标签名=标签值

​       kubectl get pods -l 标签名，标签名= 

Pod的生命周期

​        状态：Pending（挂起），Running（运行），Failed（失败），succeeded，Unknow

​         创建Pod:

​        pod生命周期中的重要行为：

​            初实化容器

​             容器探测：

​                liveness

​                 readiness

service：

   工作模式： 

​        user space：1.1-

​         IP tables： 1.10-

​         IP VS： 1.11+

   类型：

​       ExternalName，ClusterIP，NodePort，and LoadBalancer

   资源记录：

​          SVC_NAME.NS_NAME.DOMAIN.LTD.

​           redis.default.svc.cluster.local.