# 浅谈tomcat优化（内存，并发，缓存，安全，网络，系统等）

# 一.Tomcat内存优化

1. Tomcat内存优化主要是对 tomcat 启动参数优化，我们可以在 tomcat 的启动脚本 catalina.sh 中设置 java_OPTS 参数
2. JAVA_OPTS参数说明
   　　-server 启用jdk 的 server 版
      　　-Xms java虚拟机初始化时的最小内存
      　　-Xmx java虚拟机可使用的最大内存
      　　-XX: PermSize 内存永久保留区域
      　　-XX:MaxPermSize 内存最大永久保留区域

3.配置示例：
JAVA_OPTS=’-Xms1024m -Xmx2048m -XX: PermSize=256M -XX:MaxNewSize=256m -XX:MaxPermSize=256m’
说明：其内存的配置需要根据服务器（或虚拟机）的实际内存来配置

4.重启tomcat生效

------

# 二.Tomcat并发优化

1. 调整连接器connector的并发处理能力：

   ```
       maxThreads ：客户请求最大线程数
   ```

   　　minSpareThreads ：Tomcat初始化时创建的 socket 线程数
   　　maxSpareThreads： Tomcat连接器的最大空闲 socket 线程数
   　　enableLookups ：是否反查域名，取值为： true 或 false 。为了提高处理能力，应设置为 false
   　　redirectPort： 在需要基于安全通道的场合，把客户请求转发到基于SSL 的 redirectPort 端口
   　　acceptAccount： 监听端口队列最大数，满了之后客户请求会被拒绝（不能小于maxSpareThreads ）
   　　connectionTimeout： 连接超时
   　　minProcessors： 服务器创建时的最小处理线程数
   　　maxProcessors： 服务器同时最大处理线程数
   　　URIEncoding： URL统一编码

   ```
            其中和最大连接数相关的参数为maxProcessors 和 acceptCount 。如果要加大并发连接数，应同时加大这两个参数。
   ```

   2、压缩优化及参数

   ```
     ●compression="on"   打开压缩功能
   ```

   ●compressionMinSize="2048"启用压缩的输出内容大小，默认为2KB
   ●noCompressionUserAgents="gozilla,traviata" 对于以下的浏览器，不启用压缩
   ●compressableMimeType="text/html,text/xml,text/javascript,text/css,text/plain"　哪些资源类型需要压缩

   ```
     Tomcat 的压缩是在客户端请求服务器对应资源后，从服务器端将资源文件压缩，再输出到客户端，由客户端的浏览器负责解压缩并浏览。相对于普通的浏览过程 HTML、CSS、Javascript和Text，它可以节省40% 左右的流量。更为重要的是，它可以对动态生成的，包括CGI、PHP、JSP、ASP、Servlet,SHTML等输出的网页也能进行压缩，压缩效率也很高。但是， 压缩会增加 Tomcat 的负担，因此最好采用Nginx + Tomcat 或者 Apache + Tomcat 方式，将压缩的任务交由 Nginx/Apache 去做。
   ```

------

------

------

# 三、Tomcat缓存优化

1、tomcat的maxThreads、acceptCount（最大线程数、最大排队数）
说明：
maxThreads：tomcat起动的最大线程数，即同时处理的任务个数，默认值为200

```
  acceptCount：当tomcat起动的线程数达到最大时，接受排队的请求个数，默认值为100
```

这两个值如何起作用，请看下面三种情况

情况1：接受一个请求，此时tomcat起动的线程数没有到达maxThreads，tomcat会起动一个线程来处理此请求。

情况2：接受一个请求，此时tomcat起动的线程数已经到达maxThreads，tomcat会把此请求放入等待队列，等待空闲线程。

情况3：接受一个请求，此时tomcat起动的线程数已经到达maxThreads，等待队列中的请求个数也达到了acceptCount，此时tomcat会直接拒绝此次请求，返回connection refused

maxThreads如何配置

