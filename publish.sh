#!/bin/bash
cp config.yml ./auth-microservice
cp config.yml ./article-microservice
cp config.yml ./gateway-microservice

docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
docker build -t raihanlh/postgres_db ./db
docker push raihanlh/postgres_db
docker build -t raihanlh/auth-microservice ./auth-microservice
docker push raihanlh/auth-microservice
docker build -t raihanlh/article-microservice ./article-microservice
docker push raihanlh/article-microservice
docker build -t raihanlh/gateway-microservice ./gateway-microservice
docker push raihanlh/gateway-microservice

rm rm ./auth-microservice/config.yml rm ./article-microservice/config.yml ./gateway-microservice/config.yml

# remove dangling images
docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi