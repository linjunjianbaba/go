#!/bin/bash
yum -y install expect

innobackupex --user='root' --password='jnw5O6MvfI08OzmtkSGM' --databases='mysql sibu_wesale_buyer_order_07 sibu_wesale_buyer_order_08' --no-timestamp /var/backup/slave
/usr/bin/expect <<EOF
    spawn scp -r /var/backup/slave 10.66.211.4:/tmp
    set timeout 60
    expect {
           "*(yes/no)?" {send "yes\r"}
           "*assword:" {send "sibu@123..\r"}
    }
    expect eof
EOF

#!/bin/bash

IP=10.135.255.96

CHANGE MASTER 'master255147' TO MASTER_HOST='10.135.255.147',MASTER_USER='slave',MASTER_PASSWORD='zODfxPogU8hbmCTyF00c',MASTER_LOG_FILE='mysql-bin.000293',MASTER_LOG_POS=506230391;