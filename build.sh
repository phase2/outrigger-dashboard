#!/usr/bin/env bash

set -e

pushd server
gox -osarch="Linux/amd64" -output="../dist/outrigger-dashboard"
popd

pushd frontend
npm run build
popd

docker build -t phase2/outrigger-dashboard .
