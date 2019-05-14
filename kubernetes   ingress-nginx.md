kubernetes   ingress-nginx

安装：

下载 : git clone https://github.com/kubernetes/ingress-nginx.git

```shell
cd ingress-nginx/deploy
ls
kubectl apply -f ./
kubectl apply -f monitoring/
vim ingress-nginx-svc.yaml
kind: Service
apiVersion: v1
metadata:
  name: ingress-nginx
  namespace: ingress-nginx
  labels:
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
spec:
  type: NodePort
  selector:
    app.kubernetes.io/name: ingress-nginx
    app.kubernetes.io/part-of: ingress-nginx
  ports:
    - name: http
      port: 80
      targetPort: http
      nodePort: 80
    - name: https
      port: 443
      targetPort: https
      nodePort: 443
kubectl apply -f ingress-nginx-svc.yaml
```

已域名的方式访问

vim ingress.yaml

```yaml
apiVersion: extensions/v1beta1      #api版本
kind: Ingress       #清单类型
metadata:           #元数据
  name: ingress-myapp    #ingress的名称
  namespace: default     #所属名称空间
  annotations:           #注解信息
    kubernetes.io/ingress.class: "nginx"
spec:      #规格
  rules:   #定义后端转发的规则
  - host: myapp.magedu.com    #通过域名进行转发
    http:
      paths:       
      - path:       #配置访问路径，如果通过url进行转发，需要修改；空默认为访问的路径为"/"
        backend:    #配置后端服务
          serviceName: myapp
          servicePort: 80
```

加入tls服务证书方式访问：

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: ingress-tomcat-tls
  namespace: default
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  tls:
  - hosts:
    - tomcat.magedu.com
    secretName: tomcat-ingress-secret
  rules:
  - host: tomcat.magedu.com
    http:
      paths:
      - path:
        backend:
          serviceName: tomcat
          servicePort: 8080
```