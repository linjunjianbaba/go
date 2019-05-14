k8s认证

令牌认证：token，有没有预共享令牌

证书认证：

RBAC:许可授权

k8s用户账号

客户端-->API Server

​       user: username ,uid

​       group:

​       extra:

​       API:

​         Request path

​               /apis/apps/v1/namespaces/default

​        请求动作：

​                   get post put delete

​       API 请求动作：

​                 get list create update patch watch proxy redirect delete deletecollection

Resource：

 Subresource：

Namespace：

Api group：

每个命名空间都有一个default-token



```shell
kubectl create serviceaccount admin    #创建admin sa
```

```shell
kubectl config view                   #查看配置
```

  K8s创建自己的账号：

​                  生成私钥：

```shell
(umask 077; openssl genrsa -out bill.key 2048)  #生成私钥
openssl req -new -key bill.key -out bill.csr -subj "/CN=bill"                                   #生成证书签署请求
openssl x509 -req -in bill.csr -CA ./ca.crt -CAkey ./ca.key -CAcreateserial -out bill.crt -days 3650                #证书签署
openssl x509 -req -in bill.csr -CA ./ca.crt -CAkey ./ca.key -CAcreateserial -out bill.crt -days 3650

openssl x509 -in bill.crt -text -noout #查看证书签署
kubectl config set-credentials bill --client-certificate=./bill.crt --client-key=./bill.key --embed-certs=true   #添加用户验证
 kubectl config set-context bill@kubernetes --cluster=kubernetes --user=bill #配置上下文
 kubectl config use-context bill@kubernetes  #使用指定用户运行kubectl
```

RBAC权限控制：

授权插件：Node, ABAC,RBAC,Webhook（http回调）

   角色（role）

   许可（permission）

#### jenkinsfile:  https://blog.csdn.net/zbbkeepgoing/article/details/83098023

kubernetes:认证，授权

  API server：

 认证：token，tls，user/password

​           账号：user Account，ServiceAccount

​           授权：RBAC

​              role，rolebinding

​              clusterrole，clusterrolebinding

​           subject：

​                   user

​                    group

​                    serviceaccount

​              object：

​                 resouce group

​                 resouce 

​                 non-resouce url

​              actuon:  get,list,watch,patch,delete,deletecollection,...

​         subject

​       subject（主题）

```yaml
kubectl create role pods-reader --verb=get,list,watch --resource=pods --dry-run -o yaml

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: pods-reader
rules:
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - watch
  
kubectl create rolebinding bill-pods --role=pods-reader --user=bill --dry-run -o yaml

kubectl create clusterrolebinding bill-pods --clusterrole=cluster-pods --user=bill -dry-run -o yaml

kubectl create clusterrolebinding bill-pods --clusterrole=cluster-pods --user=bill --dry-run -o yaml #集群绑定
```

角色绑定相关应用：

Dashboard:

​    1.部署

​    2.将Service修改为Node

​    3.认证：

​        认证的账号必须为ServiceAccount：被dashboard pod拿来由kubernetes进行验证

​      token:

​                (1) 创建ServiceAccount，根据管理目标，使用rolebinding或clusterbinding绑定至合理的role  clusterrole

​                 (2)获取ServiceAccount的secret，查看secret的详细信息，其中就由token

​      kubeconfig：把ServiceAccount的token封装为kubeconfig文件

​                （1）创建ServiceAccount根据管理目标，使用rolebinding或clusterbinding绑定至合理的role  clusterrole

​                  （2）kubectl get secret |awk '/^SerciceAccount/ {print $1}'

​                           KUBE_TOKEN=$(kubectl get secret )

​                     (3)生成kubeconfig文件

​                                kubectl config set-cluster

​                                kubectl config set-credentials

​                                kubectl config set-context

​                                kubectl config use-context

单个命名空间权限绑定角色

kubectl create rolebinding k8sadmin --clusterrole=cluster-admin --serviceaccount=kube-system:k8sadmin   

 全部命名空间权限需要绑定集群角色

kubectl create clusterrolebinding k8sadmin --clusterrole=cluster-admin --serviceaccount=kube-system:k8sadmin       

​       生成专用证书：

```shell
#(umask 077; openssl genrsa -out dashboard.key 2048)
#openssl req -new -key dashboard.key -out dashboard.csr -subj "/O=bill/CN=ui.k8sz.com" 
#openssl x509 -req -in dashboard.csr -CA ca.crt -CAkey ca.key -CA createserial -out dashboard.crt -days 3650
# kubectl create secret generic dashboard-cert -n kube-system --from-file=./dashboard.crt --from-file=dashboard.key=./dashboard.key       #secret创建
#kubectl create secret generic dashboard-cert -n kube-system --from-file=./bill.crt --from-file=./bill.key --dry-run -o yaml      #secret创建
#kubectl create serviceaccount k8sadmin -n default  #创建serviceaccount
# kubectl create rolebinding k8sadmin --clusterrole=admin --serviceaccount=default:k8sadmin  #serviceaccount账户绑定到集群角色admin

```

kubeconfig认证:

```yaml
#kubectl config set-cluster kubernetes --certificate-authority=/etc/kubernetes/pki/ca.crt --server="https://192.168.139.134:6443" --embed-certs=true --kubeconfig=/root/k8sadmin.conf  #配置集群信息

```

