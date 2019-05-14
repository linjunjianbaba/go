#!/bin/bash
MYSQLUSER=root
MYSQLPASSWD=jnw5O6MvfI08OzmtkSGM
ALLIP=`cat mysql |awk '{print $2}'`
HOSTIP=`ifconfig eth0 |grep inet |awk '{print $2}'`
INNOCMD=/usr/bin/innobackupex
BACKPATH=/var/backup
SCPPATH=/tmp
DATA=`date +%s`
[ -d $BACKPATH ] && mv $BACKPATH/slave{,$DATA}
DATABASE1=`cat mysql |grep $HOSTIP |awk '{print $3}'`
DATABASE2=`cat mysql |grep $HOSTIP |awk '{print $4}'`

$INNOCMD --user=$MYSQLUSER --password=$MYSQLPASSWD --databases="mysql ${DATABASE1} ${DATABASE2}" --no-timestamp $BACKPATH/slave

scp -r $BACKPATH/slave 10.66.211.4:$SCPPATH
