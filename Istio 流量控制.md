Istio 流量控制

以下这些大致就是你需要遵循的，以获得Istio的不同版本的应用程序的步骤：

1. 创建一个Kubernetes集群并安装带有sidecare自动注入的Istio。
2. 使用你选择的语言创建Hello World应用程序，创建Docker镜像并将其推送到公共镜像仓库。
3. 为你的容器创建Kubernetes Deployment和Service。
4. 创建Gateway以启用到群集的HTTP(S)流量。
5. 创建[VirtualService](https://istio.io/docs/reference/config/istio.networking.v1alpha3/#VirtualService)，通过Gateway公开Kubernetes服务。
6. （可选）如果要创建多个版本应用程序，请创建[DestinationRule](https://istio.io/docs/reference/config/istio.networking.v1alpha3/#DestinationRule)以定义可从VirtualService引用的subsets

Deployment和Service

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: myapp       #两个Deployment使用同一个标签
        version: v1      #区分版本
      annotations:
        prometheus.io/scrape: 'true'
    spec:
      containers:
      - name: myapp
        image: ikubernetes/myapp:v1
        imagePullPolicy: Always
        ports:
        - containerPort: 80
          protocol: TCP
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: myapp2
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: myapp     #两个Deployment使用同一个标签
        version: v2    #区分版本
      annotations:
        prometheus.io/scrape: 'true'
    spec:
      containers:
      - name: myapp
        image: ikubernetes/myapp:v2
        imagePullPolicy: Always
        ports:
        - containerPort: 80
          protocol: TCP
---
apiVersion: v1
kind: Service
metadata:
  name: myapp
spec:
  selector:
    app: myapp   #service暴露不同版本的应用服务
  type: ClusterIP
  ports:
  - name: http
    port: 80
    targetPort: 80
```

使用kubectl apply -f xxx.yaml进行部署

到目前为止没有任何特定的针对Istio的内容。

### Gateway AND VirtualService

首先，我们需要为服务网格启用HTTP/HTTPS流量。 为此，我们需要创建一个网关。 Gateway描述了在边缘运行的负载均衡，用于接收传入或传出的HTTP/TCP连接。

让我们创建一个myapp-gateway.yaml文件：

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: myapp-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - myapp.yunfuzg.com
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: myapp
spec:
  hosts:
  - myapp.yunfuzg.com
  gateways:
  - myapp-gateway    #VirtualService对应的Gateway
  http:
    route:
    - destination:
        port:
          number: 80
        host: myapp     #kubernetes的service
```

kubectl apply -f myapp-gateway.yaml

此时，我们为集群启用了HTTP流量。 我们需要将之前创建的Kubernetes服务映射到Gateway。 我们将使用VirtualService执行此操作。

[VirtualService](https://istio.io/docs/reference/config/istio.networking.v1alpha3/#VirtualService)实际上将Kubernetes服务连接到Istio网关。 它还可以执行更多操作，例如定义一组流量路由规则，以便在主机被寻址时应用，但我们不会深入了解这些细节。VirtualService与特定网关绑定，并定义引用Kubernetes服务的主机。

当我们在浏览器中打开,您将看到应用程序的v1和v2版本交替出现

### DestinationRule

在某些时候，你希望将应用更新为新版本。 也许你想分割两个版本之间的流量。你需要创建一个[DestinationRule](https://istio.io/docs/reference/config/istio.networking.v1alpha3/#DestinationRule)来定义是哪些版本，在Istio中称为subset。

如果您想将服务仅指向v2，该怎么办？ 这可以通过在VirtualService中指定subset来完成，但我们需要首先在DestinationRules中定义这些subset。 DestinationRule本质上是将标签映射到Istio的subset。

创建一个myapp-destinationrule.yaml文件：

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: myapp-destinationrule
spec:
  host: myapp
  subsets:
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: myapp-destinationrule
spec:
  host: myapp
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
```

```shell
kubectl apply -f myapp-destinationrule.yaml
```

现在你可以从VirtualService来引用v2 subset：

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: myapp
spec:
  hosts:
  - myapp.yunfuzg.com
  gateways:
  - myapp-gateway    #VirtualService对应的Gateway
  http:
    route:
    - destination:
        port:
          number: 80
        host: myapp     #kubernetes的service
        subset: v1      #这里开启了流量的百分比，如不需要可注释掉
      weight: 20
    - destination:
        port:
          number: 80
        host: myapp
        subset: v2
      weight: 80
```

