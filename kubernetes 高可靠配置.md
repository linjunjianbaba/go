kubernetes 高可靠配置

为了保证应用可以稳定可靠的运行在Kubernetes里，本文介绍构建Kubernetes集群时的推荐配置。

## 磁盘类型及大小

**磁盘类型**

- 推荐选择SSD盘。
- 对于Worker节点，创建集群时推荐选择挂载数据盘。这个盘是专门提供给/var/lib/docker存放本地镜像。可避免后续如果镜像太多根磁盘容量不够的问题。在运行一段时间后，本地会存在很多无用的镜像。比较快捷的方式就是，先下线这台机器，重新构建这个磁盘，然后再上线。 

**磁盘大小**

Kubernetes节点需要的磁盘空间也不小，Docker镜像、系统日志、应用日志都保存在磁盘上。创建Kubernetes集群的时候，要考虑每个节点上要部署的Pod数量，每个Pod的日志大小、镜像大小、临时数据，再加上一些系统预留的值。

Kubernetes集群中，ECS操作系统占用3G左右的磁盘空间，建议预留ECS操作系统8G的空间。剩余的磁盘空间由Kubernetes资源对象使用。

## 是否立即构建Worker节点

​      创建集群时，      

节点类型

若选择：       

- 按量付费，可在创建集群时，构建Worker节点。 
- 包年包月，创建集群时，可先不构建Worker节点，根据后续需求，单独购买ECS添加进集群。 

## 网络选择

