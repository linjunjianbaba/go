#!/bin/bash
IP=`cat /root/ip |awk '{print $2}'`
#j=`cat /root/ip |grep $IP |awk '{print $1}'`
for i in $IP
do
     j=`cat /root/ip |grep $i |awk '{print $1}'`
     mysqldump -uroot -p'zWEGQq7GM8QjlIIJfyCt' -h $i --set-gtid-purged=OFF vmall_order$j | mysql -u'sibu_develop_write' -p'LU8%**DZYbZ7*c85' -h $i mall_order$j
done

