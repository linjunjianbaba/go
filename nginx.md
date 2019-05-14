nginx 

yum安装：

vim /etc/yum.repos.d/nginx.repo

[nginx]

name=nginx repo

baseurl=http://nginx.org/packages/centos/7/$basearch/

gpgcheck=0

enable=1



yum -y install nginx

源码安装：http://nginx.org/en/download.html

