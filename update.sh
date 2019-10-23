#!/bin/bash
app=gin
dbname=mysql8.0
port=9000
docker rm -f $app
docker rmi $app
docker build -t $app .
docker run -d  -v /etc/localtime:/etc/localtime -p $port:$port --name $app --link $dbname:db $app:latest