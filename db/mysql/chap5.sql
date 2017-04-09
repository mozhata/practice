-- source ~/go/src/practice/db/mysql/chap5.sql

use choose;

-- select version(), now(), pi(), null = null, null != null, null is null;
-- select version() 版本号, now() as 服务器当前时间, pi() PI的值, 1 + 2 求和;

-- select * from student;
-- select distinct department_name from classes;

-- select * from student limit 0,2;

-- insert into classes values(4, '2012 自动化4班', '机电工程学院');
-- insert into student values('2012006', '张三丰', '20000000000', null);

-- select student_no, student_name, student_contact, student.class_no, class_name, department_name from student join classes on student.class_no=classes.class_no;

-- select student_no, student_name, student_contact, student.class_no, class_name, department_name from student left join classes on student.class_no=classes.class_no;

-- select student_no, student_name, student_contact, student.class_no, class_name, department_name from student right join classes on student.class_no=classes.class_no;

-- select student_no, student_name, student_contact, student.class_no, class_name, department_name from student right join classes on student.class_no=classes.class_no;

-- select student.student_no, student_name, course.course_no, course_name, score
-- from student inner join choose on student.student_no=choose.student_no
-- inner join course on choose.course_no=course.course_no;

-- select student.student_no, student_name, choose.course_no, course_name, score
-- from classes left join student on classes.class_no=student.class_no
-- join choose on student.student_no=choose.student_no
-- join course on course.course_no=choose.course_no
-- where class_name="2012自动化2班";

-- select student_no, student_name, student_contact, student.class_no, class_name, department_name
-- from student, classes
-- where student.class_no=classes.class_no;


	/*select trigger_rule, count(trigger_rule) from tenx_alert_notify_history where strategy_id='SID-3ns8a3a3f' and trigger_rule in ('cpu > 30', 'cpu > 33') group by trigger_rule;*/