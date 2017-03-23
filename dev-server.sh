#!/usr/bin/env bash

pushd server
godep restore
gox -osarch="Linux/amd64" -output="../dist/outrigger-dashboard"
popd

echo "Starting Outrigger Dashboard Server"
docker stop outrigger-dashboard
docker rm outrigger-dashboard
docker run -it \
    --name outrigger-dashboard \
    -e DOCKER_API_VERSION=1.24 \
    -l com.dnsdock.name=dashboard \
    -l com.dnsdock.image=outrigger \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -v `pwd`/dist/outrigger-dashboard:/outrigger-dashboard \
    --rm \
    outrigger/dashboard