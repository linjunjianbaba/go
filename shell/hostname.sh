#!/bin/bash

for name in `cat /root/hostname`; do
    ssh root@$name "hostnamectl set-hostname $name"
done;
