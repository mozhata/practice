# -*- coding: utf-8 -*-

import simplejson as json1
import requests

url = "http://data.api.gkcx.eol.cn/soudaxue/queryProvince.html?messtype=json&url_sign=queryprovince&page=1&size=100"
json = requests.get(url)
dic = json1.loads(json.text)
print type(dic)
print dic["totalRecord"]
