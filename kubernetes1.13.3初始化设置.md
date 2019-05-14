

kubernetes1.13.3初始化设置

1.系统准备

修改主机名称:hostnamectl set-hostname --static

修改/etc/hosts文件

关闭防火墙：systemctl stop firewalld && systemctl disable firewalld

禁用selinux：setenforce 0

​        vi /etc/selinux/config

创建/etc/sysctl.d/k8s.conf

```shell
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
vm.swappiness=0

modprobe br_netfilter
sysctl -p /etc/sysctl.d/k8s.conf
```

kube-proxy开启ipvs前置条件：各个节点执行

```shell
cat > /etc/sysconfig/modules/ipvs.modules <<EOF
#!/bin/bash
modprobe -- ip_vs
modprobe -- ip_vs_rr
modprobe -- ip_vs_wrr
modprobe -- ip_vs_sh
modprobe -- nf_conntrack_ipv4
EOF
chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack_ipv4
```

安装docker

```shell
cd /etc/yum.repos.d
yum -y install wget
mv CentOS-Base.repo CentOS-Base.repo.bck
wget -O CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum list docker-ce --showduplicates | sort -r 
yum install -y --setopt=obsoletes=0 docker-ce-18.06.2.ce-3.el7
```

确认iptables filter规则默认策略改为ACCEPT

```SHELL
iptables -nvL
iptables -p FORWARD ACCEPT
```

安装kubeadm kubelet kubectl

```shell
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes Repo
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF


yum makecache fast
yum install -y kubelet kubeadm kubectl
```

关闭swap

```
swapoff -a
vm.swappiness=0
sysctl -p /etc/sysctl.d/k8s.conf
vi /etc/sysconfig/kubelet
KUBELET_EXTRA_ARGS=--fail-swap-on=false
systemctl enable kubelet.service && systemctl enable docker.service
systemctl start docker && systemctl start kubelet.service
```

kubeadm init

```shell
kubeadm init --kubernetes-version=v1.13.3 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.139.133 --service-cidr=10.96.0.0/12  --ignore-preflight-errors=Swap


api:
kubeadm join 10.29.58.165:6443 --token 1zyyeq.vm2ix6sjhpktquab --discovery-token-ca-cert-hash sha256:64c4a2f90c10d7129bf81d32a0740324a5f2f6d0932dda0b381bcd24f7d3ad81
```

kubernetes镜像下载

```shell
#!/bin/bash
set -e
KUBE_VERSION=v1.13.3
KUBE_PAUSE_VERSION=3.1
ETCD_VERSION=3.2.24
DNS_VERSION=1.2.6
GCR_URL=k8s.gcr.io
ALIYUN_URL=mirrorgooglecontainers
images=(kube-proxy:${KUBE_VERSION}
kube-scheduler:${KUBE_VERSION}
kube-controller-manager:${KUBE_VERSION}
kube-apiserver:${KUBE_VERSION}
pause:${KUBE_PAUSE_VERSION}
etcd:${ETCD_VERSION})
#coredns:${DNS_VERSION})
for imageName in ${images[@]}
do
  docker pull $ALIYUN_URL/$imageName
  docker tag  $ALIYUN_URL/$imageName $GCR_URL/$imageName
  docker rmi $ALIYUN_URL/$imageName
done
docker pull coredns/coredns:1.2.6
docker tag coredns/coredns:1.2.6 $GCR_URL/coredns:1.2.6
docker rmi coredns/coredns:1.2.6
docker images  
```

```
mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
yum install -y bash-completion
source /usr/share/bash-completion/bash_completion
source <(kubectl completion bash)
echo "source <(kubectl completion bash)" >> ~/.bashrc
echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> ~/.bash_profile
source ~/.bash_profile
```

集群初始化如果遇到问题，可以使用下面的命令进行清理：

```shell
kubeadm reset
ifconfig cni0 down
ip link delete cni0
ifconfig flannel.1 down
ip link delete flannel.1
rm -rf /var/lib/cni/

```

安装Pod Network

