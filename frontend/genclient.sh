#!/bin/sh
rm -rf src/client
openapi \
  --input  ../api/openapi.yml \
  --output ./src/client \
  --client fetch \
  --indent 2 \
  --name   AppClient
