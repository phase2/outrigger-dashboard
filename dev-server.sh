#!/usr/bin/env bash

pushd server
gox -osarch="Linux/amd64" -output="../dist/outrigger-dashboard"
popd

docker rm outrigger-dashboard

echo "Starting Outrigger Dashboard Server"
docker run -it \
    --name outrigger-dashboard \
    -e DOCKER_HOST \
    -e DOCKER_TLS_VERIFY \
    -e DOCKER_CERT_PATH \
    -l com.dnsdock.name=dashboard \
    -l com.dnsdock.image=outrigger \
    -v $HOME:$HOME \
    -v `pwd`/dist/outrigger-dashboard:/outrigger-dashboard \
    phase2/outrigger-dashboard
