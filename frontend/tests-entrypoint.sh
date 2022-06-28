#!/bin/sh
set -e
echo "{ \"backend\": \"$BACKEND_HOST\" }" > /usr/share/nginx/html/settings.json
echo "running docker-entrypoint"
nginx -g "daemon off;"
