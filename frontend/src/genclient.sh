#!/bin/sh
rm -rf src/client
openapi \
  --input  ../openapi.yml \
  --output ./src/client \
  --client fetch \
  --indent 2 \
  --name   AppClient
