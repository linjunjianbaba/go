V商城docker

docker build -t vshop:v1 /dockerfile/java/

docker run -d --name vshop -v /dockerfile/java:/home/app --log-driver=fluentd --log-opt tag="docker.{{.Name}}" --log-opt fluentd-async-connect=true -p 82:8080 vshop:v1



java -server -Xms8G -Xmx8G -Xmn4G -Xss256k -XX:+PrintGCDetails -XX:+PrintGCTimeStamps -Xloggc:$WORK_DATA/logs/gc_%p.log -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=5 -XX:GCLogFileSize=30m -XX:+HeapDumpOnOutOfMemoryError -jar -Djava.security.egd=file:/dev/./urandom



dobbo-admin启动：nohup java -jar dubbo-admin-0.0.1-SNAPSHOT.jar /dev/null 2>&1 &

/data/dubbo-monitor-simple-2.0.0/bin/start.sh