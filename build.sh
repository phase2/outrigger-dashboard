#!/usr/bin/env bash

set -e

pushd server
dep ensure
GOOS=linux GOARCH=amd64 go build -o ../dist/outrigger-dashboard
popd

pushd frontend
npm install
npm run build
popd

docker build -t outrigger/dashboard .
