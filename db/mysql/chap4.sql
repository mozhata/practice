-- source ~/go/src/practice/db/mysql/chap4.sql

use choose;

-- insert into teacher values('001', '张老师', '11000000000');
-- insert into teacher values('002', '李老师', '12000000000');
-- insert into teacher values('003', '王老师', '13000000000');

-- insert into classes(class_no, class_name, department_name) values(null, '2012自动化1班', '机电工程');
-- insert into classes(class_no, class_name, department_name) values(null, '2012自动化2班', '机电工程');
-- insert into classes(class_no, class_name, department_name) values(null, '2012自动化3班', '机电工程');

-- insert into course values(null, 'java语言程序设计', default, '暂无', '已审核', '001');
-- insert into course values(null, 'MySQL 数据库', 150, '暂无', '已审核', '002');
-- insert into course values(null, 'c语言程序设计', 230, '暂无', '已审核', '003');

-- insert into classes values(null, '2012计算机应用1班', '信息工程');

-- insert into student values
-- 	('2012001', '张三', '15000000000', 1),
-- 	('2012002', '李四', '16000000000', 1),
-- 	('2012003', '王五', '17000000000', 3),
-- 	('2012004', '马六', '18000000000', 2),
-- 	('2012005', '田七', '19000000000', 2);

-- create table new_student like student;
-- insert into new_student select * from student;
-- select * from new_student;

-- replace into student values ('2012001', '张三丰', '15000000000', 1);
-- replace into student values ('2012001', '张三', '15000000000', 1);

-- update classes set department_name = '机电工程学院' where class_no<=3;

-- delete from classes where class_name = '2012计算机应用1班';
-- select * from classes;


-- create table new_class like classes;
-- insert into new_class select * from classes;
-- select * from new_class;

-- show create table new_class;
-- delete from new_class;
-- show create table new_class;

-- truncate table new_class;
-- show create table new_class;


-- insert into new_student values('2012006', 'Martin', 'mar\tin@gmail.com', 3);
-- insert into new_student values('2012007', 'O\_Neil', 'o_\neil@gmail.com', 3);
-- select * from new_student;

-- select * from new_student where student_name like '%\_%';

-- insert into choose values
-- 	(null, '2012001', 2, 40, now()),
-- 	(null, '2012001', 1, 50, now()),
-- 	(null, '2012002', 3, 60, now()),
-- 	(null, '2012002', 2, 70, now()),
-- 	(null, '2012003', 1, 80, now()),
-- 	(null, '2012004', 2, 90, now()),
-- 	(null, '2012005', 3, null, now()),
-- 	(null, '2012006', 1, null, now())
-- select * from choose;