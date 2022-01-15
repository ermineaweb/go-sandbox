#!/bin/bash

DOCKER_APP_NAME="go-app-dev"
LOCAL_APP_SRC="/home/romain/workspace/go-sandbox/docker/src"

docker build -f Dockerfile-dev -t $DOCKER_APP_NAME:latest .

watch docker run -v $LOCAL_APP_SRC:/app/src $DOCKER_APP_NAME:latest