```shell
wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
vi kube-flannel.yml
containers:
      - name: kube-flannel
        image: quay.io/coreos/flannel:v0.10.0-amd64
        command:
        - /opt/bin/flanneld
        args:
        - --ip-masq
        - --kube-subnet-mgr
        - --iface=ens33
```

使用flannel网络如出现node的pod无法访问api-service，查看iptables 并添加规则

```shell
iptables -nvL
如FORWARD为DORP,需要运行以下命令修改为ACCEPT
iptables -p FORWARD ACCEPT
node的pod无法访问api-service，需要根据POD网络运行以下命令
iptables -t nat -I POSTROUTING -s 10.88.0.0/16 -j MASQUERADE
其中10.88.0.0/16为POD使用的网络
关闭nodes端口限制
修改/etc/kubernetes/manifests/kube-apiserver.yaml
添加--service-node-port-range=1-65535
```

master node参与工作负载

```shell
kubectl describe node master01 | grep Taint
kubectl taint nodes master01 node-role.kubernetes.io/master- #去除污点
kubectl taint nodes master01 node-role.kubernetes.io/master=:NoSchedule    #设置污点
```

测试dns

```shell
kubectl run curl --rm --image=radial/busyboxplus:curl -it
nslookup kubernetes.default
```

node节点加入集群

```shell
kubeadm join 192.168.139.133:6443 --token p41q9u.z538a4a9md2mnos6 --discovery-token-ca-cert-hash sha256:2771954395490827d0974771f08ae01baa9cd3e7691fcf72e17196aea2ea8387

node节点加入集群，或者忘记kubeadm join的解决方法
重新生成一条永久的token
在master节点上运行
kubeadm token create --ttl 0
kubeadm token list
获取ca证书sha256编码hash值
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
```

修改kube-proxy为ipvs

```shell
kubectl edit cm kube-proxy -n kube-system
mode: “ipvs”
```

移除node

```shell
kubectl cordon node01 #标记node01不可调度
kubectl drain node2 --delete-local-data --force --ignore-daemonsets
kubectl uncordon node01 #node01可重新调度
在node上执行
kubeadm reset
ifconfig cni0 down
ip link delete cni0
ifconfig flannel.1 down
ip link delete flannel.1ip
rm -rf /var/lib/cni/
在master上执行
kubectl delete node node01
```

# metrics-server 安装

```shell
yum -y install git
git clone https://github.com/stefanprodan/k8s-prom-hpa.git
cd k8s-prom-hpa/metrics-server/
kubectl apply -f ./
docker pull mirrorgooglecontainers/metrics-server-amd64:v0.3.1
docker tag mirrorgooglecontainers/metrics-server-amd64:v0.3.1 k8s.gcr.io/metrics-server-amd64:v0.3.1
docker rmi mirrorgooglecontainers/metrics-server-amd64:v0.3.1
cd /root/k8s-prom-hpa
kubectl apply -f ./namespaces.yaml
kubectl apply -f ./prometheus
```

kubernetes认可证书建立：

```shell
cd /etc/kubernetes/pki/
(umask 077; openssl genrsa -out bill.key 2048)  #生成私钥
openssl req -new -key bill.key -out bill.csr -subj "/CN=bill"   #生成证书签署请求
openssl x509 -req -in bill.csr -CA ./ca.crt -CAkey ./ca.key -CAcreateserial -out bill.crt -days 3650                #证书签署
openssl x509 -in bill.crt -text -noout #查看证书签署
kubectl create secret generic bill --from-file=bill.crt=./bill.crt --from-file=bill.key=./bill.key -n prom   #将证书挂载到k8s的secret上
kubectl get secret   #查看证书
cd /root/k8s-prom-hpa/custom-metrics-api/
vim custom-metrics-apiserver-deployment.yaml
    - --tls-cert-file=/var/run/serving-cert/bill.crt
    - --tls-private-key-file=/var/run/serving-cert/bill.key
    
kubectl apply -f ./
```

https://github.com/stefanprodan/k8s-prom-hpa