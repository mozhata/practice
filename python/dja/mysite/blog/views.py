# from django.shortcuts import render

# # Create your views here.
from django.http import HttpResponse
from django.template import loader,Context, Template

# def index(req):
# 	t = loader.get_template('index.html')
# 	c = Context({})
# 	return HttpResponse(t.render(c))

from django.shortcuts import render_to_response
def index(req):
	usr = {'name':'yongkang','age':23,'sex':'male'}
	# usr = Person('yang',22,'male')
	booklist = ['pytho','java','Golang']
	return render_to_response('index.html',{'title':'test title','user': usr,'booklist':booklist})

class Person(object):
	def __init__(self,name,age,sex):
		self.name = name
		self.age = age
		self.sex = sex
	def say(self):
		return 'I am ' + self.name

def page(req):
	t = loader.get_template("page.html")
	c = Context({'title':'mypage','uname':'zykzhang'})
	return HttpResponse(t.render(c))

def home(req):
	t = Template("<h1>hello {{name}}</h1>")
	c = Context({'name': 'yangyang'})
	return HttpResponse(t.render(c))

def index2(req):
	return render_to_response('page.html',{'uname': 'zhenzhen','title': "TT"})