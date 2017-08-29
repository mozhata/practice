package models

import "time"

/*
mysql:
use demo;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `nickname` varchar(100) DEFAULT NULL COMMENT '昵称',
  `email` varchar(100) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL COMMENT '手机号',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '用户状态，0表示正常，1 表示禁用，2 表示删除',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
*/
type User struct {
	ID         uint64    `orm:"column(id);auto;pk" json:"id"`
	Name       string    `orm:"column(name);size(100)" json:"name"`
	Nickname   string    `orm:"column(nickname);size(100)" json:"nickname"`
	Email      string    `orm:"column(email);size(100);null" json:"email"`
	Phone      string    `orm:"column(phone);size(100);null" json:"phone"`
	Status     int       `orm:"column(status)" json:"staus"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" json:"createTime"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" json:"updateTime"`
}

func NewUserModel() *User {
	return &User{}
}

func (m *User) TableName() string {
	return UserTable
}

func (u *User) IsValid() bool {
	return u.Name != "" || u.Email != ""
}
