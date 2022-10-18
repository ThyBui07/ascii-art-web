#!/bin/bash -e

docker image build -f Dockerfile -t ascii-art-web-img .
docker container run -p 8080:8080 --detach --name ascii-art-web-ctn ascii-art-web-img
