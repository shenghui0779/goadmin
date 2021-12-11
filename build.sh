#!/bin/bash
docker rm -f goadmin
docker rmi -f goadmin
docker build -t goadmin .
docker image prune -f
docker run -d --name=goadmin --privileged -p 10086:8000 -v /data/config/goadmin:/data/config -v /data/logs/goadmin:/data/logs goadmin