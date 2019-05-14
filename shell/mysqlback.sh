#!/bin/bash
DBPATH=/backup
MYUSER=
MYPASS=
MYSOCK=/var/lib/mysql/mysql.sock
MYCMD="mysql -u$MYUSER -p$MYPASS -S $MYSOCK"
MYDUMP="mysqldump -u$MYUSER -p$MYPASS -S $MYSOCK"

[ ! -d "$DBPATH" ] && mkdir $DBPATH

for dbname in `$MYCMD -e "show databases;" |sed '1d' |egrep -v "mysql|schema"`
do
    $MYDUMP $dbname | gzip >$DBPATH/${dbname}_$(date +%F).sql.gz
done