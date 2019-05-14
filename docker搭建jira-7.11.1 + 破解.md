# docker搭建jira-7.11.1 + 破解

幾行命令教你搭建一個jira最新版。所有步驟必不可少。破解補丁需要的請在下面留言。

1. pull docker 鏡像：  jira:7.11.1(目前的最新版本)    mysql:5.7

   - ```
     docker pull cptactionhank/atlassian-jira-software
     docker pull mysql:5.7
     ```

2. 啟動mysql docker實例

   - ```
     docker run --name atlassian-mysql --restart always -p 3306:3306 -v /opt/mysql_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=passwd -d mysql:5.7   #記得修改passwd
     ```

3. 連接mysql

   - ```
     docker run -it --link atlassian-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
     ```

4. 創建jira數據庫,並添加jira用户

   - ```
     create database jira default character set utf8 collate utf8_bin;
     CREATE USER `jira`@`%` IDENTIFIED BY 'jira';GRANT ALL ON *.* TO `jira`@`%` WITH GRANT OPTION;
     alter user 'jira'@'%' identified with mysql_native_password by 'jira';
     ```

5. 修改mysql事物隔離級別

   - ```
     set global transaction isolation level read committed;
     set session transaction isolation level read committed;
     ```

6. 啟動jira實例

   - ```
     docker run --detach --restart always -v /data/atlassian/confluence:/home --publish 8080:8080 cptactionhank/atlassian-jira-software
     ```

7. 訪問：192.168.x.x:8080 進行jira配置。配置過程略。配置完成如下圖：

   - 

      

      

8. 破解

   - ```
     docker exec --user root 97 mv /opt/atlassian/jira/atlassian-jira/WEB-INF/lib/atlassian-extras-3.2.jar /opt/atlassian/jira/atlassian-jira/WEB-INF/lib/atlassian-extras-3.2.jar_bak
     
     docker cp atlassian-extras-3.1.2.jar 97:/opt/atlassian/jira/atlassian-jira/WEB-INF/lib/
     
     docker restart 97     #97為jira容器短id
     ```

   - ![img](https://images2018.cnblogs.com/blog/1093334/201807/1093334-20180730163753390-1265631516.png)

9. 破解成功

   - ![img](https://images2018.cnblogs.com/blog/1093334/201807/1093334-20180730164021127-916880101.png)