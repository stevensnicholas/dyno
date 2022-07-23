#!/bin/sh
if [ -z "${AWS_LAMBDA_RUNTIME_API}" ]; then
# running locally in docker
  exec aws-lambda-rie python3 -m awslambdaric $@
else
# running on aws
  exec python3 -m awslambdaric $@
fi   