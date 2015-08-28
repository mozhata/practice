# encoding=utf-8

dic1 = {'a': 1, 'c': 3, 'e': 5}
dic2 = {'b': 2, 'd': 4, 'f': 6}

# ## method-1 效率较低
# mergedDic1 = dict(dic1.items() + dic2.items())
# print(mergedDic1)

# ## method-2效率较高
# mergedDic2 = dict(dic1, **dic2)
# print(mergedDic2)

# # method-3,等效于method-2
# mergedDic3 = dic1.copy()
# mergedDic3.update(dic2)
# print(mergedDic3)


# copy = dic1.copy()		# 值传递
# index = dic1   # 地址引用,对copy的操作会作用在dic1上

# del(index['a'])
# print(index)

# print(index.pop('e'))
# index.clear()
# print(index)

# print(dic1)		# index 和dic1 是引用同一个地址,index的
# print(copy)


# ## value的值会按照对应的数据类型输出
dict1 = {"a": ("apple", ), "bo": {"b": "banana", "o": "orange"}, "g": ["grape", "grapefruit"]}
# print(dict1['a'])
# print(dict1['bo'])


## 都是list
# print(dict1.keys())
# print(dict1.values())


# ## get 方式,key存在则取到value,不存在则取到None,或者自己设一个值
# print(dict1.get("a"))
# print(dict1.get("b"))
# print(dict1.get("b", "ops.."))


# ## 更新,(或者合并dic2到dic1)
# dict1 = {"a": "apple", "b": "banana"}
# print(dict1)
# dict2 = {"c": "grape", "d":  "orange", 'b': "ops.."}
# dict1.update(dict2)
# print(dict1)


# ## 设置默认值
# dic = {}
# dic.setdefault("a")		# {'a': None}
# print(dic)
# dic = {}
# dic.setdefault("a", "default")		# {'a': 'default'}
# print(dic)