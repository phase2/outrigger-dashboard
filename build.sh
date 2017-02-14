#!/usr/bin/env bash

set -e

pushd server
godep restore
gox -osarch="Linux/amd64" -output="../dist/outrigger-dashboard"
popd

pushd frontend
npm run build
popd

docker build -t outrigger/dashboard .
