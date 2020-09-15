#!/bin/bash
set -euo pipefail

IMAGE=$1
cd $IMAGE
docker build -t acadx0/$IMAGE .
docker push acadx0/$IMAGE
