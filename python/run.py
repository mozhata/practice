# import imp
# func = imp.load_source("function","/home/zyk/go/src/practice/python/func.py")
# print func.power(4)
import sys
sys.path.append("/home/zyk/go/src/practice/python")
import func
print(func.power(6))
