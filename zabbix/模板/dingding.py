#!/usr/bin/python
# -*- coding: utf-8 -*-
# Author: xxxxxxxx
import requests
import json
import sys
import os

headers = {'Content-Type': 'application/json;charset=utf-8'}
api_url = "https://oapi.dingtalk.com/robot/send?access_token=88da9ed421bf34a2e0c0cbed6a0f60cd2660d1efb547028f13257b4a5519bb25"

def msg(text):
    json_text= {
     "msgtype": "text",
        "at": {
            "atMobiles": [
                "zabbix"
            ],
            "isAtAll": False
        },
        "text": {
            "content": text
        }
    }
    print requests.post(api_url,json.dumps(json_text),headers=headers).content

if __name__ == '__main__':
    text = sys.argv[1]
    msg(text)
