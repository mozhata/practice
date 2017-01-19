from urllib import request

link = "http://www.baidu.com"
resp = request.urlopen(link)
print (resp.read().decode("utf8"))