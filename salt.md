1.更新salt yum源

yum  -y install https://repo.saltstack.com/yum/redhat/salt-repo-latest-2.el7.noarch.rpm 

2.安装master端

yum -y install salt-master

yum -y install salt-ssh

安装salt-minion端

 yum -y install salt-minion

systemctl enable salt-minion.service && systemctl start salt-minion.service

安装WEB端

yum -y install salt-api

安装halite及依赖文件

yum -y install python-pip 

pip install --upgrade pip

pip install -U halite

pip install cherrypy

pip install paste

yum -y install python-devel gcc

pip install gevent

pip install pyopenssl

```bash
vim /etc/salt/master
```

 

 

```bash
external_auth:
  pam:
    testuser:            ``//``此用户设置为系统在用的用户
      - .*
      - '@runner'
```

 

```bash
halite:
  level: 'debug'
  server: 'cherrypy'
  host: '0.0.0.0'
  port: '8080'
  cors: False
  tls: True
  certpath: '/etc/pki/tls/certs/localhost.crt'
  keypath: '/etc/pki/tls/certs/localhost.key'
  pempath: '/etc/pki/tls/certs/localhost.pem'
```

运行salt-call tls.create_self_signed_cert tls

重新启动服务 salt-minion salt-master salt-api

salt其他：https://blog.csdn.net/ZZL95415/article/details/80672553

​                  https://www.cnblogs.com/shhnwangjian/p/5985868.html

salt-syndic安装配置：

yum -y install salt-syndic

配置master：

/etc/salt/master

order_masters:  True    #开启多层master

配置syndic端：

/etc/salt/proxy

master：IP

/etc/salt/master

synddic_master: IP

重新启动两个服务

salt执行脚本

salt “*” cmd.script 脚本名称

salt-cp “*”  源目录  远程目录

salt “*”

salt “*” pillar.item

```bash
salt 'linux-node1*' grains.ls # 列出ID为linux-node1的主机，grains的所有key
salt 'linux-node1*' grains.items  # 列出主机的详细信息，可用于资产管理
salt '*' grains.item os  # 列出所有主机的系统版本
salt '*' grains.item fqdn_ip4  # 列出所有主机的IP地址
```

salt "fqdn_ip4:IP" cmd.run "comm"   #根据IP执行命令



salt-ssh：/etc/salt/roster

xd_redis_s1:
​    host: 172.16.23.120
   user: root
   priv: /root/.ssh/id_rsa





elastic search：https://blog.csdn.net/qq_29767087/article/details/79791173