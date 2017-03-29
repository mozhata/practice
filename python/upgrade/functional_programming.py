# -*- coding: utf-8 -*-
import math
import time
import functools


# 列表生成式

# print range(10)
# print range(1, 10)
# print range(1, 0, 2)
# L = range(10)
# print L[0:]
# print L[1:3]
# print L[:]
# print L[::]
# print L[::2]
# print L[::-1]
# print [x*x for x in range(10)]
# print [x*x for x in range(10) if x % 2 == 0]
# print [m + n for m in "abc" for n in "ABC"]
# d = {"java": "99", "C": "98", "c++": "97"}
# print [k + " = " + v for k, v in d.iteritems()]
# # L = ['Java' , 'C' , 'Swift' , 'Python' , 123] ， 现在有 list 中包含字符串，和整数，把list中得大写字符转为小写，推到出另外一个list":
# L = ["Java", "C", "Swift", "php", 123, "GO"]
# print [s.lower() if isinstance(s, str) else s for s in L]


# 把函数作为参数

def add(a, b, f):
    return f(a) + f(b)

# print add(2, -9, abs)


# map()函数 内置的高阶函数，它接收一个函数 f 和一个 list，并通过把函数 f 依次作用在 list 的每个元素上，得到一个新的 list 并返回。
# map()函数不改变原有的 list，而是返回一个新的 list。

def a2(x):
    return x * x

# print map(a2, [1,2,3,4])


# reduce()函数接受一个函数 f，一个list，reduce()传入的函数 f 必须接收两个参数，reduce()对list的每个元素反复调用函数f，并返回最终结果值。
# reduce()还可以接收第3个可选参数，作为计算的初始值

def add2(a, b):
    return a + b

# print reduce(add2, [1, 2, 3, 4])
# print reduce(add2, [1, 2, 3, 4], 100)


# filter()函数接收一个函数 f 和一个list，这个函数 f 的作用是对每个元素进行判断，返回 True或 False，filter()根据判断结果自动过滤掉不符合条件的元素，返回由符合条件元素组成的新list。

def is_odd(x):
    return x % 2 == 1


def is_sqrt(x):
    return x and math.sqrt(x)**2 == x

# print filter(is_odd, [1, 4, 6, 7, 9, 12, 17])
# print filter(is_odd, range(10))
# L = [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
# print filter(is_sqrt, L)


# sorted()也是一个高阶函数，它可以接收一个比较函数来实现自定义排序，比较函数的定义是，传入两个待比较的元素 x, y，如果 x 应该排在 y 的前面，返回 -1，如果 x 应该排在 y 的后面，返回 1。如果 x 和 y 相等，返回 0。

def reverse_cmp(x, y):
    if x > y:
        return -1
    else:
        return 1


def cmp_ignore_case(a, b):
    if a.upper() > b.upper():
        return 1
    else:
        return -1

# print sorted(range(10), reverse_cmp)
# print sorted(['bob', 'about', 'Zoo', 'Credit'])
# print sorted(['bob', 'about', 'Zoo', 'Credit'], cmp_ignore_case)


# 闭包
def multi(j):
    return lambda: j * j


def count():
    fs = []
    for i in range(1, 4):
        fs.append(multi(i))
    return fs

# f1, f2, f3 = count()
# print f1(), f2(), f3()
# def count():
#     fs = []
#     for i in range(1, 4):
#         def f(i):
#             return lambda : i*i
#         fs.append(f(i))
#     return fs
# f1, f2, f3 = count()
# print f1(), f2(), f3()


# 匿名函数
# 关键字lambda 表示匿名函数，冒号前面的 x 表示函数参数。
# 匿名函数有个限制，就是只能有一个表达式，不写return，返回值就是该表达式的结果

# print map(lambda x: x * 2, range(10))
# print sorted([1, 3, 9, 5, 0], lambda x, y: -cmp(x, y))
# print filter(lambda s: s and len(s.strip()) > 0, ['test', None, '', 'str', '  ', 'END'])


# 编写无参数decorator

def log(f):
    def fn(*agrs, **kw):
        print "call " + f.__name__ + "()... "
        return f(*agrs, **kw)
    return fn


def performance(f):
    @log
    def fn(*args, **kw):
        print "%s called" % time.strftime('%Y-%m-%d %H: %M: %S', time.localtime(time.time()))
        return f(*args, **kw)
    return fn


def time_unit(s):
    @performance
    def deco(f):
        def wrapper(*args, **kw):
            print "time unit is %s" % s
            f(*args, **kw)
        return wrapper
    return deco


@log
def fatorial(n):
    return reduce(lambda x, y: x*y, range(1, n+1))


@performance
def add3(x, y):
    return x + y


@time_unit("ms")
def test():
    pass

# print fatorial(7)
# print add3(1, 4)
# test()


# 完善decorator

def log2(f):
    @functools.wraps(f)
    def wrapper(*args, **kw):
        print "call " + f.__name__ + "()..."
        return f(*args, **kw)
    return wrapper


def f1():
    pass


@log
def f2():
    pass


@log2
def test2():
    pass

# print f1.__name__
# print f2.__name__
# print test2.__name__


# 偏函数

# print int("123")
# print int("123", base=8)
# int2 = functools.partial(int, base=2)
# print int2("10000")
