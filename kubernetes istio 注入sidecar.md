kubernetes istio 注入sidecar

```shell 
查看kubernetes api
kubectl api-versions | grep admissionregistration
在kube-apiserver加入参数  （非必要）
- --enable-admission-plugins=NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota
自动注入控制：
可通过在sidecar-injector的configmap中设置policy=disabled字段来设置是否启用自动注入（此处为全局控制是否启用自动注入功能）
kubectl edit cm istio-sidecar-injector -nistio-system #查看
为需要自动注入的namespace打上标签istio-injection: enabled（此处为ns级别的自动注入控制）。
kubectl get namespace -L istio-injection   #查询
kubectl label namespace default istio-injection=enabled #添加标签
kubectl logs -f istio-sidecar-injector -n istio-system #查看sidecar注入是否成功
```