- 如果需要连接外部的一些服务，如RDS等，则需要考虑复用原有的VPC，而不是创建一个新的VPC。因为VPC间是隔离的，您可以通过创建一个新的交换机，把运行Kubernetes的机器都放在这个交换机下，从而便于管理。
- 在Kubernetes集群创建时提供两种网络插件：Terway和Flannel，可参考[如何选择Kubernetes集群网络插件：Terway和Flannel](https://help.aliyun.com/document_detail/86949.html#concept-qsd-ckm-q2b)。 
- Pod网络CIDR不能设置太小，如果太小，可支持的节点数量就会受限。这个值的设置需要与高级选项中的节点Pod数量综合考虑。例如：Pod网络CIDR的网段是/16，那么就有256*256个地址，如果每个节点Pod数量是128，则最多可以支持512个节点。 

## 使用多可用区

阿里云支持多Region（地域），每个Region下又有不同的可用区。可用区是指在同一地域内，电力和网络互相独立的物理区域。多可用区能够实现跨区域的容灾能力。同时也会带来额外的网络延时。创建Kubernetes集群时，您可选择创建多可用区Kubernetes集群。参见[创建多可用区 Kubernetes 集群](https://help.aliyun.com/document_detail/86493.html#task-hyb-xwf-vdb)。 

## 声明每个Pod的resource

在使用Kubernetes集群时，经常会遇到：在一个节点上调度了太多的Pod，导致节点负载太高，没法正常对外提供服务的问题。

为避免上述问题，在Kubernetes中部署Pod时，您可以指定这个Pod需要Request及Limit的资源，Kubernetes在部署这个Pod的时候，就会根据Pod的需求找一个具有充足空闲资源的节点部署这个Pod。下面的例子中，声明Nginx这个Pod需要1核CPU，1024M的内存，运行中实际使用不能超过2核CPU和4096M内存。

​            

```
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx
    resources: # 资源声明
      requests:
        memory: "1024Mi"
        cpu: "1000m"
      limits:
        memory: "4096Mi"
        cpu: "2000m"
```

Kubernetes采用静态资源调度方式，对于每个节点上的剩余资源，它是这样计算的：`节点剩余资源=节点总资源-已经分配出去的资源`，并不是实际使用的资源。如果您自己手动运行一个很耗资源的程序，Kubernetes并不能感知到。 

另外所有Pod上都要声明resources。对于没有声明resources的Pod，它被调度到某个节点后，Kubernetes也不会在对应节点上扣掉这个Pod使用的资源。可能会导致节点上调度过去太多的Pod。

## 日常运维

- 日志

  创建集群时，请勾选日志服务。 

- 监控

  阿里云容器服务与云监控进行集成，通过配置节点监控，提供实时的监控服务。通过添加监控告警规则，节点上的资源使用量很高的时候，可快速定位问题。

  ​         通过容器服务创建Kubernetes集群时，会自动在云监控创建两个应用分组：一个对应Master节点，一个对应Worker节点。我们可以在这两个组下面添加一些报警规则，对组里所有的机器生效。后续加入的节点，也会自动出现在组里，不用单独再去配置报警规则。        
  ​        [![img](http://static-aliyun-doc.oss-cn-hangzhou.aliyuncs.com/assets/img/24069/154587392914010_zh-CN.png)](http://static-aliyun-doc.oss-cn-hangzhou.aliyuncs.com/assets/img/24069/154587392914010_zh-CN.png)        
  ​       

  ​        主要配置ECS资源的报警规则就可以了。         

  ​                   

  说明

  - 对于ECS的监控，日常运维请设置cpu，memory，磁盘等的报警规则。且尽量将/var/lib/docker放在一个独立的盘上。 

  ![img](http://static-aliyun-doc.oss-cn-hangzhou.aliyuncs.com/assets/img/24069/154587392914012_zh-CN.png)

## 启动时等待下游服务，不要直接退出

有些应用可能会有一些外部依赖，比如需要从数据库（DB）读取数据或者依赖另外一个服务的接口。应用启动的时候，外部依赖未必都能满足。手工运维的时候，通常采用依赖不满足立即退出的方式，也就是所谓的failfast，但是在Kubernetes中，这种策略不再适用。原因在于Kubernetes中多数运维操作都是自动的，不需要人工介入，比如部署应用，您不用自己选择节点，再到节点上启动应用，应用fail，也不用手动重启，Kubernetes会自动重启应用。负载增高，还可以通过HPA自动扩容。  

针对启动时依赖不满足这个场景，假设有两个应用A和B，A依赖B，刚好运行在同一个节点上。这个节点因为某些原因重启了，重启之后，A首先启动，这个时候B还没启动，对A来说就是依赖不满足。如果A还是按照传统的方式直接退出，当B启动之后，A也不会再启动，必须人工介入处理才行。

Kubernetes的最好的做法是启动时检查依赖，如果不满足，轮询等待，而不是直接退出。可以通过[Init Container](https://kubernetes.io/docs/concepts/workloads/pods/init-containers/#what-can-init-containers-be-used-for)完成这个功能。 

## 配置restart policy

Pod运行过程中进程退出是个很常见的问题，无论是代码里的一个bug，还是占用内存太多，都会导致应用进程退出，Pod退出。您可在Pod上配置restartPolicy，就能实现Pod挂掉之后自动启动。

​            

```
apiVersion: v1
kind: Pod
metadata:
  name: tomcat
spec:
  containers:
  - name: tomcat
    image: tomcat
    restartPolicy: OnFailure # 
```

restartPolicy有三个可选值

- Always： 总是自动重启 
- OnFailure：异常退出才自动重启 （进程退出状态非0） 
- Never：永远不重启 

## 配置Liveness Probe和Readiness Probe

Pod处于Running状态和Pod能正常提供服务是完全不同的概念，一个Running状态的Pod，里面的进程可能发生了死锁而无法提供服务。但是因为Pod还是Running的，Kubernetes也不会自动重启这个Pod。所以我们要在所有Pod上配置Liveness  Probe，探测Pod是否真的存活，是否还能提供服务。如果Liveness Probe发现了问题，Kubernetes会重启Pod。 

Readiness  Probe用于探测Pod是不是可以对外提供服务。应用启动过程中需要一些时间完成初始化，在这个过程中是没法对外提供服务的，通过Readiness  Probe，可以告诉Ingress或者Service能不能把流量转发给这个Pod上。当Pod出现问题的时候，Readiness  Probe能避免新流量继续转发给这个Pod。 

​            

```
apiVersion: v1
kind: Pod
metadata:
  name: tomcat
spec:
  containers:
  - name: tomcat
    image: tomcat
    livenessProbe:
      httpGet:
        path: /index.jsp
        port: 8080
      initialDelaySeconds: 3
      periodSeconds: 3
    readinessProbe:
      httpGet:
        path: /index.jsp
        port: 8080
```

## 每个进程一个容器

​      很多刚刚接触容器的人喜欢按照旧习惯把容器当作虚拟机（VM）使用，在一个容器里放多个进程：监控进程、日志进程、sshd进程、甚至整个Systemd。这样操作存在两个问题：       

- 判断Pod整体的资源占用会变复杂，不方便实施前面提到resource limit。
- 容器内只有一个进程的情况，进程挂了，外面的容器引擎可以清楚的感知到，然后重启容器。如果容器内有多个进程，某个进程挂了，容器未必受影响，外部的容器引擎感知不到容器内有进程退出，也不会对容器做任何操作，但是实际上容器已经不能正常工作了。

如果有几个进程需要协同工作，在Kubernetes里也可以实现，例如：nginx和php-fpm，通过Unix domain socket通信，我们可以用一个包含两个容器的Pod，unix socket放在两个容器的共享volume中。 

## 确保不存在SPOF（Single Point of Failure）

如果应用只有一个实例，当实例失败的时候，虽然Kubernetes能够重启实例，但是中间不可避免地存在一段时间的不可用。甚至更新应用，发布一个新版本的时候，也会出现这种情况。在Kubernetes里，尽量避免直接使用Pod，尽可能使用Deployment/StatefulSet，并且让应用的Pod在两个以上。