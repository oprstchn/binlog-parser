#!/bin/bash
cd ./docker-mysql

echo "Stop docker-compose...."
docker-compose down

echo "Build Docker Images"
docker-compose build

echo "Starting docker..."
docker-compose up -d

echo "Started docker."

#cd ..
#echo "create Event"
#./test_db.py