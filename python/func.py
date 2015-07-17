#!/usr/bin/env python
# -*-coding: utf-8 -*-
' a module '
def power(x,n=2):
	s = 1
	while n > 0:
		n = n -1
		s = s * x
	return s
# print(power(5))
def add_end(L=[]):
	L.append('end')
	return L

def str2int(s):
   def fn(x, y):
       return x * 10 + y
   def char2num(s):
       return {'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8':8, '9': 9}[s]
   return reduce(fn, map(char2num, s))
# print(add_end([1,2]))
# print(add_end())
# print(add_end())
def calc(*numbers):
	sum = 0
	for n in numbers:
		sum = sum + n * n
	return sum
# print calc(1,2,4,5)
# print calc(*(1,2,3,4))
def person(name,age,**kw):
	print('name: ', name , "age: ", age, "other: ", kw)
extra = {'city': 'Beijing', 'job': 'Engineer'}
person('kang',12,**extra)
