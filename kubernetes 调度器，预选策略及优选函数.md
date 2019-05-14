kubernetes 调度器，预选策略及优选函数

调度器

​    预选策略：

​           checkNodeCondition：检查节点第二个位置

​           GeneralPredicates

​                 HostName： 检查Pod对象是否定义了pod.spec.hostname

​                  PodFitsHostPorts: pods.spec.containers.ports.hostPort

​                  MatchNodeSelector: pods.spec.nodeSelector

​                  podFitsResources: 检查Pod资源需求是否能被节点所满足

NodeDiskConflict： 检查Pod依赖的存储卷是否满足需求

PodToleratesNodeTaints:



优先函数

   Least Requested：

   BalancedResourceAllocation：CPU和内存资源被占有率的胜出

   NodePreferAvidPods：节点注解

   TanitToleration：将Pod对象的sepc.tolerations列表与节点的taints列表项进行匹配度检查，比配

SelectorSpreading：

InterPodAffinity：

NodeLabel：

Image



```
apiVersion: v1
kind: Pod
metadata:
  name: pod-s
  labels:
    app: pod-s
spec:
  containers:
  - name: pod-s
    image: ikubernetes/myapp:v1
    imagePullPolicy: IfNotPresent
  nodeSelector:
    kubernetes.io/hostname: master01 
```

kubectl get pods --shwo-labels

节点亲和性

现实中应用的运行对于kubernetes在亲和性上提出了一些要求，可以归类到以下几个方面： 
1.Pod固定调度到某些节点之上 
2.Pod不会调度到某些节点之上 
3.Pod的多副本调度到相同的节点之上 
4.Pod的多副本调度到不同的节点之上

#### Pod调动到某些节点上

Pod的定义中通过nodeSelector指定label标签，pod将会只调度到具有该标签的node之上

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  nodeSelector:
    disktype: ssd
```

Pod间的亲和性和反亲和性

基于已经运行在Node 上pod的labels来决定需要新创建的Pods是否可以调度到node节点上，配置的时候可以指定那个namespace中的pod需要满足pod的亲和性．可以通过topologyKey来指定topology domain, 可以指定为node／cloud provider zone／cloud provider region的范围

表达的语法：支持In, NotIn, Exists, DoesNotExist

Pod的亲和性和反亲和性可以分成
requiredDuringSchedulingIgnoredDuringExecution　#硬要求
preferredDuringSchedulingIgnoredDuringExecution　＃软要求

类似上面node的亲和策略类似，requiredDuringSchedulingIgnoredDuringExecution亲和性可以用于约束不同服务的pod在同一个topology domain的Nod上．preferredDuringSchedulingIgnoredDuringExecution反亲和性可以将服务的pod分散到不同的topology domain的Node上．

topologyKey可以设置成如下几种类型
kubernetes.io/hostname　　＃Node
failure-domain.beta.kubernetes.io/zone　＃Zone
failure-domain.beta.kubernetes.io/region #Region

可以设置node上的label的值来表示node的name,zone,region等信息，pod的规则中指定topologykey的值表示指定topology范围内的node上运行的pod满足指定规则

亲和性：https://blog.csdn.net/jettery/article/details/79003562











资源限制：

```
apiVersion: v1
kind: Pod
metadata:
  name: 
  labels:
spec:
  containers:
  - name:
    image:
    imagePullPolicy:
    resources:
      requests:
        cpu: "200m"
        memory: "128Mi"
      limits:
        cpu: "500m"
        memory: "200Mi"
```

(umask 077; openssl genrsa -out serving.key 2048)



 