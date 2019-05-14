kubernetes 结合jenkins

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: jenkins
  labels:
    app: jenkins
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: jenkins
    spec:
      containers:
      - name: jenkins
        image: jenkins/jenkins:lts
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
          name: web
          protocol: TCP
        volumeMounts:
        - name: jenkins-home
          mountPath: /var/jenkins_home
        env:
        - name: JAVA_OPTS
          value: "-Duser.timezone=Asia/Shanghai -XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap -XX:MaxRAMFraction=1 -Dhudson.slaves.NodeProvisioner.MARGIN=50 -Dhudson.slaves.NodeProvisioner.MARGIN0=0.85"
      volumes:
      - name: jenkins-home
        persistentVolumeClaim:
          claimName: jenkins-pvc
---
apiVersion: v1
kind: Service
metadata:
  name: jenkins
  labels:
    app: jenkins
spec:
  selector:
    app: jenkins
  ports:
  - port: 8080
    targetPort: 8080
    name: web
    nodePort: 308080
  - port: 50000
    targetPort: 50000
    name: agent
  type: NodePort
  
```

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: jenkins-pv
  labels:
    app: jenkins-pv
spec:
  hostPath:
    path: /data/jenkins
    type: DirectoryOrCreate
  accessModes: ["ReadWriteMany","ReadWriteOnce"]
  capacity:
    storage: 10G
```

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata: 
  name: jenkins-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 8Gi
```

1.首先建立jenkins存储PV和PVC

