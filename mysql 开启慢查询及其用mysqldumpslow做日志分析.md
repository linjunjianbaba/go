mysql 开启慢查询及其用mysqldumpslow做日志分析

mysql慢查询日志是mysql提供的一种日志记录，它是用来记录在mysql中相应时间超过阈值的语句，就是指运行时间超过long_query_time值的sql，会被记录在慢查询日志中。long_query_time的默认值是10，意思是运行10S之上的语句。

慢查询日志的设置

1   、查看是否开启慢查询日志命令：

```shell 
show variables like '%slow_query_log%'
```

2、设置慢查询开启的命令

```
set global slow_query_log=1
```

注： 
 slow_query_log ON为开启，OFF为关闭 
 slow_query_log_file 为慢查询日志的存放地址

3、查询并修改慢查询定义的时间

```
show variables like 'long_query_time%'
set global long_query_time=4
```

4、未使用索引的查询被记录到慢查询日志中。如果调优的话，建议开启这个选项。如果开启了这个参数，full index scan的sql也会被记录到慢查询日志中

``` shell
show variables like 'log_queries_not_using_indexes'
set global log_queries_not_using_indexes=1
```

5.查询有多少条慢查询记录

```
show global status like '%Slow_queries%';
```

mysqldumpslow 慢日志分析工具
命令：

```
-s 按照那种方式排序
    c：访问计数
    l：锁定时间
    r:返回记录
    al：平均锁定时间
    ar：平均访问记录数
    at：平均查询时间
-t 是top n的意思，返回多少条数据。
-g 可以跟上正则匹配模式，大小写不敏感。
```

得到返回记录最多的20个sql

```
mysqldumpslow -s r -t 20 sqlslow.log
```

得到平均访问次数最多的20条sql

```
mysqldumpslow -s ar -t 20 sqlslow.log
```

得到平均访问次数最多,并且里面含有ttt字符的20条sql

```
mysqldumpslow -s ar -t 20 -g "ttt" sqldlow.log
```

注： 
 1、如果出现 -bash: mysqldumpslow: command not found 错误，请执行

```
ln -s /usr/local/mysql/bin/mysqldumpslow /usr/bin
```

2、如果出现如下错误，Died at /usr/bin/mysqldumpslow line 161, <> chunk 405659.说明你要分析的sql日志太大了，请拆分后再分析

拆分的命令为：

```
tail -100000 mysql-slow.log>mysql-slow.20180725.log
```

