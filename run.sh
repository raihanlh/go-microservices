#!/bin/bash

# example usage: ./run.sh auth
cd $1-microservice
go run src/server.go