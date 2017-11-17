# -*- coding: utf-8 -*-
import sys
import time
from selenium import webdriver
from selenium.webdriver.common.by import By
from  selenium.webdriver.remote.webelement import WebElement
import urllib2
import json
# import imp
# func = imp.load_source("function","/home/zyk/go/src/practice/python/func.py")
# print func.power(4)
# import sys
# sys.path.append("/home/zyk/go/src/practice/python")
# import func

# print(func.power(6))

DOMAIN = "crm-test.meiqia.com"

def get_ip(domain):
    import socket
    return socket.getaddrinfo(domain, "http")[0][4][0]


def get_user_version(token, uid):
    ip = get_ip(DOMAIN)
    query_api = "http://" + ip + ":7010/api/v1.0/meiqia/User/"+uid
    print("query_api: ", query_api)
    req = urllib2.Request(query_api)
    req.add_header("x-token", token)
    try:
        resp = urllib2.urlopen(req)
        if resp.getcode() != 200:
            print("get version failed, status code not 200: ", resp.getcode())
            return False, 0
        # get version
        body = json.load(resp)
        code = body.get("code", -1)
        if code != 0:
        	return False, 0
       	body = body.get("body", {})
       	version = body.get("verson", -1)
       	if version == -1:
       		return False, 0
       	return True, version
    except urllib2.HTTPError, e:
        print("get version failed, status code: ", e.code)
        return False, 0

get_user_version("AZoGABZbClp4AEFRQUNrNUh4RklheERBQUEwbEMxSFNPRjloUzNCUUFBQVFBQ2s1SHhGSVoxREFBQUE0OEV5RFM3OUJTc0F3QUHUvV5v4fAg21vp854TZS8c1IeibuKH8VG7nFGexlear_BiODGVIJi53zkQy5i4VwOZFx7kPXnQh-vujRy37tNU", "AQACk5HxFIaxDAAAZ8etigHS9hTSCAAA")

