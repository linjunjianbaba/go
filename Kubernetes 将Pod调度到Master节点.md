# Kubernetes 将Pod调度到Master节点

出于安全考虑，默认配置下Kubernetes不会将Pod调度到Master节点。如果希望将k8s-master也当作Node使用，可以执行如下命令：

```bash
kubectl taint node master01 node-role.kubernetes.io/master-
```

其中k8s-master是主机节点hostname如果要恢复Master Only状态，执行如下命令：

```bash
kubectl taint nodes master1 node-role.kubernetes.io/master=:NoSchedule
```

```bash
kubectl taint node [node] key=value[effect]   
     其中[effect] 可取值: [ NoSchedule | PreferNoSchedule | NoExecute ]
      NoSchedule: 一定不能被调度
      PreferNoSchedule: 尽量不要调度
      NoExecute: 不仅不会调度, 还会驱逐Node上已有的Pod
```

## 容忍tolerations主节点的taints

```yaml
tolerations:
- key: "node-role.kubernetes.io/master"
  operator: "Equal"
  value: ""
  effect: "NoSchedule"
```



## Kubernetes1.13.1部署Kuberneted-dashboard v1.10.1

下载镜像并修改镜像标签

docker pull registry.cn-beijing.aliyuncs.com/minminmsn/kubernetes-dashboard:v1.10.1

docker tag registry.cn-beijing.aliyuncs.com/minminmsn/kubernetes-dashboard:v1.10.1 k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1

下载官方yaml文件

wget https://raw.githubusercontent.com/kubernetes/dashboard/master/aio/deploy/recommended/kubernetes-dashboard.yaml

修改kubernetes-dashboard.yaml

```yaml
kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kube-system
spec:
  type: NodePort    #使用节点模式
  ports:
    - port: 443
      targetPort: 8443
      nodePort: 30001   #绑定节点端口
  selector:
    k8s-app: kubernetes-dashboard
```

新建登陆UI账户，vi ui.yaml

```yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: admin
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: ServiceAccount
  name: admin
  namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin
  namespace: kube-system
  labels:
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
```

获取UI登陆token

```
kubectl apply -f admin-token.yaml 
kubectl describe secret/$(kubectl get secret -nkube-system |grep admin|awk '{print $1}') -n kube-system
```

Kubernetes1.13.1部署ingress-nginx控制器

下载ingress-nginx控制器yaml文件

wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml

建立服务ingress-nginx.yaml文件

```yaml
kind: Service
apiVersion: v1
metadata:
  name: ingress-nginx
  namespace: ingress-nginx
  labels:
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
spec:
  type: NodePort
  selector:
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
  ports:
    - name: http
      port: 80
      targetPort: http
      protocol: TCP
      nodePort: 30080
    - name: https
      port: 443
      targetPort: https
      protocol: TCP
      nodePort: 30443

---
```

建立ingress.yaml

```yaml
apiVersion: extensions/v1beta1      #api版本
kind: Ingress       #清单类型
metadata:           #元数据
  name: ingress-myapp    #ingress的名称
  namespace: default     #所属名称空间
  annotations:           #注解信息
    kubernetes.io/ingress.class: "nginx"
spec:      #规格
  rules:   #定义后端转发的规则
  - host: myapp.k8sz.com    #通过域名进行转发
    http:
      paths:
      - path:       #配置访问路径，如果通过url进行转发，需要修改；空默认为访问的路径为"/"
        backend:    #配置后端服务
          serviceName: myapp
          servicePort: 80
```



