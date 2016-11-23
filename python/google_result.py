#conding=utf-8

import simplejson as json1
import requests
from bs4 import BeautifulSoup
import urllib
import csv

googleAPI = "https://www.google.co.jp/search"
source = "site.txt"

sourceFile = open(source)
destFile = open("test.csv", "wb")
writer = csv.writer(destFile,dialect="excel")
# while line:
#     print line
count = 1
while True:
	if count > 10:
		break
	line = sourceFile.readline()
	if not line:
		print("finished")
		break

	dic = {"q": "site:" + line[:-1]}
	link = googleAPI + "?" + urllib.urlencode(dic) + "&oq=site%3A&aqs=chrome.0.69i59j69i57j69i58.3524j0j7&sourceid=chrome&ie=UTF-8"
	print("link: ", link)

	doc = requests.get(link)
	soup = BeautifulSoup(doc.text)
	result = soup.select("#resultStats").txt
	print("result: ", result)
	record = [link, result]
	writer.writerow(record)

	count = count + 1

sourceFile.close()
destFile.close()