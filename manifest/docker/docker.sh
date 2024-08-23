#!/bin/bash

# This shell is executed before docker build.
# loop over all folder in internal and build docker image

# create variable internal
INTERNAL_DIR=$1
DOCKERFILE=$2
APP_DIR=$3
# len of internal
len=${#INTERNAL_DIR}

for d in ${INTERNAL_DIR}/*; do
    if [ -d "$d" ]; then
        echo "Building ${d:len+1} using Dockerfile"

        docker build -t 1layar/${d:len+1} -f ${DOCKERFILE} --build-arg APP_DIR=${ROOT_DIR} --build-arg SERVICE_NAME=${d:len+1} ${APP_DIR}
    fi
done





