volumes_linux:
	sudo mkdir -p /mnt/sda1/practice/data/mysql

clean_mysql:
	sudo rm -rf /mnt/sda1/practice/data/mysql*

mysql_server:
	docker run --name mozhata_mysql -v /mnt/sda1/practice/data/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=mysql_access_key -p 3333:3306 mozhata/mysql:utf8 -d

connect_mysql:
	mysql -h 127.0.0.1 -P 3333 -u root -pmysql_access_key

bash:
	docker exec -it mozhata_mysql bash

containers:
	docker-compose up -d

