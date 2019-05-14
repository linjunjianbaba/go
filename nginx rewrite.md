nginx rewrite

1，将www.myweb.com/connect 跳转到connect.myweb.com

rewrite ^/connect$ http://connect.myweb.com permanent;

rewrite ^/connect/(.*)$ http://connect.myweb.com/$1 permanent; 

2，将connect.myweb.com 301跳转到www.myweb.com/connect/ 

if ($host = "connect.myweb.com"){

rewrite ^/(.*)$ http://www.myweb.com/connect/$1 permanent;

​    } 

3，myweb.com 跳转到www.myweb.com

if ($host != 'www.myweb.com' ) { 

rewrite ^/(.*)$ http://www.myweb.com/$1 permanent; 

​    }

4，www.myweb.com/category/123.html 跳转为 category/?cd=123

rewrite "/category/(.*).html$" /category/?cd=$1 last;

5，www.myweb.com/admin/ 下跳转为www.myweb.com/admin/index.php?s=

if (!-e $request_filename){

rewrite ^/admin/(.*)$ /admin/index.php?s=/$1 last;

​    } 

6，在后面添加/index.php?s=

if (!-e $request_filename){

​    rewrite ^/(.*)$ /index.php?s=/$1 last;

​    } 

7，www.myweb.com/xinwen/123.html  等xinwen下面数字+html的链接跳转为404

rewrite ^/xinwen/([0-9]+)\.html$ /404.html last; 

8，http://www.myweb.com/news/radaier.html 301跳转 http://www.myweb.com/strategy/

rewrite ^/news/radaier.html http://www.myweb.com/strategy/ permanent;

9，重定向 链接为404页面

rewrite http://www.myweb.com/123/456.php /404.html last; 

10, 禁止htaccess

location ~//.ht {

​         deny all;

​     } 

11, 可以禁止/data/下多级目录下.log.txt等请求;

location ~ ^/data {

​     deny all;

​     }

12, 禁止单个文件

location ~ /www/log/123.log {

​      deny all;

​     }

 13, http://www.myweb.com/news/activies/2014-08-26/123.html 跳转为 http://www.myweb.com/news/activies/123.html

rewrite ^/news/activies/2014\-([0-9]+)\-([0-9]+)/(.*)$ http://www.myweb.com/news/activies/$3 permanent;

14，nginx多条件重定向rewrite

如果需要打开带有play的链接就跳转到play，不过/admin/play这个不能跳转

​        if ($request_filename ~ (.*)/play){ set $payvar '1';}
​        if ($request_filename ~ (.*)/admin){ set $payvar '0';}
​        if ($payvar ~ '1'){
​                rewrite ^/ http://play.myweb.com/ break;
​        }

15，http://www.myweb.com/?gid=6 跳转为http://www.myweb.com/123.html

 if ($request_uri ~ "/\?gid\=6"){return  http://www.myweb.com/123.html;}

正则表达式匹配，其中：

\* ~ 为区分大小写匹配

\* ~* 为不区分大小写匹配

\* !~和!~*分别为区分大小写不匹配及不区分大小写不匹配

文件及目录匹配，其中：

\* -f和!-f用来判断是否存在文件

\* -d和!-d用来判断是否存在目录

\* -e和!-e用来判断是否存在文件或目录

\* -x和!-x用来判断文件是否可执行

flag标记有：

\* last 相当于Apache里的[L]标记，表示完成rewrite

\* break 终止匹配, 不再匹配后面的规则

\* redirect 返回302临时重定向 地址栏会显示跳转后的地址

\* permanent 返回301永久重定向 地址栏会显示跳转后的地址

$args ：这个变量等于请求行中的参数，同$query_string
$content_length ： 请求头中的Content-length字段。
$content_type ： 请求头中的Content-Type字段。
$document_root ： 当前请求在root指令中指定的值。
$host ： 请求主机头字段，否则为服务器名称。
$http_user_agent ： 客户端agent信息
$http_cookie ： 客户端cookie信息
$limit_rate ： 这个变量可以限制连接速率。
$request_method ： 客户端请求的动作，通常为GET或POST。
$remote_addr ： 客户端的IP地址。
$remote_port ： 客户端的端口。
$remote_user ： 已经经过Auth Basic Module验证的用户名。
$request_filename ： 当前请求的文件路径，由root或alias指令与URI请求生成。
$scheme ： HTTP方法（如http，https）。
$server_protocol ： 请求使用的协议，通常是HTTP/1.0或HTTP/1.1。
$server_addr ： 服务器地址，在完成一次系统调用后可以确定这个值。
$server_name ： 服务器名称。
$server_port ： 请求到达服务器的端口号。
$request_uri ： 包含请求参数的原始URI，不包含主机名，如：”/foo/bar.php?arg=baz”。
$uri ： 不带请求参数的当前URI，$uri不包含主机名，如”/foo/bar.html”。
$document_uri ： 与$uri相同。

    符号解释：
    ^ 匹配字符串的开始
    / 匹配域名的分隔符
    . 匹配除换行符以外的任意字符
    * 重复零次或更多次
    (.*) 匹配任意字符
    .* 匹配任意文本
    $ 匹配字符串的结束
