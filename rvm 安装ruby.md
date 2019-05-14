rvm 安装ruby

首先安装rvm安装会使用的包：

yum install gcc-c++ patch readline readline-devel zlib zlib-devel libyaml-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison iconv-devel sqlite-devel

之后便是安装rvm:

```shell
curl -sSL https://rvm.io/mpapis.asc | gpg --import -
curl -L get.rvm.io | bash -s stable
```

配置rvm的运行环境

```bash
source /etc/profile.d/rvm.sh
rvm reload
```

输入一下命令检查安装情况

```undefined
rvm requirements run
```

将显示：

```shell
Checking requirements for centos.
Requirements installation successful.
```

最后便可安装ruby了，当然版本可以任选，反正我选2.4.4

```
rvm install 2.4.4
```

检查安装情

```
rvm list
```

设置默认运行的ruby版本

```
rvm use 2.4.2 --default
```