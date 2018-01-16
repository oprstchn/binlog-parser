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
echo "Waiting MYSQL Server"
./wait_for_it.sh ./test_db.py

echo "Copy MYSQL binlog"
docker cp dockermysql_mysql_1:/var/log/mysql ./data