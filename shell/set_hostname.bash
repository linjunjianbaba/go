#!/bin/bash
ip=`ifconfig eth0 |grep 'inet' |awk {print $1}`
hostname=`cat /root/host |grep $ip |awk {print $2}`
echo $ip
echo $hostname

hostnamectl set-hostname --static $hostname
hostname $hostname