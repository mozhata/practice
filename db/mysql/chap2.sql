-- source ~/go/src/practice/db/mysql/chap2.sql

-- create table my_table(
-- 	today datetime,
-- 	name char(20)
-- );

-- insert into my_table values (now(), 'a');
-- insert into my_table values (now(), 'a');
-- insert into my_table values (now(), NULL);
-- insert into my_table values (now(), '');

-- use choose;
-- create table second_table(
-- 	today datetime,
-- 	name char(20)
-- );

-- use choose;
-- create table today(
-- 	t1 datetime,
-- 	t2 timestamp
-- );

-- insert into today values (now(), now());
-- insert into today values (NULL, NULL);

-- use choose;
-- create table person(
-- 	sex enum('男', '女'),
-- 	interest set('听音乐', '看电影', '旅游', '购物', '游泳', '游戏')
-- );

-- insert into person values ('男', '看电影,游泳,听音乐');

-- use  choose;
-- create table nowadays (
-- 	t1 datetime,
-- 	t2 timestamp,
-- 	primary key(t1, t2)
-- );

-- use choose;
-- create table test(
-- 	test_no char(10),
-- 	test_name char(10),
-- 	constraint test_pk primary key (test_no),
-- 	constraint name_unique unique (test_name)
-- );

-- use choose;
-- create table today1 like today;
-- show create table today1;
-- select * from today1;

-- use choose;
-- create table today2 select * from today;
-- show create table today2;
-- select * from today2;

-- alter table person drop interest;
-- alter table person add person_no int auto_increment primary key first;
-- alter table person add person_name char(10) not null after person_no;
-- alter table person change person_name name char(20);
-- alter table person modify name char(30);

-- select constraint_name, constraint_type from information_schema.table_constraints
-- where table_schema='choose' and table_name='person';

-- delete from person;
-- alter table person add constraint name_unique unique (name);

-- alter table person drop index name_unique;

-- alter table person rename human;


-- DROP TABLE IF EXISTS `book`;
-- CREATE TABLE `book`(
-- 	`isbn` char(20) PRIMARY KEY,
-- 	`name` char(100) NOT NULL,
-- 	`brief_introduction` text NOT NULL,
-- 	`price` decimal(6,2),
-- 	`publish_time` date NOT NULL
-- 	-- unique index isbn_unique (`isbn`),
-- 	-- index name_index (`name`, (20)),
-- 	-- fulltext index brief_fulltext (`name`, `brief_introduction`),
-- 	-- index complex_index (`price`, `publish_time`)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- | book  | CREATE TABLE `book` (
--   `isbn` char(20) NOT NULL,
--   `name` char(100) NOT NULL,
--   `brief_introduction` text NOT NULL,
--   `price` decimal(6,2) DEFAULT NULL,
--   `publish_time` date NOT NULL,
--   PRIMARY KEY (`isbn`),
--   UNIQUE KEY `isbn_unique` (`isbn`),
--   FULLTEXT KEY `brief_introduction` (`name`,`brief_introduction`)
-- )

-- create unique index isbn_unique on book (isbn);
-- alter table book add fulltext index brief_introduction (name, brief_introduction);
-- create index name_index on book (name (20));
-- alter table book add index complex_index (price,publish_time);

-- alter table course add fulltext index description_fulltext (description);

-- drop index complex_index on book;

