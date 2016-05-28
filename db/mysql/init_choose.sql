-- source ~/go/src/practice/db/mysql/init_choose.sql

use choose;

DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`(
	`teacher_no` char(15) PRIMARY KEY,
	`teacher_name` char(15) NOT NULL,
	`teacher_contact` char(30) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `classes`;
CREATE TABLE `classes`(
	`class_no` int auto_increment PRIMARY KEY,
	`class_name` char(30) NOT NULL unique,
	`department_name` char(30) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`(
	`course_no` int auto_increment PRIMARY KEY,
	`course_name` char(15) NOT NULL,
	`up_limit` int default 60,
	`description`text NOT NULL,
	`status` char(10) default '未审核',
	`teacher_no` char(15) NOT NULL unique,
	CONSTRAINT course_teacher_fk FOREIGN KEY(teacher_no) REFERENCES teacher(teacher_no)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`(
	`student_no` char(11) PRIMARY KEY,
	`student_name` char(15) NOT NULL,
	`student_contact` char(30) NOT NULL,
	`class_no` int,
	CONSTRAINT student_class_fk FOREIGN KEY(class_no) REFERENCES classes(class_no)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `choose`;
CREATE TABLE `choose`(
	`choose_no` int auto_increment PRIMARY KEY,
	`student_no` char(11) NOT NULL,
	`course_no` int NOT NULL,
	`score` tinyint unsigned,
	`choose_time` datetime NOT NULL,
	CONSTRAINT choose_student_fk FOREIGN KEY(student_no) REFERENCES student(student_no),
	CONSTRAINT choose_course_fk FOREIGN KEY(course_no) REFERENCES course(course_no)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
