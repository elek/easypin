#!/usr/bin/env bash
set -ex
cd cmd/easypin
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . 
cd -
cd web
npm install
npx vite build
cd -
TAG=$(git log -n 1 --pretty=format:"%h")
docker build -t ghcr.io/elek/easypin:$TAG .

