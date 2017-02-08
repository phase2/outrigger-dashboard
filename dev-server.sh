#!/usr/bin/env bash

pushd server
gox -osarch="Linux/amd64" -output="../dist/outrigger-dashboard"
popd

docker rm outrigger-dashboard

echo "Starting Outrigger Dashboard Server"
docker-compose up
