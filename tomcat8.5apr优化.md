tomcat8.5优化

1.软件包下载

​      tomcat：https://tomcat.apache.org/download-80.cgi

​       APR的三个依赖 http://apr.apache.org/download.cgi

​       jdk：https://www.oracle.com/technetwork/java/javase/downloads

2.软件安装：

​     编译APR组件和Tomcat-native组件，

​      依赖安装：yum -y install cmake gcc expat-devel

​      安装apr，apr-iconv，apr-util

```shell
tar xf apr-1.6.5.tar.gz && tar xf apr-iconv-1.2.2.tar.gz && tar xf apr-util-1.6.1.tar.gz
cd apr-1.6.5
./configure --prefix=/usr/local/apr  
make  
make install 
cd ../apr-iconv-1.2.2
./configure --prefix=/usr/local/apr-iconv --with-apr=/usr/local/apr
make  
make install 
cd ../apr-util-1.6.1
./configure --prefix=/usr/local/apr-util --with-apr=/usr/local/apr --with-apr-iconv=/usr/local/apr-iconv/bin/apriconv 
make  
make install 
```

​     安装jdk

```shell
tar xf jdk-8u201-linux-x64.tar.gz && mv jdk1.8.0_201 /usr/local/java
vi /etc/profile
#Java Env
export JAVA_HOME=/usr/local/java
export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
export PATH=$PATH:$JAVA_HOME/bin
source /etc/profile
```

​     安装tomcat-native

```shell
tar xf apache-tomcat-8.5.38.tar.gz
tar xf apache-tomcat-8.5.38/bin/tomcat-native.tar.gz
cd apache-tomcat-8.5.38/bin/tomcat-native-1.2.21-src/native
./configure --with-apr=/usr/local/apr --with-java-home=/usr/local/java
make 4 
make install
```

​       修改环境变量

```shell
vim /etc/profile
# apr
export LD_LIBRARY_PATH=/usr/local/apr/lib
```

```xml
修改Tomcat 下 conf/server.xml protocol的值  HTTP/1.1为org.apache.coyote.http11.Http11AprProtocol
 <Connector port="8080" protocol="org.apache.coyote.http11.Http11AprProtocol"
               connectionTimeout="20000"
               redirectPort="8443" />
```

```xml
修改SSLEngine 为off
<Listener className="org.apache.catalina.core.AprLifecycleListener" SSLEngine="off" />
```



常见错误FAQ：


错误提示 ：architecture word width mismatch
解决方式： java版本问题，可能是32位java运行64位的lib.so,更换java版本即可


错误提示：安装apr-util时也遇到类似的问题,找不到expat.h
解决方式：yum install expat-devel

错误提示：.lifecycleEvent Failed to initialize the SSLEngine
解决方式：<Listener className="org.apache.catalina.core.AprLifecycleListener" SSLEngine="off" />  默认为on  改为off 即可

在tomcat/catalina.sh中加入下面的配置，内存要根据机器实际情况配置，如果配置内存太大了有可能机器很慢。

JAVA_OPTS="-server -Xms512m -Xmx512m -Xss1024K -XX:PermSize=64m -XX:MaxPermSize=128m"

JAVA_HOME=/root/jdk1.8.0_131

CATALINA_HOME=/root/apache-tomcat-8.5.13

JVM虚拟机的启动配置 bin/catalina.sh

然后在tomat/conf/server.xml配置如下，增加并发性能

maxThreads="500" minSpareThreads="100" enableLookups="false" URIEncoding="utf-8"

acceptCount="500" connectionTimeout="20000" disableUploadTimeout="ture" redirectPort="8443"/>

设定虚拟机的server启动方式，以及堆内存的初始分配大小，垃圾收集机制，线程最大堆栈配置数，新生代内存大小等等

a 、JVM Server模式与client模式启动，最主要的差别在于：-Server模式启动时，速度较慢，但是一旦运行起来后，性能将会有很大的提升。JVM如果不显式指定是-Server模式还是-client模式，JVM能够根据下列原则进行自动判断（适用于Java5版本或者Java以上版本）；

