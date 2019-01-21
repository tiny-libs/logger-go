#!/bin/bash

docker build -t logger-go-workspace .
docker run -it -v "$PWD"."/../":/usr/src/logger --rm --name logger-go logger-go-workspace