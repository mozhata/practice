#!/usr/bin/env python
# -*-coding: utf-8 -*-
' a module '
from collections import Iterable
import os
import types
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
# person('kang',12,**extra)

def f1(a,b=2,*args,**kw):
	print('a =', a, 'b =', b, 'args =', args, 'kw= ', kw)
# f1('a',2,12,'qw',city='sdgl')
def fact(n):
	if n==1:
		return 1
	return n * fact(n-1)
# print(fact(90))
L = list(range(10))
# print(L[:10:2])
d = {'a': '1', 'b': '2', 'c': '3'}
# for key in d:
# # 	print(key,d[key])
# for i ,v in enumerate(L):
# 	print(i,v)
# print(isinstance('abs',Iterable))
# print([x * x for x in range(1,11) if x % 2 == 0])
# print([m + n for m in 'abs' for n in 'ABS'])
# print([d for d in os.listdir('.')])
# for k, v in d.items():
# 	print(k, '=', v)

# L = ['Hello', 'World', 'IBM', 'Apple']
# print([s.lower() for s in L])
# print(isinstance(1,(str,int)))
def fib(max):
	n, a, b = 0, 0, 1
	while n < max:
		yield b
		a, b = b, a + b
		n = n + 1
	# return 'done'
f = fib(6)
# print(type(f))
# print(type(fib))
# for i in f:
# 	print(i)
# print(isinstance((x for x in range(1,9)),Iterable))
class Student(object):
	def __init__(self,name,score):
		self.name = name
		self.score = score

	def print_score(self):
		print('%s: %s' % (self.name, self.score))

bart = Student('Bart Simpson',55)
# print(Student)
# print(bart)
# bart.print_score()

class Student1(object):
	pass

bart = Student1()
bart.name = 'yang'
# print(bart.name)
class Student2(object):
	def __init__(self, name, score):
		self.__name = name
		self.__score = score

	def print_score(self):
		print('%s: %s' % (self.__name, self.__score))

# st = Student2('yang',99)
# print(st._Student2__name)
# st.print_score()
class Animal(object):
	def run(self):
		print('Animal is running')

class Dog(Animal):
	pass

dog = Dog()
# dog.run()
# print(isinstance(dog,Animal))
# print(type(123)==int)
# print(type('asgd')==str)
def fn():
	pass

# print(type(fn)==types.FunctionType)
# print(type(abs)==types.BuiltinFunctionType)
# print(type(lambda x: x)==types.LambdaType)
# print(type((x for x in range(10)))==types.GeneratorType)
# print(isinstance(dog,Animal))
class MyObject(object):
    def __init__(self):
        self.x = 9
    def power(self):
        return self.x * self.x

obj = MyObject()
# print(hasattr(obj,'x'))
# print(obj.x)
# print(hasattr(obj,'y'))
setattr(obj,'y',19)
# print(hasattr(obj,'y'))
# print(getattr(obj,'y'),getattr(obj,'x'))
# print(hasattr(obj,'power'))
class Studentt(object):
	k = 'king'

	def __init__(self,name):
		self.name = name

s = Studentt('bo')
# print(s.name)
# s.score = 90
# print(s.score)
print(s.k,Studentt.k)