一般的服务器操作都包括量方面：1计算（主要消耗cpu），2等待（io、数据库等）

第一种极端情况，如果我们的操作是纯粹的计算，那么系统响应时间的主要限制就是cpu的运算能力，此时maxThreads应该尽量设的小，降低同一时间内争抢cpu的线程个数，可以提高计算效率，提高系统的整体处理能力。

第二种极端情况，如果我们的操作纯粹是IO或者数据库，那么响应时间的主要限制就变为等待外部资源，此时maxThreads应该尽量设的大，这样才能提高同时处理请求的个数，从而提高系统整体的处理能力。此情况下因为tomcat同时处理的请求量会比较大，所以需要关注一下tomcat的虚拟机内存设置和linux的open  file限制。

我在测试时遇到一个问题，maxThreads我设置的比较大比如3000，当服务的线程数大到一定程度时，一般是2000出头，单次请求的响应时间就会急剧的增加，

百思不得其解这是为什么，四处寻求答案无果，最后我总结的原因可能是cpu在线程切换时消耗的时间随着线程数量的增加越来越大，

cpu把大多数时间都用来在这2000多个线程直接切换上了，当然cpu就没有时间来处理我们的程序了。

以前一直简单的认为多线程=高效率。。其实多线程本身并不能提高cpu效率，线程过多反而会降低cpu效率。

当cpu核心数<线程数时，cpu就需要在多个线程直接来回切换，以保证每个线程都会获得cpu时间，即通常我们说的并发执行。

所以maxThreads的配置绝对不是越大越好。

现实应用中，我们的操作都会包含以上两种类型（计算、等待），所以maxThreads的配置并没有一个最优值，一定要根据具体情况来配置。

最好的做法是：在不断测试的基础上，不断调整、优化，才能得到最合理的配置。

acceptCount的配置，我一般是设置的跟maxThreads一样大，这个值应该是主要根据应用的访问峰值与平均值来权衡配置的。

如果设的较小，可以保证接受的请求较快相应，但是超出的请求可能就直接被拒绝

如果设的较大，可能就会出现大量的请求超时的情况，因为我们系统的处理能力是一定的。

maxThreads 配置要结合 JVM -Xmx 参数调整，也就是要考虑内存开销。

------

------

------

# 四、tomcat的协议类型优化：

1、关闭AJP端口
AJP是为 Tomcat 与 HTTP 服务器之间通信而定制的协议，能提供较高的通信速度和效率。如果tomcat前端放的是apache的时候，会使用到AJP这个连接器。若tomcat未与apache配合使用，因此不使用此连接器，因此需要注销掉该连接器。
<!-- <Connector port="8009" protocol="AJP/1.3" redirectPort="8443" /> -->
2、bio模式：
默认的模式,性能非常低下,没有经过任何优化处理和支持.
3、nio模式：
01、nio(new  I/O)，是Java SE 1.4及后续版本提供的一种新的I/O操作方式(即java.nio包及其子包)。Java  nio是一个基于缓冲区、并能提供非阻塞I/O操作的Java API，因此nio也被看成是non-blocking  I/O的缩写。它拥有比传统I/O操作(bio)更好的并发运行性能。
02、如何启动此模式：
修改server.xml里的Connector节点,修改protocol为org.apache.coyote.http11.Http11NioProtocol
4、apr模式：
apr是从操作系统级别解决异步IO问题，大幅度提高服务器的并发处理性能，也是Tomcat生产环境运行的首选方式
目前Tomcat 8.x默认情况下全部是运行在nio模式下，而apr的本质就是使用jni技术调用操作系统底层的IO接口，所以需要提前安装所需要的依赖，首先是需要安装openssl和apr，命令如下：

yum -y install openssl-devel
yum -y install apr-devel

安装之后，去tomcat官网下载native组件，native可以看成是tomcat和apr交互的中间环节，下载地址是：<http://tomcat.apache.org/download-native.cgi> 这里下载最新的版本1.2.10

　　解压之后上传至服务器执行解压并安装：

