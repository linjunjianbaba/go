kubernetes deployment yaml

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: myapp
  namespace: default
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: ikubernetes/myapp:v1
        imagePullPolicy: IfNotPresent
        ports:
        - name: http
          containerPort: 80
        volumeMounts:
        - name: html
          mountPath: /usr/share/nginx/html
        - name: nginxwww
          mountPath: /etc/nginx/conf.d/
      volumes:
      - name: html
        hostPath:
          path: /data/pod/volume1
          type: DirectoryOrCreate
      - name: nginxwww
        configMap:
          name: nginx-www
---
apiVersion: v1
kind: Service
metadata:
  name: myservice
  namespace: default
spec:
  selector:
    app: myapp
  type: NodePort
  ports:
  - nema: http
    port: 80
    targetPort: 80
    protocol: TCP
    nodePort: 81
     
```

kubectl-shell

下载python2.7.14

```shell
# wget https://www.python.org/ftp/python/2.7.14/Python-2.7.14.tgz
```

解压Python包

```shell
tar -zxvf Python-2.7.14.tgz
```

检查&准备编译环境

```shell
yum install gcc* openssl openssl-devel ncurses-devel.x86_64  bzip2-devel sqlite-devel python-devel zlib
```

**安装**

```shell
cd Python-2.7.14
./configure --prefix=/usr/local
make && make install 
```

备份旧版，yum等组件依赖于2.7.5工作

```shell
mv /usr/bin/python /usr/bin/python2.7.5
ln -s /usr/local/bin/python2.7 /usr/bin/python
```

修正yum等组件python

```shell
[root@localhost bin]# vim /usr/bin/yum
首行的#!/usr/bin/python 改为 #!/usr/bin/python2.7.5 
[root@localhost bin]# vim /usr/libexec/urlgrabber-ext-down
首行的#!/usr/bin/python 改为 #!/usr/bin/python2.7.5
```
### Pip安装

```shell
wget https://bootstrap.pypa.io/get-pip.py
python get-pip.py
ln -s /usr/local/bin/pip2.7 /usr/bin/pip   
```

daemonset yaml

```yaml
apiVersion:  extensions/v1beta1
kind: Daemonset
metadata:
  name: mydaemon
  labels:
    app: mydaemon
spec:
  selector:
    matchLabels:
      app: mydaemon
  template:
    metadata:
      labels:
        app: mydaemon
    spec:
      containers:
      - name: mydaemon
        image: ikubernetes/myapp:v3
        imagePullPolicy: IfNotPresent
        volumeMounts:
        - name: nginxwww
          mountPath: /etc/nginx/conf.d/
      volumes:
      - name：nginxwww
        configMag:
          name: nginx-www
       
```

