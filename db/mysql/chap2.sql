-- source ~/go/src/practice/db/mysql/chap2.sql

-- create table my_table(
-- 	today datetime,
-- 	name char(20)
-- );

-- insert into my_table values (now(), 'a');
-- insert into my_table values (now(), 'a');
-- insert into my_table values (now(), NULL);
-- insert into my_table values (now(), '');

-- use test_mysql;
-- create table second_table(
-- 	today datetime,
-- 	name char(20)
-- );

-- use test_mysql;
-- create table today(
-- 	t1 datetime,
-- 	t2 timestamp
-- );

-- insert into today values (now(), now());
-- insert into today values (NULL, NULL);

-- use test_mysql;
-- create table person(
-- 	sex enum('男', '女'),
-- 	interest set('听音乐', '看电影', '旅游', '购物', '游泳', '游戏')
-- );

-- insert into person values ('男', '看电影,游泳,听音乐');

-- use  test_mysql;
-- create table nowadays (
-- 	t1 datetime,
-- 	t2 timestamp,
-- 	primary key(t1, t2)
-- );

-- use test_mysql;
-- create table test(
-- 	test_no char(10),
-- 	test_name char(10),
-- 	constraint test_pk primary key (test_no),
-- 	constraint name_unique unique (test_name)
-- );

