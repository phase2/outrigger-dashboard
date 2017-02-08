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
    -e DOCKER_HOST \
    -e DOCKER_TLS_VERIFY \
    -e DOCKER_CERT_PATH \
    -l com.dnsdock.name=dashboard \
    -l com.dnsdock.image=outrigger \
    -v $HOME:$HOME \
    -v `pwd`/dist/outrigger-dashboard:/outrigger-dashboard \
    --rm \
    phase2/outrigger-dashboard