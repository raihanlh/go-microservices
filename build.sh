#!/bin/env bash

# Create and run auth microservice
cp config.yml ./auth-microservice
docker stop auth-microservice || true
docker container prune -f || true
# docker rmi $(docker images | grep '/auth-microservice') || true
docker build -t auth-microservice ./auth-microservice
docker run -d -p 3001:3001 --network host --name auth-microservice auth-microservice:latest
rm ./auth-microservice/config.yml

# Create and run gateway microservice
cp config.yml ./gateway-microservice
docker stop gateway-microservice || true
docker container prune -f || true
# docker rmi $(docker images | grep '/gateway-microservice') || true
docker build -t gateway-microservice ./gateway-microservice
docker run -d -p 3001:3001 --network host --name gateway-microservice gateway-microservice:latest
rm ./gateway-microservice/config.yml

# remove dangling images
docker images -a | grep none | awk '{ print $3; }' | xargs docker rmi