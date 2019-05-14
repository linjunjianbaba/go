kubernetes-在pod里面的容器不能ping外部ip

- 节点 192.168.0.6 中pod的ip段为 172.30.60.2/24
- 节点 192.168.0.7 中pod的ip段为 172.30.25.2/24

在 节点 192.168.0.6 增加一条路由规则：

```bash
# /sbin/iptables -t nat -I POSTROUTING -s 172.30.60.0/24 -j MASQUERADE
```

在 节点 192.168.0.7 增加一条路由规则：

```bash
# /sbin/iptables -t nat -I POSTROUTING -s 172.30.25.0/24 -j MASQUERADE
```

