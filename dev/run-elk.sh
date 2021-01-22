#!/bin/bash
set -a

root_dir="$(dirname $(dirname $(realpath "$0")))"
export APP_PATH="${root_dir}/app/bin"
export CONFIG_PATH="${root_dir}/filebeat/config/local"

echo "APP_PATH = $APP_PATH, CONFIG_PATH = $CONFIG_PATH"
pushd "$(dirname $(realpath $0))"

rm -f $APP_PATH/app.log

echo "Starting up..."
docker-compose up -V