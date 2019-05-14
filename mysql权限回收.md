mysql权限回收：

回收全部：revoke all privileges,grant option from 'sibu_draw'@'%';

重新赋权：grant select on `sibu_wesale\_%`.* to 'sibu_wesale'@'%' identified by '5IcMJTpvRQZyr5H11pEy';

刷新权限：flush privileges；