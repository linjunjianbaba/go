Fluentd

安装：curl -L https://toolbelt.treasuredata.com/sh/install-redhat-td-agent2.sh | sh

启动：/etc/init.d/td-agent start

配置文件：/etc/td-agent/td-agentd.conf

可执行文件：/opt/td/agent/embedded/bin

<source>

  type forward

  port 24224

  bind 0.0.0.0

  log_level error

</source>

<match docker.*>

  type forest

  subtype file

  <template>

​    type file_alternative

​    path /home/lee/fluentd-log/${tag_parts[1]}/temp

  </template>

</match>