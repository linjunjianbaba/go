腾讯云自建VPN上外网iptables设置：

1.iptables -t nat -A POSTROUTING -s 10.10.10.0/24 -o eth0 -j MASQUERADE

2.iptables -A INPUT -p tcp --dport 1723 -j ACCEPT

3.iptables -I FORWARD -p tcp --syn -i ppp+ -j TCPMSS --set-mss 1300   （关键）