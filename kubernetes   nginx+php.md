kubernetes   nginx+php

helm create nginx

1.生成挂载证书yaml，或直接挂载证书

kubectl create secret tls nginx-web --cert=tls.pem --key=tls.key --dry-run -o yaml

新建的yaml文件复制到nginx/templates下

修改ingress.yaml文件

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: nginx
    kubernetes.io/tls-acme: "true"
  generation: 1
  labels:
    app.kubernetes.io/instance: honking-chipmunk
    app.kubernetes.io/managed-by: Tiller
    app.kubernetes.io/name: nginx
    helm.sh/chart: nginx-0.1.0
  name: honking-chipmunk-nginx
  namespace: default
spec:
  rules:
  - host: www.yunfuzg.com
    http:
      paths:
      - backend:
          serviceName: honking-chipmunk-nginx
          servicePort: http
        path: /
  tls:
  - hosts:
    - www.yunfuzg.com
    secretName: yunfuzg
```

helm create mysql