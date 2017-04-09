# -*- coding: utf-8 -*-


class Person(object):
    def __init__(self, name, gender):
        self.name = name
        self.gender = gender

    def __str__(self):
        return "(Person: %s, %s)" % (self.name, self.gender)

    __repr__ = __str__


# 函数super(Student, self)将返回当前类继承的父类，即 Person ，然后调用__init__()方法，注意self参数已在super()中传入，在__init__()中将隐式传递，不需要写出（也不能写）。
class Student(Person):
    def __init__(self, name, gender, score):
        super(Student, self).__init__(name, gender)
        self.score = score

    def who_am_i(self):
        return "I am a Student, my name is %s" % self.name


class Teacher(Person):
    def __init__(self, name, gender, course):
        super(Teacher, self).__init__(name, gender)
        self.course = course


# t = Teacher('Alice', 'Female', 'English')
# print t.name
# print t.course

# p = Person('Tim', 'Male')
s = Student('Bob', 'Male', 88)
# t = Teacher('Alice', 'Female', 'English')
# print isinstance(p, Person)
# print isinstance(p, Student)
# print isinstance(p, Teacher)
# print isinstance(s, Person)
# print isinstance(s, Student)
# print isinstance(s, Teacher)
# print isinstance(s, object)
# print s
# s


# 多重继承

class A(object):
    def __init__(self, a):
        print "init A ..."
        self.a = a


class B(A):
    def __init__(self, a):
        super(B, self).__init__(a)
        print "init B ..."


class C(A):
    def __init__(self, a):
        super(C, self).__init__(a)
        print "init C ..."


class D(B, C):
    def __init__(self, a):
        super(D, self).__init__(a)
        print "init D ..."


# d = D("d")


# 获取对象信息
# 如果已知一个属性名称，要获取或者设置对象的属性，就需要用 getattr() 和 setattr( )函数了：

# s = Student("Bob", "Male", 99)
# print getattr(s, "name")
# setattr(s, "name", "Adam")
# print s.name
# print getattr(s, "skill", "Ops..")  # 获取age属性，如果属性不存在，就返回默认值20


class Fib(object):
    def __init__(self, num):
        self.num = num
        self.fibo = [0, 1]
        i = 2
        while i < num:
            self.fibo.append(self.fibo[i-2] + self.fibo[i-1])
            i = i + 1

    def __str__(self):
        return str(self.fibo)

    def __len__(self):
        return len(self.fibo)

# print Fib(10)
# print len(Fib(12))


# __slots__
# 由于Python是动态语言，任何实例在运行期都可以动态地添加属性。
# 如果要限制添加的属性，例如，Student类只允许添加 name、gender和score 这3个属性，就可以利用Python的一个特殊的__slots__来实现。

class Person22(object):
    __slots__ = ("name", "gender")

    def __init__(self, name, gender):
        self.name = name
        self.gender = gender


# __call__
# 在Python中，函数其实是一个对象：
# 一个类实例也可以变成一个可调用对象，只需要实现一个特殊方法__call__()。

class Fib2(object):

    def __call__(self, num):
        L = [0, 1]
        for i in range(2, num):
            L.append(sum(L[-2:]))
        return L


f = Fib2()
print f(10)
