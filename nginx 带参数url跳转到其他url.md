nginx 带参数url跳转到其他url

```shell
location / {
        root /home/webapp/public;
        index index.html;
        if ($query_string ~* ^(.*)rid=17332(.*)$) {
             rewrite ^(.*) http://h5.eqxiu.com/ls/7WTbGi7c;
        }
}
```

