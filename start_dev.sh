#!/bin/bash
cd ./docker-mysql

echo "Stop docker-compose...."
docker-compose down

echo "Build Docker Images"
docker-compose build

echo "Starting docker..."
docker-compose up -d

echo "Started docker."

cd ..
echo "Waiting MYSQL master Server"
./wait_for_it.sh 3306

echo "Waiting MYSQL slave Server"
./wait_for_it.sh 13306


#echo "Copy MYSQL binlog"
#docker cp dockermysql_mysql_1:/var/log/mysql ./data

# create Replication
mysql -u root -h 0.0.0.0 -P 3306 -e "GRANT REPLICATION SLAVE ON *.* TO 'root'@'172.17.0.%' IDENTIFIED BY 'root';"
mysql -u root -h 0.0.0.0 -P 3306 -e "flush privileges;"

# slave setting
master_port=3306
file=$(mysql -u root -h 0.0.0.0 -P $master_port -e "show master status\G" | grep File | awk '{print $2}')
position=$(mysql -u root -h 0.0.0.0 -P $master_port -e "show master status\G" | grep Position | awk '{print $2}')
master_host=`docker inspect --format '{{ .NetworkSettings.Networks.dockermysql_default.IPAddress }}' dockermysql_master_1`

echo $file $position $master_host

mysql -u root -h 0.0.0.0 -P 13306 -e "CHANGE MASTER TO MASTER_HOST='$master_host', MASTER_USER='root', MASTER_LOG_FILE='$file', MASTER_LOG_POS=$position;"
mysql -u root -h 0.0.0.0 -P 13306 -e "start slave"

# create test data
./test_db.py

sleep 5
echo "Copy MYSQL slave binlog"

# change slave binlog-format
mysql -u root -h 0.0.0.0 -P 13306 -e "FLUSH TABLES WITH READ LOCK;
FLUSH LOGS;
SET GLOBAL binlog_format = 'ROW';
FLUSH LOGS;
UNLOCK TABLES;"
./test_db.py
docker cp dockermysql_slave_1:/var/log/mysql/ ./data/binlog/
docker cp dockermysql_slave_1:/var/lib/mysql/ ./data/relay/