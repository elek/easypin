#!/usr/bin/env bash
set -ex
TAG=$(git log -n 1 --pretty=format:"%h")
docker push ghcr.io/elek/easypin:$TAG

