istio https ingressgateway

将证书文件写入到istio-gateway中

```shell
kubectl create secret tls -n istio-system istio-ingressgateway-certs --key ./tls.key --cert ./tls.pem  #必须以istio-ingressgateway-cert命名才能被istio-gateway正确挂载
kubectl exec -it -n istio-system $(kubectl -n istio-system get pods -l istio=ingressgateway -o jsonpath='{.items[0].metadata.name}') -- ls -al /etc/istio/ingressgateway-certs    #查看证书挂在是否正确
```

编写istiogateway  yaml文件

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: nginx-gateway
spec:
  selector:
    istio: ingressgateway # use Istio default gateway implementation
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    tls:
      httpsRedirect: true  #访问http自动跳转到https
    hosts:
    - "*.yunfuzg.com"
  - port:
      number: 443
      name: https
      protocol: HTTPS
    hosts:
    - "*.yunfuzg.com"
    tls:
      mode: SIMPLE
      serverCertificate: /etc/istio/ingressgateway-certs/tls.crt           #证书是挂在指定位置
      privateKey: /etc/istio/ingressgateway-certs/tls.key
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: nginx
spec:
  hosts:
  - "*.yunfuzg.com"
  gateways:
  - nginx-gateway
  http:
  - match:
    - uri:
        prefix: /
    route:
    - destination:
        port:
          number: 80    #对应pod service port 
        host: nginx     #对应pod service

```

