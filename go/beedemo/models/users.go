package models

import "time"

/*
mysql:
use demo;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '用户状态，0表示正常，1 表示禁用，2 表示删除',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/
type User struct {
	ID         int       `orm:"column(id);auto;pk",json:"id"`
	Name       string    `orm:"column(name);size(100)",json:"name"`
	Email      string    `orm:"column(email);size(100);null","email"`
	Status     int       `orm:"column(status)",json:"staus"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null",json:"createTime"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null",json:"updateTime"`
}

func NewUserModel() *User {
	return &User{}
}
func (m *User) TableName() string {
	return "users"
}
