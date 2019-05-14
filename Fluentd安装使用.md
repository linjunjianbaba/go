Fluentd安装使用

\# 安装

curl -L https://toolbelt.treasuredata.com/sh/install-redhat-td-agent2.sh | sh 

\# 支持命令

/etc/init.d/td-agent start

/etc/init.d/td-agent stop

/etc/init.d/td-agent restart

/etc/init.d/td-agent status

根据配置文件启动

fluentd -c dokcer_in01.conf 

fluentd配置文件所在目录：/etc/td-agent/

fluentd二进制文件所在目录：/opt/td-agent/embedded/bin/

安装插件的方法：/opt/td-agent/embedded/bin/fluent-gem install [插件名称]

docker run -d --name test01 --log-driver=fluentd --log-opt tag="docker.{{.Name}}" --log-opt fluentd-async-connect=true -p 82:8080 vshop:v1

fluentd使用插件：

用于路径中加入tag：[fluent-plugin-forest](https://github.com/tagomoris/fluent-plugin-forest)

用于修改record：[fluent-plugin-record-reformer](https://github.com/sonots/fluent-plugin-record-reformer)

用于修改tag：[fluent-plugin-rewrite-tag-filter](https://github.com/fluent/fluent-plugin-rewrite-tag-filter)

用于正则匹配日志内容，进行筛选：[fluent-plugin-grep](https://github.com/sonots/fluent-plugin-grep)

\# 客户端需要安装的插件

/opt/td-agent/embedded/bin/fluent-gem install fluent-plugin-rewrite-tag-filter

/opt/td-agent/embedded/bin/fluent-gem install fluent-plugin-grep

/opt/td-agent/embedded/bin/fluent-gem install fluent-plugin-record-reformer

\# 服务器端需要安装的插件

/opt/td-agent/embedded/bin/fluent-gem install fluent-plugin-forest



elasticsearch配置

<source>
  type forward
  port 24224
  bind 0.0.0.0
</source>

<match docker.*>
  type elasticsearch
  host 127.0.0.1
  port 9200
  index_name docker
  type_name docker
  logstash_format true
  logstash_prefix docker
  logstash_dateformat %Y.%m.
  time_key vtm
  utc_index true
  flush_interval 5
</match>

elasticsearch docker启动：docker run -it  --name 容器名 --privileged -p 9200:9200 -p 9300:9300 -d 镜像名称

kibana dockers启动：docker run  --name --link elastic容器名:别名 -p 5601:5601 -d kibana

