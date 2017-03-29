# coding=utf-8

for i in range(4):
	print("this is the first level, i: %s" % i)
	for j in range(2):
		if i > 1:
			print("i > 1, break, i: %s" % i)
		print("i: %s, j: %s" % (i, j))