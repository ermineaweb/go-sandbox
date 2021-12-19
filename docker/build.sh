#!/bin/bash

DOCKER_APP_NAME="go-app"

docker build -f Dockerfile-build -t $DOCKER_APP_NAME:latest .