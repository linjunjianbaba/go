- 用`screen -dmS *session name*`来建立一个处于断开模式下的会话（并指定其会话名）。
- 用`screen -list `来列出所有会话。
- 用`screen -r *session name*`来重新连接指定会话。
- 用快捷键`CTRL-a d `来暂时断开当前会话。

Linux

linux初始化

1.更改yum源

yum clean

yum makecache

yum update

yum install yum-priorities 源优先级工具

repo下添加priority=对应优先级数字

/etc/yum/pluginconf.d/priorities.conf时间服务

*/5 * * * * /usr/sbin/ntpdate ntp.api.bz >> /dev/null 2>&1

lsmod

/etc/security/limits.conf

1. *soft nofile 65535 * hard nof ile 65535 

echo ” fs.file-max=419430”> / e tc/sysctl.conf修改系统限制



CHANGE MASTER 'master215' TO MASTER_HOST='10.30.222.215',MASTER_USER='slave',MASTER_PASSWORD='zODfxPogU8hbmCTyF00c',MASTER_LOG_FILE='mysql-bin.000279',MASTER_LOG_POS=692183009;

CHANGE MASTER 'master241'  TO MASTER_HOST='10.30.222.241',MASTER_USER='slave',MASTER_PASSWORD='zODfxPogU8hbmCTyF00c',MASTER_LOG_FILE='mysql-bin.000271',MASTER_LOG_POS=705350361;

CHANGE MASTER  'master122'  TO MASTER_HOST='10.135.255.122',MASTER_USER='slave',MASTER_PASSWORD='zODfxPogU8hbmCTyF00c',MASTER_LOG_FILE='mysql-bin.000268',MASTER_LOG_POS=685747206;

CHANGE MASTER 'master195'  TO MASTER_HOST='10.30.222.195',MASTER_USER='slave',MASTER_PASSWORD='zODfxPogU8hbmCTyF00c',MASTER_LOG_FILE='mysql-bin.000270',MASTER_LOG_POS=679772092;





修改密码：echo "密码" | passwd --stdin 用户名

docker-compose up -d 启动容器，如果镜像不存在则先下载镜像，如果容器没创建则创建容器，如果容器没启动则启动
docker-compose down 停止并移除容器
docker-compose restart 重启服务

acl权限

getfacl /test

setfacl -x u:code /test     #取消acl权限

setfacl -x m /test                 //恢复有效权限

setfacl -m -R u:code:rwx /test  设置acl权限

chatter: 锁定文件，不能删除，不能更改
​        +a:  只能给文件添加内容，但是删除不了，
​              chattr +a  /etc/passwd
​        -d:      不可删除
​        加锁：chattr +i  /etc/passwd       文件不能删除，不能更改，不能移动
​        查看加锁： lsattr /etc/passwd      文件加了一个参数 i 表示锁定
​        解锁：chattr -i /home/omd/h.txt    - 表示解除

随机生成：

```
openssl rand -base64 40
```

后台运行进程方法：

安装：yum -y install screen

- 用`screen -dmS *session name*`来建立一个处于断开模式下的会话（并指定其会话名）。
- 用`screen -list `来列出所有会话。
- 用`screen -r *session name*`来重新连接指定会话。
- 用快捷键`CTRL-a d `来暂时断开当前会话。

Linux pstree命令将所有行程以树状图显示，树状图将会以 pid (如果有指定) 或是以 init 这个基本行程为根 (root)，如果有指定使用者 id，则树状图会只显示该使用者所拥有的行程。

sh命令

-n 对脚本进行语法检查

-v

-x跟踪脚本执行过程



expect自动化交互（自动输入）

rpm -qa expect

yum -y install expect

```bash
#!/usr/bin/expect
set file [lindex $argv 0]
spawn ssh root@120.78.200.229 uptime   #开始自动交互所使用命令
expect {
    "yes/no" {exp_send "yes\r";exp_continue}  #多次匹配
    "*password:" {exp_send "Sibu@2019..\r"}
}
expect eof  #终结expect
```





shell脚本编写技巧

set -e   脚本有错误就退出

for i in {100..1};do

   sleep 1

done          #一个倒计时语句



**top 运行中可以通过 top 的内部命令对进程的显示方式进行控制。内部命令如下：**
s – 改变画面更新频率
l – 关闭或开启第一部分第一行 top 信息的表示
t – 关闭或开启第一部分第二行 Tasks 和第三行 Cpus 信息的表示
m – 关闭或开启第一部分第四行 Mem 和 第五行 Swap 信息的表示
N – 以 PID 的大小的顺序排列表示进程列表
P – 以 CPU 占用率大小的顺序排列进程列表
M – 以内存占用率大小的顺序排列进程列表
h – 显示帮助
n – 设置在进程列表所显示进程的数量
q – 退出 top
s – 改变画面更新周期

H-以 CPU 占用率大小的顺序排列进程列表

top -Hp pid

jvm分析：http://www.importnew.com/28916.html

jstack -l pid

jstack -m pid

jstack -F pid

jmap

\#提取进程内存信息，用于分析OOM导致原因
jmap -dump:format=b,file=HeapDump.bin <pid>
\#输出堆信息
jmap -heap <PID> 

jhat简单分析内存中对象情况

\#读取dump文件，生成报告，并启动WEB服务器，默认端口为7000
jhat -J-mx768m -stack false HeapDump.bin 

jstat

jstat -gcutil <pid> 2000 100 # 每2秒输出一次内存情况，连续输出100次

jstat -gc<pid> 输出heap各个分区大小

- S0C：第一个幸存区的大小
- S1C：第二个幸存区的大小
- S0U：第一个幸存区的使用大小
- S1U：第二个幸存区的使用大小
- EC：伊甸园区的大小
- EU：伊甸园区的使用大小
- OC：老年代大小
- OU：老年代使用大小
- MC：方法区大小
- MU：方法区使用大小
- CCSC:压缩类空间大小
- CCSU:压缩类空间使用大小
- YGC：年轻代垃圾回收次数
- YGCT：年轻代垃圾回收消耗时间
- FGC：老年代垃圾回收次数
- FGCT：老年代垃圾回收消耗时间
- GCT：垃圾回收消耗总时间

jstat -gccapacity pid 250 20

- NGCMN：新生代最小容量
- NGCMX：新生代最大容量
- NGC：当前新生代容量
- S0C：第一个幸存区大小
- S1C：第二个幸存区的大小
- EC：伊甸园区的大小
- OGCMN：老年代最小容量
- OGCMX：老年代最大容量
- OGC：当前老年代大小
- OC:当前老年代大小
- MCMN:最小元数据容量
- MCMX：最大元数据容量
- MC：当前元数据空间大小
- CCSMN：最小压缩类空间大小
- CCSMX：最大压缩类空间大小
- CCSC：当前压缩类空间大小
- YGC：年轻代gc次数
- FGC：老年代GC次数

查看运行时jvm参数 

jinfo -flag <jvm参数> <pid>
举例：
jinfo -flag MaxHeapSize 107249



获取当前JVM默认参数 

java -XX:+PrintFlagsFinal -version | grep MaxHeapSize