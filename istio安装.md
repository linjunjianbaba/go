istio安装

什么是istio

- Istio 可以在虚拟机和容器中运行
- Istio 的组成
  - Pilot：服务发现、流量管理
  - Mixer：访问控制、遥测
  - Citadel：终端用户认证、流量加密
- Service mesh 关注的方面
  - 可观察性
  - 安全性
  - 可运维性
- Istio 是可定制可扩展的，组建是可拔插的
- Istio 作为控制平面，在每个服务中注入一个 Envoy 代理以 Sidecar 形式运行来拦截所有进出服务的流量，同时对流量加以控制
- 应用程序应该关注于业务逻辑（这才能生钱），非功能性需求交给 Service Mesh

下载安装包：https://github.com/istio/istio/releases

解压 安装包

使用helm安装：

```shell
helm install install/kubernetes/helm/istio --name istio --namespace istio-system
helm template install/kubernetes/helm/istio --name istio --namespace istio-system > $HOME/istio.yaml   #生成yaml，方便修改，使用kubectl进行不俗
kubectl apply -f $HOME/istio.yaml  
```

