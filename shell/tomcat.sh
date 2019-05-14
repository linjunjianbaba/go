#!/bin/bash
#JAVA_HOME=
#CATALINA_HOME=
TOMCATHOME=/var/tomcat
PORT=8080
PID=`/usr/sbin/lsof -n -P -i :$PORT`
RETEL=0
. /etc/init.d/functions
TOMCATUID=`id -u tomcat`
if [ ! -n "$TOMCATUID"  ];then
    useradd tomcat
fi
chown -R tomcat.tomcat /var/tomcat

start_tomcat() {
    if [ -n "$PID" ];then
        echo "tomcat is running."
    else
        if [ `id -u` -eq 0 ];then
            sudo su - tomcat -c "$TOMCATHOME/bin/startup.sh >/dev/null"
		    RETEL=$?
		    if [ $RETEL -eq 0 ];then
		        action $"Startting start tomcat" /bin/true
		    else
		        action $"Startting start tomcat" /bin/false
		    fi
        else
            $TOMCATHOME/bin/startup.sh >/dev/null 
        fi
    fi
}
stop_tomcat() {
    if [ -n "$PID" ];then
        $TOMCATHOME/bin/shutdown.sh >/dev/null
        sleep 1
        PID=`/usr/sbin/lsof -n -P -i :$PORT`
        if [ -n "$PID" ];then
            kill -9 $PID >/dev/null
        fi
        RETEL=$?
		if [ $RETEL -eq 0 ];then
		    action $"Shutting down tomcat" /bin/true
		else
		    action $"Shutting down tomcat" /bin/false
		fi

    fi
}
status_tomcat(){
    if [ -n "$PID" ];then
	    /usr/sbin/lsof -i:$PORT
    fi	
}
case $1 in
  start)
      start_tomcat
      ;;
  stop)
      stop_tomcat
      ;;
  restart)
      stop_tomcat
      sleep 1
      start_tomcat
      ;;
  status)
      status_tomcat
      #/var/tomcat/bin/catalina.sh version
      ;;
  *)
      echo "......................"
      ;;
esac


