

# kubernetes  存储

## emptyDir存储卷演示

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-demo
  labels:
    app: myapp
spec:
  containers:
  - name: httpd
    image: busybox:latest
    imagePullPolicy: IfNotPresent
    command: ['/bin/httpd','-f','-h /data/web/html']
    ports: 
    - name: http
      containerPort: 80
    volumeMounts:
    - name: html
      mountPath: /data/web/html
  - name: busybox
    image: busybox
    imagePullPolicy: IfNotPresent
    volumeMounts:
    - name: html
      mountPath: /data
    command: ['/bin/sh','-c','while true; do echo ${date} >> /data/index.html; sleep 2; done' ]
  volumes:
  - name: html
    emptyDir: {}

```

mysql binlog日志反写sql语句 https://github.com/danfengcao/binlog2sql

```shell
python binlog2sql/binlog2sql.py -h127.0.0.1 -P3306 -u'root' -p'jnw5O6MvfI08OzmtkSGM' -d sibu_wesale_qrcode_03 -t safe_code_10 --start-file='mysql-bin.000544' --stop-file='mysql-bin.000545' -B |grep 91FEE3F8-6C38-4E63-943E-9892AECC963E > 4.sql
```

## hostPath存储卷演示

```shell
type：
DirectoryOrCreate  宿主机上不存在创建此目录  
Directory 必须存在挂载目录  
FileOrCreate 宿主机上不存在挂载文件就创建  
File 必须存在文件
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-vo
  namespaces: default
spec:
  containers:
  - name: pod-vo
    image: ikubernetes/myapp:v1
    volumeMounts:
    - name: html
      mountPath: /usr/share/nginx/html
  volumes:
    - name: html
      hostPath:
        path: /data/pod/volume1
        type: DirectoryOrCreate
```

## nfs共享存储卷演示

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nfs-pod
spec:
  containers:
  - name: nfs-pod
    image: ikubernetes/myapp:v1
    volumeMounts:
    - name: html
      mountPath: /usr/share/nginx/html
   volumes:
     - name: html
       nfs:
         path: /data/volumes
         server: server01
```

## PVC和PV的概念

我们前面提到kubernetes提供那么多存储接口，但是首先kubernetes的各个Node节点能管理这些存储，但是各种存储参数也需要专业的存储工程师才能了解，由此我们的kubernetes管理变的更加复杂的。由此kubernetes提出了PV和PVC的概念，这样开发人员和使用者就不需要关注后端存储是什么，使用什么参数等问题。

PersistentVolume（PV）是集群中已由管理员配置的一段网络存储。 集群中的资源就像一个节点是一个集群资源。 
PV是诸如卷之类的卷插件，但是具有独立于使用PV的任何单个pod的生命周期。 
该API对象捕获存储的实现细节，即NFS，iSCSI或云提供商特定的存储系统。

PersistentVolumeClaim（PVC）是用户存储的请求。PVC的使用逻辑：在pod中定义一个存储卷（该存储卷类型为PVC），定义的时候直接指定大小，pvc必须与对应的pv建立关系，pvc会根据定义去pv申请，而pv是由存储空间创建出来的。pv和pvc是kubernetes抽象出来的一种存储资源

虽然PersistentVolumeClaims允许用户使用抽象存储资源，但是常见的需求是，用户需要根据不同的需求去创建PV，用于不同的场景。而此时需要集群管理员提供不同需求的PV，而不仅仅是PV的大小和访问模式，但又不需要用户了解这些卷的实现细节。
对于这样的需求，此时可以采用StorageClass资源。这个在前面就已经提到过此方案。

PV是集群中的资源。 PVC是对这些资源的请求，也是对资源的索赔检查。 PV和PVC之间的相互作用遵循这个生命周期：

# PV的建立

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv001
  labels:
    name: pv001
spec:
  hostPath:
    path: /data/pv001
    type: DirectoryOrCreate
  accessModes: ["ReadWriteMany","ReadWriteOnce"]
  capacity:
    storage: 2G
```

# PVC的建立

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mypvc
  namespace: default
spec:
  accessModes: ["ReadWriteMany"]
  resources:
    requests:
      storage: 2Gi
---
apiVersion: v1
kind: Pod
metadata:
  name: pod-vol-pvc
  namespace: default
spec:
  containers:
  - name: pod-vol-pvc
    image: ikubernetes/myapp:v1
    volumeMounts:
    - name: html
      mountPath: /usr/share/nginx/html
  volumes:
    - name: html
      persistentVolumeClaim:
        claimName: mypvc
```

## Secret和configMap

创建secret的四种方式：

1、通过 --from-literal：

```shell
kubectl create secret  generic mysecret --form-literal=username=admin --from-literal=password=123456
```

2、通过 --from-file

3、通过 --from-env-file：

文件 env.txt 中每行 Key=Value 对应一个信息条目。如：

```shell
username=admin
password=123456
```

4、通过 YAML 配置文件

```shell
echo -n admin | base64
echo -n 123456 | base64
```

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
  data:
    username: YWRtaW4=
    password: MTIzNDU2
