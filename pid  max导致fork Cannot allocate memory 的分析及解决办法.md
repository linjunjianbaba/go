pid  max导致fork: Cannot allocate memory 的分析及解决办法

 今天遇到[服务器](https://www.baidu.com/s?wd=%E6%9C%8D%E5%8A%A1%E5%99%A8&tn=24004469_oem_dg&rsv_dl=gh_pl_sl_csd)无法SSH，VNC操作命令提示fork:cannot allocate memory

 free查看内存还有（注意，命令可能要多敲几次才会出来）

查看最大进程数 sysctl kernel.pid_max

ps -eLf | wc -l查看进程数

确认是进程数满了

修改最大进程数后系统恢复

echo 1000000 > /proc/sys/kernel/pid_max

永久生效
echo "kernel.pid_max=1000000 " >> /etc/sysctl.conf
sysctl -p