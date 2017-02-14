from urllib import request

link = "http://www.baidu.com"
# resp = request.urlopen(link)
req = request.Request(link)
req.add_header(
    "User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")

resp = request.urlopen(req)
print(resp.read().decode("utf8"))


def abc():
    pass

abc()
