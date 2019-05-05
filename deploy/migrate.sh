#!/usr/bin/env sh

env=$(echo ${DRONE_TAG} | cut -d . -f 4) ./main migrate