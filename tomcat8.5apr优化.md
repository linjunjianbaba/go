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