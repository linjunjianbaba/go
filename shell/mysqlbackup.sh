#!/bin/bash
#Date: 2018-11-14
#Author: Created by bill
#Blog: 
#Description: This scripts function is mysql backup
#Version:1.0
set -e
INNOBACK=/usr/bin/innobackupex
BAKPUPPATH=/var/backup/
MYSQLUSER=root
MYSQLPASSWD=sibu2018
DATA=`date +%s`

if [ !-d $BAKPUPPATH ];then
    /usr/bin/mkdir -p $BAKPUPPATH
fi

if [ -d $BAKPUPPATH/slave ];then
    mv $BAKPUPPATH/slave{,$DATA}
fi
mysql_backup() {
    innobackupex --user='$MYSQLUSER' --password='$MYSQLPASSWD' --no-timestamp $BAKPUPPATH/slave	
}




    
