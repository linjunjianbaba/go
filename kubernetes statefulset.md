kubernetes statefulset

statefulset

1.稳定且唯一的网络标识符

2.稳定且持久的存储

3.有序，平滑的部署和扩展

4.有序，平滑的删除和终止

5.有序的滚动更新 10->1

三个组件：Headless Service(无头服务)  ,StatufulSet（控制器）, volumeClaimTemplate（pvc模板）

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysvc
  namespace: default
  labels:
    app: mysvc
spec:
  ports:
  - name: http
    port: 80
  clusterIP: None
  selector: 
    app: mystate-pod
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mystateful
  namespace: default
spec:
  serviceName: mysvc
  reolicas: 2
  selector:
    matchLabels:
      app: mystate-pod
  template:
    metadata:
      labels:
        app: mystate-pod
    spec:
      containers:
      - name: mystate-pod
        image: ikubernetes/myapp:v5
        imagePullPolicy: IfNotPresent
        ports:
        - name: http
          containerPort: 80
        volumeMounts:
        - name: myappdata
          mountPath: /usr/share/nginx/html/
      volumeClaimTemplate:
      - metadata:
          name: myappdata
        spec:
          accessModes: ["ReadWriteOnce"]
          storageClassName: "gluster-dynamic"
          resources:
            reauests:
              storage: 2Gi
```

