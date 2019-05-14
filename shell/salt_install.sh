#!/bin/bash
MASTERSERVER=172.16.0.15
CONFFILE=/etc/salt/minion
RPMURL=https://repo.saltstack.com/yum/redhat/salt-repo-latest-2.el7.noarch.rpm
RETEL=0
HOSTNAME=`cat /etc/hostname`

if [ -f /etc/salt/minion ];then
    exit 0
fi

repo_install() {
    yum -y install $RPMURL
}
salt_minion() {
    yum -y install salt-minion
}
conffile_chage() {
    sed -i s/#master:\ salt/master:\ $MASTERSERVER/g $CONFFILE
    sed -i s/#id:/id:\ $HOSTNAME/g $CONFFILE
}
main() {
    if [ -f /etc/yum.repos.d/salt-latest.repo ];then
        RETEL=0
    else
        repo_install
        RETEL=$?
    fi
    if [ $RETEL -eq 0 ];then
        salt_minion
        RETEL=$?
        if [ $RETEL -eq 0 ];then
            conffile_chage
            RETEL=$?
            if [ $RETEL -eq 0 ];then
                systemctl enable salt-minion.service && systemctl start salt-minion.service
            fi
        else
            echo "plasse install salt-minion."
            exit 2
        fi
    else
        echo "plasse install install_repo."
        exit 1
    fi
}

main
rm -f $0