```

Pod 可以通过 Volume 或者环境变量的方式使用 Secret

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-secret
spec:
  containers:
  - name: pod-secret
    image: busybox:v1
    args:
      - /bin/sh
      - -c
      - sleep 10;touch /tmp/healthy;sleep 30000
    volumeMounts:
    - name: foo
      mountPath: "/etc/foo"
      readOnly: true
  volumes:
  - name: foo
    secret:
      secretName: mysecret
---
apiVersion: v1
kind: Pod
metadata:
  name: pod-secret
spec:
  containers:
  - name: pod-secret
    image: busybox
    args:
      - /bin/sh
      - -c
      - sleep 10;touch /tmp/healthy;sleep 30000
    volumeMounts:   
    - name: foo
      mountPath: "/etc/foo"
      readOnly: true
  volumes:
  - name: foo
    secret:
      secretName: mysecret
      items:    #自定义存放数据的文件名
      - key: username
        path: my-secret/my-username
      - key: password
        path: my-secret/my-password
```

通过 Volume 使用 Secret，容器必须从文件读取数据，会稍显麻烦，Kubernetes 还支持通过环境变量使用 Secret

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: po-secret-env
spec:
  containers:
  - name: po-secret-env
    image: busybox:v1
    args:
      - /bin/sh
      - -c
      - sleep 10;touch /tmp/healthy;sleep 30000
    env:
      - name: USERNAME
        valueFrom:
          secretKeyRef:
            name: mysecret
            key: username
      - name: PASSWORD
        valueFrom:
          secretKeyRef:
            name: mysecret
            key: password
            
```

通过环境变量 SECRET_USERNAME 和 SECRET_PASSWORD 成功读取到 Secret 的数据。
需要注意的是，环境变量读取 Secret 很方便，但无法支撑 Secret 动态更新。
Secret 可以为 Pod 提供密码、Token、私钥等敏感数据；对于一些非敏感数据，比如应用的配置信息，则可以用 ConfigMap

configmap是让配置文件从镜像中解耦，让镜像的可移植性和可复制性。许多应用程序会从配置文件、命令行参数或环境变量中读取配置信息。这些配置信息需要与docker
image解耦，你总不能每修改一个配置就重做一个image吧？ConfigMap 
API给我们提供了向容器中注入配置信息的机制，ConfigMap可以被用来保存单个属性，也可以用来保存整个配置文件或者JSON二进制大对象。

ConfigMap API资源用来保存key-value 
pair配置数据，这个数据可以在pods里使用，或者被用来为像controller一样的系统组件存储配置数据。虽然ConfigMap跟Secrets类似，但是ConfigMap更方便的处理不含敏感信息的字符串。
注意：ConfigMaps不是属性配置文件的替代品。ConfigMaps只是作为多个properties文件的引用。可以把它理解为Linux系统中的/etc目录，专门用来存储配置文件的目录。下面举个例子，使用ConfigMap配置来创建Kuberntes
Volumes，ConfigMap中的每个data项都会成为一个新文件。

与 Secret 一样，ConfigMap 也支持四种创建方式：
1、 通过 --from-literal：
每个 --from-literal 对应一个信息条目。

```shell
kubectl create configmap nginx-config --from-literal=nginx_port=80 --from-literal=server_name=myapp.k8sz.com
```

通过文件：

```shell
kubectl create configmap nginx-www --from-file=./www.conf 
kubectl create configmap tomcat-server --from-file=./server.xml
kubectl create configmap tomcat-cataline --from-file=./catalina.sh
```

1、环境变量方式注入到pod

```
apiVersion: v1
kind: Pod
metadata:
  name: pod-cm-1
spec:
  containers:
  - name: mypod
    image: ikubernetes/myapp:v1
    imagePullPolicy: IfNotPresent
    ports:
    - name: http
      containerPort: 80
    env:
    - name: NGINX_PORT
      valueFrom:
        configMapKeyRef:
          name: nginx-config
          key: nginx_port
    - name: NGINX_SERVAR
      valueFrom:
        configMapKeyRef:
          name: nginx-config
          key: server_name
    
  
```

2、存储卷方式挂载configmap：

Volume 形式的 ConfigMap 也支持动态更新

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-cm-2
  namespace: default
  labels:
    app: pod-cm-2
spec:
  containers:
  - name: pod-cm-2
    image: ikubernetes/myapp:v1
    imagePullPolicy: IfNotPresent
    ports:
    - name: http
      containerPort: 80
    volumeMounts:
    - name: nginxconf
      mountPath: /etc/nginx/conf.d/
      readOnly: true
  volumes:
  - name: nginxconf
    configMap:
      name: nginx-config
```

3、以nginx-www配置nginx

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod-cm-3
  namespace: default
  labels:
    app: cm3
spec:
  containers:
  - name: pod-cm-3
    image: ikubernetes/myapp:v1
    imagePullPolicy: IfNotPresent
    ports:
    - name: http
      containerPort: 80
    volumeMounts:
    - name: nginxwww
      mountPath: /etc/nginx/conf.d/
      readOnly: true
  volumes:
  - name: nginxwww
    configMap:
      name: nginx-www
```

