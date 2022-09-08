#!/bin/env bash

cd proto && make

for name in gateway article auth user comment
do
    cp *.pb.go ../$name-microservice/proto
done
rm *.pb.go