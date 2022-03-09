#!/usr/bin/env bash
set -ex
cd cmd/easypin
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . 
cd -
cd web
npx vite build
cd -
docker build -t ghcr.io/elek/easypin .