[JVM client模式和Server模式的区别](https://link.jianshu.com?t=http%3A%2F%2Fblog.csdn.net%2Ftang_123_%2Farticle%2Fdetails%2F6018219)

b、线程堆栈 -Xss 1024K 可以根据业务服务器的每次请求的大小来进行分配；

c、－xms -xmx  是 jvm占用最小和最大物理内存配置参数，一般讲两者配置一样大，这样就免去了内存不够用时申请内存的耗时；

d、-XX:PermSize=128M -XX:MaxPermSize=128m

从前人的各类文章上了解到jvm的垃圾回收机制，这里只是简单提一下， jvm的内存分为2大类型，一个是perm型，另一个是generation型。perm区域存放的是class这些静态信息，一般默认64m，如果你的项目很大，有可能一启动就报错，out of memory permsize什么的，另外如果用spring框架的话很多类是动态反射加载的，运行一段时间有可能出现此异常，这种情况，设置下permsize就可以了。

另外一个类型才是重点，应用的代码基本上在这个区域活动，new的类都会在这个区域，而且jvm决大部分工作都在这里搞了，这个区详细说很复杂，有空去看sun资料，这里也只大概提下：这个区包含新生代和老生代区域，所有new出来的会放置在新区域，而多次回收失败的一些一直被使用的实例则被转移到老生代区域，所以新生代区域活动是最频繁的。新生代内存不足时会促发一次 这个区的gc －－－－然后再到老生代的gc－－－最后才轮到full gc。full gc代价很高，应该尽量避免，尽量在newsize参数的这个区gc，一般配置 newsize分配到总内存1／4左右，－－－最终，如果full gc 还是内存不足，那就会引发out of memory 常见的那种。

-----------摘自[jvm 参数优化－－－笔记](https://link.jianshu.com?t=http%3A%2F%2Fblog.csdn.net%2Fljwhx2002%2Farticle%2Fdetails%2F5968848)

e、-XX:+UseParallelGC -XX:ParallelGCThreads=2 -XX:+UseAdaptiveSizePolicy

这几个参数，一般的应用没什么必要，UseParallelGC 并行回收，XX:ParallelGCThreads 并行回收线程数，只有配置了UseParallelGC有效。UseAdaptiveSizePolicy，让jvm根据情况动态适配参数，当然如果你指定了某些参数，jvm就不会对那些参数再去调整的，加这个参数只要是让我们考虑不全的其它参数能让jvm帮忙做微处理。 总之UseParallelGC目的是 加快jvm回收频率 。

关于垃圾回收更详细的文章请见:

[tomcat查看GC信息](https://link.jianshu.com?t=http%3A%2F%2Fblog.csdn.net%2Fjimmy1980%2Farticle%2Fdetails%2F4968308)

下图只是简单的配置了:

JAVA_OPTS="-server -Xms1536m -Xmx1536m -Xss1024K -XX:PermSize=128m -XX:MaxPermSize=256m"

Tomcat堆内存的垃圾回收情况，可以看到默认是当系统使用到了阀值后进行GC回收

![img](https:////upload-images.jianshu.io/upload_images/9369836-0f34ddd807b202ea.JPG?imageMogr2/auto-orient/strip%7CimageView2/2/w/576/format/webp)

JVM 优化

模型资料来源：[http://xmuzyq.iteye.com/blog/599750](https://link.jianshu.com?t=http%3A%2F%2Fxmuzyq.iteye.com%2Fblog%2F599750)

配比资料：[http://www.jianshu.com/p/d45e12241af4](https://www.jianshu.com/p/d45e12241af4)

Java 的内存模型分为：

Young，年轻代（易被 GC）。Young 区被划分为三部分，Eden 区和两个大小严格相同的 Survivor 区，其中 Survivor 区间中，某一时刻只有其中一个是被使用的，另外一个留做垃圾收集时复制对象用，在 Young 区间变满的时候，minor GC 就会将存活的对象移到空闲的Survivor 区间中，根据 JVM 的策略，在经过几次垃圾收集后，任然存活于 Survivor 的对象将被移动到 Tenured 区间。

Tenured，终身代。Tenured 区主要保存生命周期长的对象，一般是一些老的对象，当一些对象在 Young 复制转移一定的次数以后，对象就会被转移到 Tenured 区，一般如果系统中用了 application 级别的缓存，缓存中的对象往往会被转移到这一区间。

Perm，永久代。主要保存 class,method,filed 对象，这部门的空间一般不会溢出，除非一次性加载了很多的类，不过在涉及到热部署的应用服务器的时候，有时候会遇到 java.lang.OutOfMemoryError : PermGen space 的错误，造成这个错误的很大原因就有可能是每次都重新部署，但是重新部署后，类的 class 没有被卸载掉，这样就造成了大量的 class 对象保存在了 perm 中，这种情况下，一般重新启动应用服务器可以解决问题。

Linux 修改 /usr/program/tomcat7/bin/catalina.sh 文件，把下面信息添加到文件第一行。

如果服务器只运行一个 Tomcat

机子内存如果是 4G：

CATALINA_OPTS="-Dfile.encoding=UTF-8 -server -Xms2048m -Xmx2048m -Xmn1024m -XX:PermSize=256m -XX:MaxPermSize=512m -XX:SurvivorRatio=10 -XX:MaxTenuringThreshold=15 -XX:NewRatio=2 -XX:+DisableExplicitGC"

机子内存如果是 8G：

CATALINA_OPTS="-Dfile.encoding=UTF-8 -server -Xms4096m -Xmx4096m -Xmn2048m -XX:PermSize=256m -XX:MaxPermSize=512m -XX:SurvivorRatio=10 -XX:MaxTenuringThreshold=15 -XX:NewRatio=2 -XX:+DisableExplicitGC"

机子内存如果是 16G：

CATALINA_OPTS="-Dfile.encoding=UTF-8 -server -Xms8192m -Xmx8192m -Xmn4096m -XX:PermSize=256m -XX:MaxPermSize=512m -XX:SurvivorRatio=10 -XX:MaxTenuringThreshold=15 -XX:NewRatio=2 -XX:+DisableExplicitGC"

机子内存如果是 32G：

CATALINA_OPTS="-Dfile.encoding=UTF-8 -server -Xms16384m -Xmx16384m -Xmn8192m -XX:PermSize=256m -XX:MaxPermSize=512m -XX:SurvivorRatio=10 -XX:MaxTenuringThreshold=15 -XX:NewRatio=2 -XX:+DisableExplicitGC"

如果是 8G 开发机

-Xms2048m -Xmx2048m -XX:NewSize=512m -XX:MaxNewSize=1024m -XX:PermSize=256m -XX:MaxPermSize=512m

如果是 16G 开发机

-Xms4096m -Xmx4096m -XX:NewSize=1024m -XX:MaxNewSize=2048m -XX:PermSize=256m -XX:MaxPermSize=512m

参数说明：

-Dfile.encoding：默认文件编码 -server：表示这是应用于服务器的配置，JVM 内部会有特殊处理的 -Xmx1024m：设置JVM最大可用内存为1024MB -Xms1024m：设置JVM最小内存为1024m。此值可以设置与-Xmx相同，以避免每次垃圾回收完成后JVM重新分配内存。 -Xmn1024m：设置JVM新生代大小（JDK1.4之后版本）。一般-Xmn的大小是-Xms的1/2左右，不要设置的过大或过小，过大导致老年代变小，频繁Full GC，过小导致minor GC频繁。如果不设置-Xmn，可以采用-XX:NewRatio=2来设置，也是一样的效果 -XX:NewSize：设置新生代大小 -XX:MaxNewSize：设置最大的新生代大小 -XX:PermSize：设置永久代大小 -XX:MaxPermSize：设置最大永久代大小 -XX:NewRatio=4：设置年轻代（包括 Eden 和两个 Survivor 区）与终身代的比值（除去永久代）。设置为 4，则年轻代与终身代所占比值为 1：4，年轻代占整个堆栈的 1/5 -XX:MaxTenuringThreshold=10：设置垃圾最大年龄，默认为：15。如果设置为0 的话，则年轻代对象不经过 Survivor 区，直接进入年老代。对于年老代比较多的应用，可以提高效率。如果将此值设置为一个较大值，则年轻代对象会在 Survivor 区进行多次复制，这样可以增加对象再年轻代的存活时间，增加在年轻代即被回收的概论。需要注意的是，设置了 -XX:MaxTenuringThreshold，并不代表着，对象一定在年轻代存活15次才被晋升进入老年代，它只是一个最大值，事实上，存在一个动态计算机制，计算每次晋入老年代的阈值，取阈值和MaxTenuringThreshold中较小的一个为准。

 -XX:+DisableExplicitGC：这个将会忽略手动调用 GC 的代码使得 System.gc() 的调用就会变成一个空调用，完全不会触发任何 GC

Windows 修改 /tomcat7/bin/catalina.bat 文件，找到这一行：echo Using CATALINA_BASE: "%CATALINA_BASE%"，然后在其上面添加如下内容，此方法只对解压版的 Tomcat 有效果，对于安装版本的需要点击安装后任务栏上的那个 Tomcat 图标，打开配置中有一个JavaTab 的进行编辑。

set JAVA_OPTS=%JAVA_OPTS% -Dfile.encoding="UTF-8" -Dsun.jnu.encoding="UTF8" -Ddefault.client.encoding="UTF-8" -Duser.language=Zhset JAVA_OPTS=%JAVA_OPTS% -server -Xms4096m -Xmx4096m -Xmn2048m -XX:PermSize=256m -XX:MaxPermSize=512m -XX:SurvivorRatio=10 -XX:MaxTenuringThreshold=15 -XX:NewRatio=2 -XX:+DisableExplicitGC

Server.xml的Connection优化

Tomcat的Connector是Tomcat接收HTTP请求的关键模块，我们可以配置它来指定IO模式，以及处理通过这个Connector接受到的请求的处理线程数以及其它一些常用的HTTP策略。其主要配置参数如下：

**1.指定使用NIO模型来接受HTTP请求**

protocol="org.apache.coyote.http11.Http11NioProtocol" 指定使用NIO模型来接受HTTP请求。默认是BlockingIO，配置为protocol="HTTP/1.1"

acceptorThreadCount="2" 使用NIO模型时接收线程的数目

**2.指定使用线程池来处理HTTP请求**

首先要配置一个线程池来处理请求（与Connector是平级的，多个Connector可以使用同一个线程池来处理请求）

maxThreads="1000" minSpareThreads="50" maxIdleTime="600000"/>

executor="tomcatThreadPool" 指定使用的线程池

**3.指定BlockingIO模式下的处理线程数目**

maxThreads="150"//Tomcat使用线程来处理接收的每个请求。这个值表示Tomcat可创建的最大的线程数。默认值200。可以根据机器的时期性能和内存大小调整，一般可以在400-500。最大可以在800左右。

minSpareThreads="25"---Tomcat初始化时创建的线程数。默认值4。如果当前没有空闲线程，且没有超过maxThreads，一次性创建的空闲线程数量。Tomcat初始化时创建的线程数量也由此值设置。

maxSpareThreads="75"--一旦创建的线程超过这个值，Tomcat就会关闭不再需要的socket线程。默认值50。一旦创建的线程 超过此数值，Tomcat会关闭不再需要的线程。线程数可以大致上用 “同时在线人数*每秒用户操作次数*系统平均操作时间” 来计算。

acceptCount="100"----指定当所有可以使用的处理请求的线程数都被使用时，可以放到处理队列中的请求数，超过这个数的请求将不予处 理。默认值10。如果当前可用线程数为0，则将请求放入处理队列中。这个值限定了请求队列的大小，超过这个数值的请求将不予处理。

connectionTimeout="20000" --网络连接超时，默认值20000，单位：毫秒。设置为0表示永不超时，这样设置有隐患的。通常可设置为30000毫秒。

**4.其它常用设置**

maxHttpHeaderSize="8192" http请求头信息的最大程度，超过此长度的部分不予处理。一般8K。

URIEncoding="UTF-8" 指定Tomcat容器的URL编码格式。

disableUploadTimeout="true" 上传时是否使用超时机制

enableLookups="false"--是否反查域名，默认值为true。为了提高处理能力，应设置为false

compression="on"   打开压缩功能

compressionMinSize="10240" 启用压缩的输出内容大小，默认为2KB

noCompressionUserAgents="gozilla, traviata"   对于以下的浏览器，不启用压缩

compressableMimeType="text/html,text/xml,text/javascript,text/css,text/plain" 哪些资源类型需要压缩

**5.小结**

关于Tomcat的Nio和ThreadPool，本身的引入就提高了处理的复杂性，所以对于效率的提高有多少，需要实际验证一下。

**6.配置示例**

redirectPort="8443"

maxThreads="150"

minSpareThreads="25"

maxSpareThreads="75"

acceptCount="100"

connectionTimeout="20000"

protocol="HTTP/1.1"

maxHttpHeaderSize="8192"

URIEncoding="UTF-8"

disableUploadTimeout="true"

enableLookups="false"

compression="on"

compressionMinSize="10240"

noCompressionUserAgents="gozilla, traviata"

compressableMimeType="text/html,text/xml,text/javascript,text/css,text/plain">

...

管理AJP端口

AJP是为 Tomcat 与 HTTP 服务器之间通信而定制的协议，能提供较高的通信速度和效率。如果tomcat前端放的是apache的时候，会使用到AJP这个连接器。由于我们公司前端是由nginx做的反向代理，因此不使用此连接器，因此需要注销掉该连接器。

默认 Tomcat 是开启了对war包的热部署的。为了防止被植入木马等恶意程序，因此我们要关闭自动部署。





# Tomcat 调优及 JVM 参数优化

Tomcat 的缺省配置是不能稳定长期运行的，也就是不适合生产环境，它会死机，让你不断重新启动，甚至在午夜时分唤醒你。对于操作系统优化来说，是尽可能的增大可使用的内存容量、提高CPU 的频率，保证文件系统的读写速率等。经过压力测试验证，在并发连接很多的情况下，CPU 的处理能力越强，系统运行速度越快。

![Tomcat 7.png](http://blog.chopmoon.com/ueditor/php/upload/image/20151031/1446271963871899.png)

Tomcat 的优化不像其它软件那样，简简单单的修改几个参数就可以了，它的优化主要有三方面，分为系统优化，Tomcat 本身的优化，Java 虚拟机（JVM）调优。系统优化就不在介绍了，接下来就详细的介绍一下 Tomcat 本身与 JVM 优化，以 Tomcat 7 为例。

**一、Tomcat 本身优化**

Tomcat 的自身参数的优化，这块很像 ApacheHttp Server。修改一下 xml 配置文件中的参数，调整最大连接数，超时等。此外，我们安装 Tomcat 是，优化就已经开始了。

1、工作方式选择

为了提升性能，首先就要对代码进行动静分离，让 Tomcat 只负责 jsp 文件的解析工作。如采用 Apache 和 Tomcat 的整合方式，他们之间的连接方案有三种选择，JK、http_proxy 和 ajp_proxy。相对于 JK 的连接方式，后两种在配置上比较简单的，灵活性方面也一点都不逊色。但就稳定性而言不像JK 这样久经考验，所以建议采用 JK 的连接方式。 

2、Connector 连接器的配置

之前文件介绍过的 Tomcat 连接器的三种方式： bio、nio 和 apr，三种方式性能差别很大，apr 的性能最优， bio 的性能最差。而 Tomcat 7 使用的 Connector  默认就启用的 Apr 协议，但需要系统安装 Apr 库，否则就会使用 bio 方式。

3、配置文件优化

配置文件优化其实就是对 server.xml 优化，可以提大大提高 Tomcat 的处理请求的能力，下面我们来看 Tomcat 容器内的优化。

默认配置下，Tomcat 会为每个连接器创建一个绑定的线程池（最大线程数 200），服务启动时，默认创建了 5 个空闲线程随时等待用户请求。

首先，打开 ${TOMCAT_HOME}/conf/server.xml，搜索【<Executor name="tomcatThreadPool"】，开启并调整为

```
`    ``<``Executor` `name``=``"tomcatThreadPool"` `namePrefix``=``"catalina-exec-"``        ``maxThreads``=``"500"` `minSpareThreads``=``"20"` `maxSpareThreads``=``"50"` `maxIdleTime``=``"60000"``/>`
```

注意， Tomcat 7 在开启线程池前，一定要安装好 Apr 库，并可以启用，否则会有错误报出，shutdown.sh 脚本无法关闭进程。

然后，修改<Connector …>节点，增加 executor 属性，搜索【port="8080"】，调整为

```
`    ``<``Connector` `executor``=``"tomcatThreadPool"``               ``port``=``"8080"` `protocol``=``"HTTP/1.1"``               ``URIEncoding``=``"UTF-8"``               ``connectionTimeout``=``"30000"``               ``enableLookups``=``"false"``               ``disableUploadTimeout``=``"false"``               ``connectionUploadTimeout``=``"150000"``               ``acceptCount``=``"300"``               ``keepAliveTimeout``=``"120000"``               ``maxKeepAliveRequests``=``"1"``               ``compression``=``"on"``               ``compressionMinSize``=``"2048"``               ``compressableMimeType``=``"text/html,text/xml,text/javascript,text/css,text/plain,image/gif,image/jpg,image/png"` `               ``redirectPort``=``"8443"` `/>`
```

maxThreads :Tomcat 使用线程来处理接收的每个请求，这个值表示 Tomcat 可创建的最大的线程数，默认值是 200

minSpareThreads：最小空闲线程数，Tomcat 启动时的初始化的线程数，表示即使没有人使用也开这么多空线程等待，默认值是 10。

maxSpareThreads：最大备用线程数，一旦创建的线程超过这个值，Tomcat 就会关闭不再需要的 socket 线程。

上边配置的参数，最大线程 500（一般服务器足以），要根据自己的实际情况合理设置，设置越大会耗费内存和 CPU，因为 CPU 疲于线程上下文切换，没有精力提供请求服务了，最小空闲线程数 20，线程最大空闲时间 60 秒，当然允许的最大线程连接数还受制于操作系统的内核参数设置，设置多大要根据自己的需求与环境。当然线程可以配置在“tomcatThreadPool”中，也可以直接配置在“Connector”中，但不可以重复配置。

URIEncoding：指定 Tomcat 容器的 URL 编码格式，语言编码格式这块倒不如其它 WEB 服务器软件配置方便，需要分别指定。

connnectionTimeout： 网络连接超时，单位：毫秒，设置为 0 表示永不超时，这样设置有隐患的。通常可设置为 30000 毫秒，可根据检测实际情况，适当修改。

enableLookups： 是否反查域名，以返回远程主机的主机名，取值为：true 或 false，如果设置为false，则直接返回IP地址，为了提高处理能力，应设置为 false。

disableUploadTimeout：上传时是否使用超时机制。

connectionUploadTimeout：上传超时时间，毕竟文件上传可能需要消耗更多的时间，这个根据你自己的业务需要自己调，以使Servlet有较长的时间来完成它的执行，需要与上一个参数一起配合使用才会生效。

acceptCount：指定当所有可以使用的处理请求的线程数都被使用时，可传入连接请求的最大队列长度，超过这个数的请求将不予处理，默认为100个。

keepAliveTimeout：长连接最大保持时间（毫秒），表示在下次请求过来之前，Tomcat 保持该连接多久，默认是使用 connectionTimeout 时间，-1 为不限制超时。

maxKeepAliveRequests：表示在服务器关闭之前，该连接最大支持的请求数。超过该请求数的连接也将被关闭，1表示禁用，-1表示不限制个数，默认100个，一般设置在100~200之间。

compression：是否对响应的数据进行 GZIP 压缩，off：表示禁止压缩；on：表示允许压缩（文本将被压缩）、force：表示所有情况下都进行压缩，默认值为off，压缩数据后可以有效的减少页面的大小，一般可以减小1/3左右，节省带宽。

compressionMinSize：表示压缩响应的最小值，只有当响应报文大小大于这个值的时候才会对报文进行压缩，如果开启了压缩功能，默认值就是2048。

compressableMimeType：压缩类型，指定对哪些类型的文件进行数据压缩。

noCompressionUserAgents="gozilla, traviata"： 对于以下的浏览器，不启用压缩。

如果已经对代码进行了动静分离，静态页面和图片等数据就不需要 Tomcat 处理了，那么也就不需要配置在 Tomcat 中配置压缩了。

以上是一些常用的配置参数属性，当然还有好多其它的参数设置，还可以继续深入的优化，HTTP Connector 与 AJP Connector 的参数属性值，可以参考官方文档的详细说明：

<https://tomcat.apache.org/tomcat-7.0-doc/config/http.html>

<https://tomcat.apache.org/tomcat-7.0-doc/config/ajp.html>

**二、JVM 优化**

 Tomcat 启动命令行中的优化参数，就是 JVM 的优化 。Tomcat 首先跑在 JVM 之上的，因为它的启动其实也只是一个 java 命令行，首先我们需要对这个 JAVA 的启动命令行进行调优。不管是 YGC 还是 Full GC，GC 过程中都会对导致程序运行中中断，正确的选择不同的 GC 策略，调整 JVM、GC 的参数，可以极大的减少由于 GC 工作，而导致的程序运行中断方面的问题，进而适当的提高 Java 程序的工作效率。但是调整 GC 是以个极为复杂的过程，由于各个程序具备不同的特点，如：web 和 GUI 程序就有很大区别（Web可以适当的停顿，但GUI停顿是客户无法接受的），而且由于跑在各个机器上的配置不同（主要 cup 个数，内存不同），所以使用的 GC 种类也会不同。

1、JVM 参数配置方法

Tomcat 的启动参数位于安装目录 ${JAVA_HOME}/bin目录下，Linux 操作系统就是 catalina.sh 文件。JAVA_OPTS，就是用来设置 JVM 相关运行参数的变量，还可以在 CATALINA_OPTS 变量中设置。关于这 2 个变量，还是多少有些区别的：

JAVA_OPTS：用于当 Java 运行时选项“start”、“stop”或“run”命令执行。

CATALINA_OPTS：用于当 Java 运行时选项“start”或“run”命令执行。

为什么有两个不同的变量？它们之间都有什么区别呢？

首先，在启动 Tomcat 时，任何指定变量的传递方式都是相同的，可以传递到执行“start”或“run”命令中，但只有设定在 JAVA_OPTS 变量里的参数被传递到“stop”命令中。对于 Tomcat 运行过程，可能没什么区别，影响的是结束程序，而不是启动程序。

第二个区别是更微妙，其他应用程序也可以使用 JAVA_OPTS 变量，但只有在 Tomcat 中使用 CATALINA_OPTS 变量。如果你设置环境变量为只使用 Tomcat，最好你会建议使用 CATALINA_OPTS 变量，而如果你设置环境变量使用其它的 Java 应用程序，例如 JBoss，你应该把你的设置放在JAVA_OPTS 变量中。

2、JVM 参数属性

32 位系统下 JVM 对内存的限制：不能突破 2GB ，那么这时你的 Tomcat 要优化，就要讲究点技巧了，而在 64 位操作系统上无论是系统内存还是 JVM 都没有受到 2GB 这样的限制。

针对于 JMX 远程监控也是在这里设置，以下为 64 位系统环境下的配置，内存加入的参数如下：

```
`CATALINA_OPTS="``-server ``-Xms6000M ``-Xmx6000M ``-Xss512k ``-XX:NewSize=2250M ``-XX:MaxNewSize=2250M ``-XX:PermSize=128M``-XX:MaxPermSize=256M  ``-XX:+AggressiveOpts ``-XX:+UseBiasedLocking ``-XX:+DisableExplicitGC ``-XX:+UseParNewGC ``-XX:+UseConcMarkSweepGC ``-XX:MaxTenuringThreshold=31 ``-XX:+CMSParallelRemarkEnabled ``-XX:+UseCMSCompactAtFullCollection ``-XX:LargePageSizeInBytes=128m ``-XX:+UseFastAccessorMethods ``-XX:+UseCMSInitiatingOccupancyOnly``-Duser.timezone=Asia``/Shanghai` `-Djava.awt.headless=``true``"`
```

为了看着方便，将每个参数单独写一行。上面参数好多啊，可能有人写到现在都没见过一个在 Tomcat 的启动命令里加了这么多参数，当然，这些参数只是我机器上的，不一定适合你，尤其是参数后的 value（值）是需要根据你自己的实际情况来设置的。

上述这样的配置，基本上可以达到：

系统响应时间增快；

JVM回收速度增快同时又不影响系统的响应率；

JVM内存最大化利用；

线程阻塞情况最小化。

JVM 常用参数详解：

-server：一定要作为第一个参数，在多个 CPU 时性能佳，还有一种叫 -client 的模式，特点是启动速度比较快，但运行时性能和内存管理效率不高，通常用于客户端应用程序或开发调试，在 32 位环境下直接运行 Java 程序默认启用该模式。Server 模式的特点是启动速度比较慢，但运行时性能和内存管理效率很高，适用于生产环境，在具有 64 位能力的 JDK 环境下默认启用该模式，可以不配置该参数。

-Xms：表示 Java 初始化堆的大小，-Xms 与-Xmx 设成一样的值，避免 JVM 反复重新申请内存，导致性能大起大落，默认值为物理内存的 1/64，默认（MinHeapFreeRatio参数可以调整）空余堆内存小于 40% 时，JVM 就会增大堆直到 -Xmx 的最大限制。

-Xmx：表示最大 Java 堆大小，当应用程序需要的内存超出堆的最大值时虚拟机就会提示内存溢出，并且导致应用服务崩溃，因此一般建议堆的最大值设置为可用内存的最大值的80%。如何知道我的 JVM 能够使用最大值，使用 java -Xmx512M -version 命令来进行测试，然后逐渐的增大 512 的值,如果执行正常就表示指定的内存大小可用，否则会打印错误信息，默认值为物理内存的 1/4，默认（MinHeapFreeRatio参数可以调整）空余堆内存大于 70% 时，JVM 会减少堆直到-Xms 的最小限制。

-Xss：表示每个 Java 线程堆栈大小，JDK 5.0 以后每个线程堆栈大小为 1M，以前每个线程堆栈大小为 256K。根据应用的线程所需内存大小进行调整，在相同物理内存下，减小这个值能生成更多的线程，但是操作系统对一个进程内的线程数还是有限制的，不能无限生成，经验值在 3000~5000 左右。一般小的应用， 如果栈不是很深， 应该是128k 够用的，大的应用建议使用 256k 或 512K，一般不易设置超过 1M，要不然容易出现out ofmemory。这个选项对性能影响比较大，需要严格的测试。

-XX:NewSize：设置新生代内存大小。

-XX:MaxNewSize：设置最大新生代新生代内存大小

-XX:PermSize：设置持久代内存大小

-XX:MaxPermSize：设置最大值持久代内存大小，永久代不属于堆内存，堆内存只包含新生代和老年代。

-XX:+AggressiveOpts：作用如其名（aggressive），启用这个参数，则每当 JDK 版本升级时，你的 JVM 都会使用最新加入的优化技术（如果有的话）。

-XX:+UseBiasedLocking：启用一个优化了的线程锁，我们知道在我们的appserver，每个http请求就是一个线程，有的请求短有的请求长，就会有请求排队的现象，甚至还会出现线程阻塞，这个优化了的线程锁使得你的appserver内对线程处理自动进行最优调配。

-XX:+DisableExplicitGC：在 程序代码中不允许有显示的调用“System.gc()”。每次在到操作结束时手动调用 System.gc() 一下，付出的代价就是系统响应时间严重降低，就和关于 Xms，Xmx 里的解释的原理一样，这样去调用 GC 导致系统的 JVM 大起大落。

-XX:+UseConcMarkSweepGC：设置年老代为并发收集，即 CMS gc，这一特性只有 jdk1.5
后续版本才具有的功能，它使用的是 gc 估算触发和 heap 占用触发。我们知道频频繁的 GC 会造面 JVM
的大起大落从而影响到系统的效率，因此使用了 CMS GC 后可以在 GC 次数增多的情况下，每次 GC 的响应时间却很短，比如说使用了 CMS
GC 后经过 jprofiler 的观察，GC 被触发次数非常多，而每次 GC 耗时仅为几毫秒。

-XX:+UseParNewGC：对新生代采用多线程并行回收，这样收得快，注意最新的 JVM 版本，当使用 -XX:+UseConcMarkSweepGC 时，-XX:UseParNewGC 会自动开启。因此，如果年轻代的并行 GC 不想开启，可以通过设置 -XX：-UseParNewGC 来关掉。

-XX:MaxTenuringThreshold：设置垃圾最大年龄。如果设置为0的话，则新生代对象不经过 Survivor 区，直接进入老年代。对于老年代比较多的应用（需要大量常驻内存的应用），可以提高效率。如果将此值设置为一 个较大值，则新生代对象会在 Survivor 区进行多次复制，这样可以增加对象在新生代的存活时间，增加在新生代即被回收的概率，减少Full GC的频率，这样做可以在某种程度上提高服务稳定性。该参数只有在串行 GC 时才有效，这个值的设置是根据本地的 jprofiler 监控后得到的一个理想的值，不能一概而论原搬照抄。

-XX:+CMSParallelRemarkEnabled：在使用 UseParNewGC 的情况下，尽量减少 mark 的时间。

-XX:+UseCMSCompactAtFullCollection：在使用 concurrent gc 的情况下，防止 memoryfragmention，对 live object 进行整理，使 memory 碎片减少。

-XX:LargePageSizeInBytes：指定 Java heap 的分页页面大小，内存页的大小不可设置过大， 会影响 Perm 的大小。

-XX:+UseFastAccessorMethods：使用 get，set 方法转成本地代码，原始类型的快速优化。

-XX:+UseCMSInitiatingOccupancyOnly：只有在 oldgeneration 在使用了初始化的比例后 concurrent collector 启动收集。

-Duser.timezone=Asia/Shanghai：设置用户所在时区。

-Djava.awt.headless=true：这个参数一般我们都是放在最后使用的，这全参数的作用是这样的，有时我们会在我们的 J2EE 工程中使用一些图表工具如：jfreechart，用于在 web 网页输出 GIF/JPG 等流，在 winodws 环境下，一般我们的 app server 在输出图形时不会碰到什么问题，但是在linux/unix 环境下经常会碰到一个 exception 导致你在 winodws 开发环境下图片显示的好好可是在 linux/unix 下却显示不出来，因此加上这个参数以免避这样的情况出现。

-Xmn：新生代的内存空间大小，注意：此处的大小是（eden+ 2 survivor space)。与 jmap -heap 中显示的 New gen 是不同的。整个堆大小 = 新生代大小 + 老生代大小 + 永久代大小。在保证堆大小不变的情况下，增大新生代后，将会减小老生代大小。此值对系统性能影响较大，Sun官方推荐配置为整个堆的 3/8。

-XX:CMSInitiatingOccupancyFraction：当堆满之后，并行收集器便开始进行垃圾收集，例如，当没有足够的空间来容纳新分配或提升的对象。对于 CMS 收集器，长时间等待是不可取的，因为在并发垃圾收集期间应用持续在运行（并且分配对象）。因此，为了在应用程序使用完内存之前完成垃圾收集周期，CMS 收集器要比并行收集器更先启动。因为不同的应用会有不同对象分配模式，JVM 会收集实际的对象分配（和释放）的运行时数据，并且分析这些数据，来决定什么时候启动一次 CMS 垃圾收集周期。这个参数设置有很大技巧，基本上满足(Xmx-Xmn)*(100-CMSInitiatingOccupancyFraction)/100 >= Xmn 就不会出现 promotion failed。例如在应用中 Xmx 是6000，Xmn 是 512，那么 Xmx-Xmn 是 5488M，也就是老年代有 5488M，CMSInitiatingOccupancyFraction=90 说明老年代到 90% 满的时候开始执行对老年代的并发垃圾回收（CMS），这时还 剩 10% 的空间是 5488*10% = 548M，所以即使 Xmn（也就是新生代共512M）里所有对象都搬到老年代里，548M 的空间也足够了，所以只要满足上面的公式，就不会出现垃圾回收时的 promotion failed，因此这个参数的设置必须与 Xmn 关联在一起。

-XX:+CMSIncrementalMode：该标志将开启 CMS 收集器的增量模式。增量模式经常暂停 CMS 过程，以便对应用程序线程作出完全的让步。因此，收集器将花更长的时间完成整个收集周期。因此，只有通过测试后发现正常 CMS 周期对应用程序线程干扰太大时，才应该使用增量模式。由于现代服务器有足够的处理器来适应并发的垃圾收集，所以这种情况发生得很少，用于但 CPU情况。

-XX:NewRatio：年轻代（包括 Eden 和两个 Survivor 区）与年老代的比值（除去持久代），-XX:NewRatio=4 表示年轻代与年老代所占比值为 1:4，年轻代占整个堆栈的 1/5，Xms=Xmx 并且设置了 Xmn 的情况下，该参数不需要进行设置。

-XX:SurvivorRatio：Eden 区与 Survivor 区的大小比值，设置为 8，表示 2 个 Survivor 区（JVM 堆内存年轻代中默认有 2 个大小相等的 Survivor 区）与 1 个 Eden 区的比值为 2:8，即 1 个 Survivor 区占整个年轻代大小的 1/10。

-XX:+UseSerialGC：设置串行收集器。

-XX:+UseParallelGC：设置为并行收集器。此配置仅对年轻代有效。即年轻代使用并行收集，而年老代仍使用串行收集。

-XX:+UseParallelOldGC：配置年老代垃圾收集方式为并行收集，JDK6.0 开始支持对年老代并行收集。

-XX:ConcGCThreads：早期 JVM 版本也叫-XX:ParallelCMSThreads，定义并发 CMS 过程运行时的线程数。比如 value=4 意味着 CMS 周期的所有阶段都以 4 个线程来执行。尽管更多的线程会加快并发 CMS 过程，但其也会带来额外的同步开销。因此，对于特定的应用程序，应该通过测试来判断增加 CMS 线程数是否真的能够带来性能的提升。如果还标志未设置，JVM 会根据并行收集器中的 -XX:ParallelGCThreads 参数的值来计算出默认的并行 CMS 线程数。

-XX:ParallelGCThreads：配置并行收集器的线程数，即：同时有多少个线程一起进行垃圾回收，此值建议配置与 CPU 数目相等。

-XX:OldSize：设置 JVM 启动分配的老年代内存大小，类似于新生代内存的初始大小 -XX:NewSize。

以上就是一些常用的配置参数，有些参数是可以被替代的，配置思路需要考虑的是 Java 提供的垃圾回收机制。虚拟机的堆大小决定了虚拟机花费在收集垃圾上的时间和频度。收集垃圾能够接受的速度和应用有关，应该通过分析实际的垃圾收集的时间和频率来调整。假如堆的大小很大，那么完全垃圾收集就会很慢，但是频度会降低。假如您把堆的大小和内存的需要一致，完全收集就很快，但是会更加频繁。调整堆大小的的目的是最小化垃圾收集的时间，以在特定的时间内最大化处理客户的请求。在基准测试的时候，为确保最好的性能，要把堆的大小设大，确保垃圾收集不在整个基准测试的过程中出现。

假如系统花费很多的时间收集垃圾，请减小堆大小。一次完全的垃圾收集应该不超过 3-5 秒。假如垃圾收集成为瓶颈，那么需要指定代的大小，检查垃圾收集的周详输出，研究垃圾收集参数对性能的影响。当增加处理器时，记得增加内存，因为分配能够并行进行，而垃圾收集不是并行的。

3、设置系统属性

之前说过，Tomcat 的语言编码，配置起来很慢，要经过多次设置才可以了，否则中文很有可能出现乱码情况。譬如汉字“中”，以 UTF-8 编码后得到的是 3 字节的值 %E4%B8%AD，然后通过 GET 或者 POST 方式把这 3 个字节提交到 Tomcat 容器，如果你不告诉 Tomcat 我的参数是用 UTF-8编码的，那么 Tomcat 就认为你是用 ISO-8859-1 来编码的，而 ISO8859-1（兼容 URI 中的标准字符集 US-ASCII）是兼容 ASCII 的单字节编码并且使用了单字节内的所有空间，因此 Tomcat 就以为你传递的用 ISO-8859-1 字符集编码过的 3 个字符，然后它就用 ISO-8859-1 来解码。

设置起来不难使用“ -D<名称>=<值> ”来设置系统属性：

-Djavax.servlet.request.encoding=UTF-8

-Djavax.servlet.response.encoding=UTF-8 

-Dfile.encoding=UTF-8 

-Duser.country=CN 

-Duser.language=zh

4、常见的 Java 内存溢出有以下三种

（1） java.lang.OutOfMemoryError: Java heap space —-JVM Heap（堆）溢出

JVM 在启动的时候会自动设置 JVM Heap 的值，其初始空间（即-Xms）是物理内存的1/64，最大空间（-Xmx）不可超过物理内存。可以利用 JVM提供的 -Xmn -Xms -Xmx 等选项可进行设置。Heap 的大小是 Young Generation 和 Tenured Generaion 之和。在 JVM 中如果 98％ 的时间是用于 GC，且可用的 Heap size 不足 2％ 的时候将抛出此异常信息。

解决方法：手动设置 JVM Heap（堆）的大小。  
（2） java.lang.OutOfMemoryError: PermGen space  —- PermGen space溢出。

PermGen space 的全称是 Permanent Generation space，是指内存的永久保存区域。为什么会内存溢出，这是由于这块内存主要是被 JVM 存放Class 和 Meta 信息的，Class 在被 Load 的时候被放入 PermGen space 区域，它和存放 Instance 的 Heap 区域不同，sun 的 GC 不会在主程序运行期对 PermGen space 进行清理，所以如果你的 APP 会载入很多 CLASS 的话，就很可能出现 PermGen space 溢出。

解决方法： 手动设置 MaxPermSize 大小

（3） java.lang.StackOverflowError   —- 栈溢出

栈溢出了，JVM 依然是采用栈式的虚拟机，这个和 C 与 Pascal 都是一样的。函数的调用过程都体现在堆栈和退栈上了。调用构造函数的 “层”太多了，以致于把栈区溢出了。通常来讲，一般栈区远远小于堆区的，因为函数调用过程往往不会多于上千层，而即便每个函数调用需要 1K 的空间（这个大约相当于在一个 C 函数内声明了 256 个 int 类型的变量），那么栈区也不过是需要 1MB 的空间。通常栈的大小是 1－2MB 的。
通常递归也不要递归的层次过多，很容易溢出。

解决方法：修改程序。

更多信息，请参考以下文章：

JVM 垃圾回收调优总结

<http://developer.51cto.com/art/201201/312639.htm>

JVM调优总结：典型配置举例

<http://developer.51cto.com/art/201201/311739.htm>

JVM基础：JVM参数设置、分析 

<http://developer.51cto.com/art/201201/312018.htm>

JVM 堆内存相关的启动参数：年轻代、老年代和永久代的内存分配

<http://www.2cto.com/kf/201409/334840.html>

Java 虚拟机–新生代与老年代GC

<http://my.oschina.net/sunnywu/blog/332870>

JVM（Java虚拟机）优化大全和案例实战

<http://blog.csdn.net/kthq/article/details/8618052>

JVM内存区域划分Eden Space、Survivor Space、Tenured Gen，Perm Gen解释 

<http://blog.chinaunix.net/xmlrpc.php?r=blog/article&uid=29632145&id=4616836>