tar -xvzf tomcat-native-1.2.10-src.tar.gz -C /usr/local
cd /usr/local/tomcat-native-1.2.10-src/native/
./configure 编译安装

然后进入tomcat安装目录，编辑配置文件：conf/server.xml

　　![img](http://i2.51cto.com/images/blog/201801/18/283507e8796d3ee520b094acef1dd4f8.png?x-oss-process=image/watermark,size_16,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_100,g_se,x_10,y_10,shadow_90,type_ZmFuZ3poZW5naGVpdGk=)

　　如图所示，将默认的protocol="HTTP/1.1"修改为protocol="org.apache.coyote.http11.Http11AprProtocol"

　　apr引入方法：

　　配置tomcat安装目录下:bin/catalina.sh文件引入apr，推荐这种方式：

　　![img](http://i2.51cto.com/images/blog/201801/18/91dd3c643c338a7dcc2db3586483dce6.png?x-oss-process=image/watermark,size_16,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_100,g_se,x_10,y_10,shadow_90,type_ZmFuZ3poZW5naGVpdGk=)

　　如图所示在原有变量JAVA_OPTS后面追加对应的配置即可，添加一行新的就可以：JAVA_OPTS="$JAVA_OPTS -Djava.library.path=/usr/local/apr/lib"

　　然后保存并退出

------

------

------

5、系统参数优化：
优化网卡驱动可以有效提升性能，这个对于集群环境工作的时候尤为重要。由于我们采用了linux服务器，所以优化内核参数也是一个非常重要的工作。给一个参考的优化参数：
01、 修改/etc/sysctl.cnf文件，在最后追加如下内容：
net.core.netdev_max_backlog = 32768
net.core.somaxconn = 32768
net.core.wmem_default = 8388608
net.core.rmem_default = 8388608
net.core.rmem_max = 16777216
net.core.wmem_max = 16777216
net.ipv4.ip_local_port_range = 1024 65000

net.ipv4.route.gc_timeout = 100
net.ipv4.tcp_fin_timeout = 30
net.ipv4.tcp_keepalive_time = 1200
net.ipv4.tcp_timestamps = 0
net.ipv4.tcp_synack_retries = 2
net.ipv4.tcp_syn_retries = 2
net.ipv4.tcp_tw_recycle = 1
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_mem = 94500000 915000000 927000000 

net.ipv4.tcp_max_orphans = 3276800 

net.ipv4.tcp_max_syn_backlog = 65536

02、 保存退出，执行sysctl -p生效

------

------

------

　　

# 五、tomcat的安全配置：

```
1、当Tomcat完成安装后你首先要做的事情如下：
```

首次安装完成后立即删除webapps下面的所有代码
rm -rf /srv/apache-tomcat/webapps/*
注释或删除 tomcat-users.xml 所有用户权限，看上去如下：

# cat conf/tomcat-users.xml

<?xml version='1.0' encoding='utf-8'?>

<tomcat-users>
</tomcat-users>

2、隐藏tomcat版本
01.首先找到这个jar包，$TOMCAT_HOME/lib/catalina.jar
02.解压catalina.jar之后按照路径\org\apache\catalina\util\ServerInfo.properties找到文件
03.打开ServerInfo.properties文件修改如下：把server.number、server.built置空
server.info=Apache Tomcat
server.number=
server.built=
04.重新打成jar包，重启tomcat。
3、隐藏tomcat 的服务类型
conf/server.xml文件中，为connector元素添加server="
"，注意不是空字符串，是空格组成的长度为1的字符串，或者输入其他的服务类型，这时候，在response header中就没有server的信息啦！
4、应用程序安全
关闭war自动部署 unpackWARs="false" autoDeploy="false"。防止被植入木马等恶意程序
5、修改服务监听端口
一般公司的 Tomcat 都是放在内网的，因此我们针对 Tomcat 服务的监听地址都是内网地址。
修改实例：

<Connector port="8080" address="172.16.100.1" />