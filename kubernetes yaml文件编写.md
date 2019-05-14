kubernetes yaml文件编写

```
apiVersion： apps/v1 #api版本
kind： Deployment          #资源类型，一般已大写开头如DaemonSet
metadata:
  name:
  namespace:
  labels:
     app: value
spec:
  replicas: 2
  selector:
    machLabels:
      app: value
  spec:
    containers:
    - name: myapp
      image: buxy
      imagePullPolicy: IfNotPresent
      ports:
      - name:
        containerPort:        
      command:
      args:
      env:
      - name:
        value:
      volumounts:
      - name:
        mountPath:
        readOnly:
      resources:
        limits:
          cpu:
          memory:
        requests:
          cpu:
          memory:
        
        
```

