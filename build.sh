#!/bin/env bash

# Create and run auth microservice
cp config-local.yml ./auth-microservice/config.yml
docker stop auth-microservice || true
docker container prune -f || true
# docker rmi $(docker images | grep '/auth-microservice') || true
docker build -t auth-microservice ./auth-microservice
docker run -d -p 3001:3001 --network host --name auth-microservice auth-microservice:latest
# rm ./auth-microservice/config.yml

# Create and run article microservice
cp config-local.yml ./article-microservice/config.yml
docker stop article-microservice || true
docker container prune -f || true
# docker rmi $(docker images | grep '/article-microservice') || true
docker build -t article-microservice ./article-microservice
docker run -d -p 3002:3002 --network host --name article-microservice article-microservice:latest
# rm ./article-microservice/config.yml

# Create and run gateway microservice
cp config-local.yml ./gateway-microservice/config.yml
docker stop gateway-microservice || true
docker container prune -f || true
# docker rmi $(docker images | grep '/gateway-microservice') || true
docker build -t gateway-microservice ./gateway-microservice
docker run -d -p 3000:3000 --network host --name gateway-microservice gateway-microservice:latest
# rm ./gateway-microservice/config.yml

docker stop db || true
docker container prune -f || true
docker build -t postgres_db ./db
docker run -d -p 5431:5432 --name db postgres_db:latest

# remove dangling images
docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi