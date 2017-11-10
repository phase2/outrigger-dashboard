#!/usr/bin/env bash

pushd server
dep ensure
GOOS=linux GOOS=amd64 go build -o ../dist/outrigger-dashboard
popd

echo "Starting Outrigger Dashboard Server"
docker stop outrigger-dashboard
docker rm outrigger-dashboard
docker run -it \
    --name outrigger-dashboard \
    -e DOCKER_API_VERSION=1.32 \
    -l com.dnsdock.name=dashboard \
    -l com.dnsdock.image=outrigger \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v `pwd`/dist/outrigger-dashboard:/outrigger-dashboard \
    --rm \
    outrigger/dashboard
