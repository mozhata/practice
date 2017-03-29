from django.conf.urls import url

import views

urlpatterns = [
    url(r'^$', views.index, name='index'),
    url(r'^page/$', views.page, name='page'),
    url(r'^home/$', views.home, name='home'),
    url(r'^2b/$', views.index2, name='index2'),
]