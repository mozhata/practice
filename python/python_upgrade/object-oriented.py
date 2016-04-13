# -*- coding: utf-8 -*-


# class Person(object):
#     def __init__(self, name, gender, birth, **kw):
#         self.name = name
#         self.gender = gender
#         self.birth = birth
#         self.__dict__.update(kw)

# xiaoming = Person("Xiao Ming", "Male", "1991-01-09")
# xiaohong = Person("Xiao Hong", "Female", "1992-01-09", job="coder")
# print xiaohong.name
# print xiaoming.name, xiaohong.birth, xiaohong.job


# # 修改类属性会导致所有实例访问到的类属性全部都受影响
# # 当实例属性和类属性重名时，实例属性优先级高
# class Person(object):
#     def __init__(self, name, gender, birth, **kw):
#         self.name = name
#         self.gender = gender
#         self.birth = birth
#         self.__dict__.update(kw)
#         self._title = "Mr"
#         self.__job = "student"
#         Person.count = Person.count + 1
#     address = "Earth"
#     count = 0

# # print Person.count
# xiaoming = Person("Xiao Ming", "Male", "1991-01-09")
# # print Person.count, xiaoming.count
# xiaohong = Person("Xiao Hong", "Female", "1992-01-09", job="coder")
# # print Person.count, xiaoming.count, xiaohong.count
# # print xiaohong.name
# # print xiaoming.name, xiaohong.birth, xiaohong.job
# # print xiaohong._title
# # print xiaohong.address, Person.address, xiaoming.address
# # Person.address = "beijing"
# # print xiaohong.address, Person.address, xiaoming.address
# # xiaohong.address = "shanghai"
# # print xiaohong.address, Person.address, xiaoming.address
# # del xiaohong.address
# # print xiaohong.address, Person.address, xiaoming.address


# # 定义实例方法
# class Person(object):

#     def __init__(self, name, score):
#         self.name = name
#         self.__score = score
#         self.grade = lambda: 'A'

#     def get_grade(self):
#         if self.__score > 90:
#             return "A"
#         elif self.__score > 60:
#             return "B"
#         else:
#             return "C"

# p1 = Person("Bob", 88)
# print p1.get_grade()
# print p1.grade
# print p1.grade()


# 定义类方法

class Person(object):

    count = 0
    __counting = 0

    def __init__(self, name):
        self.name = name
        Person.count = Person.count + 1
        Person.__counting = Person.__counting + 1

    @classmethod
    def how_many(cls):
        return cls.count

    @classmethod
    def get_counting(cls):
        return cls.__counting

print Person.how_many(), Person.get_counting()
bob = Person("Bob")
print Person.how_many(), Person.get_counting()
print bob.how_many(), bob.get_counting()
