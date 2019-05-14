#!/bin/bash
#Date: 2018-11-14
#Author: Created by warren
#Blog: 
#Description: This scripts function is mysql backup
#Version:1.0
set -e
REVTE=0
RPOURL="https://www.percona.com/redir/downloads/percona-release/redhat/latest/percona-release-0.1-6.noarch.rpm"
rpo_install () {
    if [ -f /etc/yum.repos.d/percona-release.repo ];then
	    echo "percona-repo is install"
		REVTE=$?
	else
	    yum -y install $RPOURL && echo "$HOSTNAME install percona-repo success!!"
		REVTE=$?
    fi    	
}
innobackup_install() {
    
    if [ $REVTE -eq 0 ];then
	    in_no=`rpm -qa |grep xtrabackup |wc -l`
	    if [ $in_no -eq 0 ];then
	         yum -y install percona-xtrabackup && echo "$HOSTNAME install percona-xtrabackup success!!"
	    else
		     echo "percona-xtrabackup is install"  
	    fi
    fi 	
}
main(){
  rpo_install
  innobackup_install  
}
main