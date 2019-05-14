mysql在线修改参数

在线修改innodb_buffer_pool

show variables like 'innodb_buffer_pool%';

SET GLOBAL innodb_buffer_pool_size=268435456;



# 查询最高内存占用

使用以下命令可以知道mysql的配置使用多少 RAM

```
SELECT ( @@key_buffer_size
+ @@query_cache_size
+ @@innodb_buffer_pool_size
+ @@innodb_additional_mem_pool_size
+ @@innodb_log_buffer_size
+ @@max_connections * ( @@read_buffer_size
+ @@read_rnd_buffer_size
+ @@sort_buffer_size
+ @@join_buffer_size
+ @@binlog_cache_size
+ @@thread_stack
+ @@tmp_table_size
)
) / (1024 * 1024 * 1024) AS MAX_MEMORY_GB;
```

可以使用[mysql计算器](http://www.mysqlcalculator.com/)来计算内存使用

下面是理论，可以直接到推荐配置

# 如何调整配置

## key_buffer_size（MyISAM索引用）

> 指定索引缓冲区的大小，它决定索引处理的速度，尤其是索引读的速度。为了最小化磁盘的 I/O ， MyISAM  存储引擎的表使用键高速缓存来缓存索引，这个键高速缓存的大小则通过 key-buffer-size 参数来设置。如果应用系统中使用的表以  MyISAM 存储引擎为主，则应该适当增加该参数的值，以便尽可能的缓存索引，提高访问的速度。

### 怎么设

```
show global status like 'key_read%';

+------------------------+-------------+
| Variable_name | Value |
+------------------------+-------------+
| Key_read_requests | 27813678764 |
| Key_reads | 6798830 |
--------------------- 
```

- key_buffer_size通过检查状态值Key_read_requests和Key_reads，可以知道key_buffer_size设置是否合理。
- 比例key_reads / key_read_requests应该尽可能的低，至少是1:100，1:1000更好。

```
show global status like '%created_tmp_disk_tables%';
```

- key_buffer_size只对MyISAM表起作用。即使你不使用MyISAM表，但是内部的临时磁盘表是MyISAM表，也要使用该值。可以使用检查状态值created_tmp_disk_tables得知详情。
- 对于1G内存的机器，如果不使用MyISAM表，推荐值是16M（8-64M）

**另一个参考如下**

```
show global status like 'key_blocks_u%';
+------------------------+-------------+
| Variable_name | Value |
+------------------------+-------------+
| Key_blocks_unused | 0 |
| Key_blocks_used | 413543 |
+------------------------+-------------+
```

Key_blocks_unused表示未使用的缓存簇(blocks)数，Key_blocks_used表示曾经用到的最大的blocks数，比如这台服务器，所有的缓存都用到了，要么增加key_buffer_size，要么就是过渡索引了，把缓存占满了。比较理想的设置：

- 可以根据此工式来动态的调整`Key_blocks_used / (Key_blocks_unused + Key_blocks_used) * 100% ≈ 80%`

```
show engines;
```

- 查询存储引擎

## innodb_buffer_pool_size （innodb索引用）

> 这个参数和MyISAM的`key_buffer_size`有相似之处，但也是有差别的。这个参数主要缓存innodb表的索引，数据，插入数据时的缓冲。为Innodb加速优化首要参数。　　

该参数分配内存的原则：这个参数默认分配只有8M，可以说是非常小的一个值。

- 如果是专用的DB服务器，且以InnoDB引擎为主的场景，通常可设置物理内存的50%，这个参数不能动态更改，所以分配需多考虑。分配过大，会使Swap占用过多，致使Mysql的查询特慢。
- 如果是非专用DB服务器，可以先尝试设置成内存的1/4，如果有问题再调整

## query_cache_size（查询缓存）

> 缓存机制简单的说就是缓存sql文本及查询结果，如果运行相同的sql，服务器直接从缓存中取到结果，而不需要再去解析和执行sql。如果表更改了，那么使用这个表的所有缓冲查询将不再有效，查询缓存值的相关条目被清空。更改指的是表中任何数据或是结构的改变，包括INSERT、UPDATE、DELETE、TRUNCATE、ALTER  TABLE、DROP TABLE或DROP  DATABASE等，也包括那些映射到改变了的表的使用MERGE表的查询。显然，这对于频繁更新的表，查询缓存是不适合的，而对于一些不常改变数据且有大量相同sql查询的表，查询缓存会节约很大的性能。

- 注意：如果你查询的表更新比较频繁，而且很少有相同的查询，最好不要使用查询缓存。因为这样会消耗很大的系统性能还没有任何的效果

### 要不要打开？

先设置成这样跑一段时间

```
query_cache_size=128M 
query_cache_type=1 
```

看看命中结果来进行进一步的判断

```
mysql> show status like '%Qcache%';
+-------------------------+-----------+
| Variable_name           | Value     |
+-------------------------+-----------+
| Qcache_free_blocks      | 669       |
| Qcache_free_memory      | 132519160 |
| Qcache_hits             | 1158      |
| Qcache_inserts          | 284824    |
| Qcache_lowmem_prunes    | 2741      |
| Qcache_not_cached       | 1755767   |
| Qcache_queries_in_cache | 579       |
| Qcache_total_blocks     | 1853      |
+-------------------------+-----------+
8 rows in set (0.00 sec)
```

> Qcache_free_blocks:表示查询缓存中目前还有多少剩余的blocks，如果该值显示较大，则说明查询缓存中的内存碎片过多了，可能在一定的时间进行整理。

Qcache_free_memory:查询缓存的内存大小，通过这个参数可以很清晰的知道当前系统的查询内存是否够用，是多了，还是不够用，DBA可以根据实际情况做出调整。

Qcache_hits:表示有多少次命中缓存。我们主要可以通过该值来验证我们的查询缓存的效果。数字越大，缓存效果越理想。

Qcache_inserts:  表示多少次未命中然后插入，意思是新来的SQL请求在缓存中未找到，不得不执行查询处理，执行查询处理后把结果insert到查询缓存中。这样的情况的次数，次数越多，表示查询缓存应用到的比较少，效果也就不理想。当然系统刚启动后，查询缓存是空的，这很正常。

Qcache_lowmem_prunes:该参数记录有多少条查询因为内存不足而被移除出查询缓存。通过这个值，用户可以适当的调整缓存大小。

Qcache_not_cached: 表示因为query_cache_type的设置而没有被缓存的查询数量。

Qcache_queries_in_cache:当前缓存中缓存的查询数量。

Qcache_total_blocks:当前缓存的block数量。

- 我们可以看到现网命中1158，未缓存的有1755767次，说明我们这个系统命中的太少了，表变动比较多，不什么开启这个功能涉及参数
- query_cache_limit：允许 Cache 的单条 Query 结果集的最大容量，默认是1MB，超过此参数设置的 Query 结果集将不会被 Cache
- query_cache_min_res_unit：设置 Query Cache 中每次分配内存的最小空间大小，也就是每个 Query 的 Cache 最小占用的内存空间大小
- query_cache_size：设置 Query Cache 所使用的内存大小，默认值为0，大小必须是1024的整数倍，如果不是整数倍，MySQL 会自动调整降低最小量以达到1024的倍数
- query_cache_type：控制 Query Cache  功能的开关，可以设置为0(OFF),1(ON)和2(DEMAND)三种，意义分别如下： 0(OFF)：关闭 Query Cache  功能，任何情况下都不会使用 Query Cache 1(ON)：开启 Query Cache 功能，但是当 SELECT 语句中使用的  SQL_NO_CACHE 提示后，将不使用Query Cache 2(DEMAND)：开启 Query Cache 功能，但是只有当  SELECT 语句中使用了 SQL_CACHE 提示后，才使用 Query Cache
- query_cache_wlock_invalidate：控制当有写锁定发生在表上的时刻是否先失效该表相关的 Query  Cache，如果设置为 1(TRUE)，则在写锁定的同时将失效该表相关的所有 Query  Cache，如果设置为0(FALSE)则在锁定时刻仍然允许读取该表相关的 Query Cache。

## innodb_additional_mem_pool_size（InnoDB内部目录大小）

InnoDB 字典信息缓存主要用来存放 InnoDB 存储引擎的字典信息以及一些 internal 的共享数据结构信息，也就是存放Innodb的内部目录，所以其大小也与系统中所使用的 InnoDB 存储引擎表的数量有较大关系。

这个值不用分配太大，通常设置16Ｍ够用了，默认8M，如果设置的内存大小不够，InnoDB 会自动申请更多的内存，并在 MySQL 的 Error Log 中记录警告信息。

## innodb_log_buffer_size （日志缓冲）

表示InnoDB写入到磁盘上的日志文件时使用的缓冲区的字节数，默认值为16M。一个大的日志缓冲区允许大量的事务在提交之前不用写日志到磁盘，所以如果有更新，插入或删除许多行的事务，则使日志缓冲区更大一些可以节省磁盘IO

通常最大设为64M足够

## max_connections (最大并发连接)

MySQL的max_connections参数用来设置最大连接（用户）数。每个连接MySQL的用户均算作一个连接，max_connections的默认值为100。

- 这个参数实际起作用的最大值（实际最大可连接数）为16384，即该参数最大值不能超过16384，即使超过也以16384为准；
- 增加max_connections参数的值，不会占用太多系统资源。系统资源（CPU、内存）的占用主要取决于查询的密度、效率等；
- 该参数设置过小的最明显特征是出现”Too many connections”错误

```
mysql> show variables like '%max_connect%';
+-----------------------+-------+
| Variable_name         | Value |
+-----------------------+-------+
| extra_max_connections | 1     |
| max_connect_errors    | 100   |
| max_connections       | 2048  |
+-----------------------+-------+
3 rows in set (0.00 sec)

mysql> show status like 'Threads%';
+-------------------+---------+
| Variable_name     | Value   |
+-------------------+---------+
| Threads_cached    | 0       |
| Threads_connected | 1       |
| Threads_created   | 9626717 |
| Threads_running   | 1       |
+-------------------+---------+
4 rows in set (0.00 sec)
```

可以看到此时的并发数也就是Threads_connected=1，还远远达不到2048

```
mysql> show variables like 'open_files_limit';
+------------------+-------+
| Variable_name    | Value |
+------------------+-------+
| open_files_limit | 65535 |
+------------------+-------+
1 row in set (0.00 sec)
```

max_connections 还取决于操作系统对单进程允许打开最大文件数的限制

也就是说如果操作系统限制单个进程最大可以打开100个文件

那么 max_connections 设置为200也没什么用

MySQL 的 open_files_limit 参数值是在MySQL启动时记录的操作系统对单进程打开最大文件数限制的值

可以使用 show variables like 'open_files_limit'; 查看 open_files_limit 值

```
ulimit -n
65535
```

或者直接在 Linux 下通过ulimit -n命令查看操作系统对单进程打开最大文件数限制 ( 默认为1024 )

# connection级内存参数(线程独享)

connection级参数，是在每个connection第一次需要使用这个buffer的时候，一次性分配设置的内存。

## 排序性能

mysql对于排序,使用了两个变量来控制sort_buffer_size和 max_length_for_sort_data, 不象oracle使用SGA控制. 这种方式的缺点是要单独控制,容易出现排序性能问题.

```
mysql> SHOW GLOBAL STATUS like '%sort%';
+---------------------------+--------+
| Variable_name             | Value  |
+---------------------------+--------+
| Sort_merge_passes         | 0      |
| Sort_priority_queue_sorts | 1409   |
| Sort_range                | 0      |
| Sort_rows                 | 843479 |
| Sort_scan                 | 13053  |
+---------------------------+--------+
5 rows in set (0.00 sec)
```

- 如果发现`Sort_merge_passes`的值比较大，你可以考虑增加`sort_buffer_size` 来加速ORDER BY 或者GROUP BY 操作,不能通过查询或者索引优化的。我们这为0，那就没必要设置那么大。

## 读取缓存

read_buffer_size = 128K(默认128K)为需要全表扫描的MYISAM数据表线程指定缓存

read_rnd_buffer_size = 4M：(默认256K)首先，该变量可以被任何存储引擎使用，当从一个已经排序的键值表中读取行时，会先从该缓冲区中获取而不再从磁盘上获取。

## 大事务binlog

```
mysql> show global status like 'binlog_cache%';
+-----------------------+----------+
| Variable_name         | Value    |
+-----------------------+----------+
| Binlog_cache_disk_use | 220840   |
| Binlog_cache_use      | 67604667 |
+-----------------------+----------+
2 rows in set (0.00 sec)
```

- Binlog_cache_disk_use表示因为我们binlog_cache_size设计的内存不足导致缓存二进制日志用到了临时文件的次数
- Binlog_cache_use 表示 用binlog_cache_size缓存的次数
- 当对应的Binlog_cache_disk_use 值比较大的时候 我们可以考虑适当的调高 binlog_cache_size 对应的值
- 如上图，现网是32K，我们加到64K

## join语句内存影响

如果应用中，很少出现join语句，则可以不用太在乎join_buffer_size参数的设置大小。

如果join语句不是很少的话，个人建议可以适当增大join_buffer_size到1MB左右，如果内存充足可以设置为2MB。

## 线程内存影响

Thread_stack：每个连接线程被创建时，MySQL给它分配的内存大小。当MySQL创建一个新的连接线程时，需要给它分配一定大小的内存堆栈空间，以便存放客户端的请求的Query及自身的各种状态和处理信息。

```
mysql> show status like '%threads%';
+-------------------------+---------+
| Variable_name           | Value   |
+-------------------------+---------+
| Delayed_insert_threads  | 0       |
| Slow_launch_threads     | 0       |
| Threadpool_idle_threads | 0       |
| Threadpool_threads      | 0       |
| Threads_cached          | 0       |
| Threads_connected       | 1       |
| Threads_created         | 9649301 |
| Threads_running         | 1       |
+-------------------------+---------+
8 rows in set (0.00 sec)

mysql> show status like 'connections';
+---------------+---------+
| Variable_name | Value   |
+---------------+---------+
| Connections   | 9649311 |
+---------------+---------+
1 row in set (0.00 sec)
```

如上：系统启动到现在共接受到客户端的连接9649311次，共创建了9649301个连接线程，当前有1个连接线程处于和客户端连接的状态。而在Thread Cache池中共缓存了0个连接线程(Threads_cached)。

Thread Cache 命中率：

```
Thread_Cache_Hit = (Connections - Threads_created) / Connections * 100%;
```

一般在系统稳定运行一段时间后，Thread Cache命中率应该保持在90%左右才算正常。

## 内存临时表

tmp_table_size 控制内存临时表的最大值,超过限值后就往硬盘写，写的位置由变量 tmpdir 决定

max_heap_table_size 用户可以创建的内存表(memory table)的大小.这个值用来计算内存表的最大行数值。

Order By 或者Group By操作多的话，加大这两个值，默认16M

```
mysql> show status like 'Created_tmp_%';
+-------------------------+-------+
| Variable_name           | Value |
+-------------------------+-------+
| Created_tmp_disk_tables | 0     |
| Created_tmp_files       | 626   |
| Created_tmp_tables      | 3     |
+-------------------------+-------+
3 rows in set (0.00 sec)
```

- 如上图，写入硬盘的为0，3次中间表，说明我们的默认值足够用了

# mariadb 推荐配置

- 注意这里只推荐innodb引擎
- 内存配置只关注有注释的行

```
[mysqld]
datadir=/var/lib/mysql
socket=/var/lib/mysql/mysql.sock
default-storage-engine=INNODB

character-set-server=utf8
collation-server=utf8_general_ci

user=mysql
symbolic-links=0

# global settings
table_cache=65535
table_definition_cache=65535

max_allowed_packet=4M
net_buffer_length=1M
bulk_insert_buffer_size=16M

query_cache_type=0              #是否使用查询缓冲,0关闭
query_cache_size=0              #0关闭，因为改表操作多，命中低，开启消耗cpu

# shared
key_buffer_size=8M              #保持8M MyISAM索引用
innodb_buffer_pool_size=4G              #DB专用mem*50%，非DB专用mem*15%到25%
myisam_sort_buffer_size=32M
max_heap_table_size=16M             #最大中间表大小
tmp_table_size=16M              #中间表大小

# per-thread
sort_buffer_size=256K               #加速排序缓存大小
read_buffer_size=128k               #为需要全表扫描的MYISAM数据表线程指定缓存
read_rnd_buffer_size=4M             #已排序的表读取时缓存，如果比较大内存就到6M
join_buffer_size=1M             #join语句多时加大，1-2M
thread_stack=256k               #线程空间，256K or 512K
binlog_cache_size=64K               #大事务binlog


# big-tables
innodb_file_per_table = 1
skip-external-locking
max_connections=2048                #最大连接数
skip-name-resolve

# slow_query_log
slow_query_log_file = /var/log/mysql-slow.log
long_query_time = 30
group_concat_max_len=65536

# according to tuning-primer.sh
thread_cache_size = 8
thread_concurrency = 16

# set variables
concurrent_insert=2
```

# 运行时修改

使用以下命令来修改变量

```
set global {要改的key} = {值}; （立即生效重启后失效）
set @@{要改的key} = {值}; （立即生效重启后失效）
set @@global.{要改的key} = {值}; （立即生效重启后失效）
```

试验

```
mysql> set @@global.innodb_buffer_pool_size=4294967296;
ERROR 1238 (HY000): Variable 'innodb_buffer_pool_size' is a read only variable
mysql> set @@global.thread_stack=262144;
ERROR 1238 (HY000): Variable 'thread_stack' is a read only variable
mysql> set @@global.binlog_cache_size=65536;
Query OK, 0 rows affected (0.00 sec)
mysql> set @@join_buffer_size=1048576;
Query OK, 0 rows affected (0.00 sec)
mysql> set @@read_rnd_buffer_size=4194304;
Query OK, 0 rows affected (0.00 sec)
mysql> set @@sort_buffer_size=262144;
Query OK, 0 rows affected (0.00 sec)
mysql> set @@read_buffer_size=131072;
Query OK, 0 rows affected (0.00 sec)
mysql> set global key_buffer_size=8388608;
Query OK, 0 rows affected (0.39 sec)
```

- 我们可以看到`innodb_buffer_pool_size`和`thread_stack`报错了，他们只能改配置文件，在运行时是只读的。 以下直接复制使用

```
set @@global.binlog_cache_size=65536;
set @@join_buffer_size=1048576;
set @@read_rnd_buffer_size=4194304;
set @@sort_buffer_size=262144;
set @@read_buffer_size=131072;
set global key_buffer_size=8388608;
```