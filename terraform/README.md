# Infrastructure

The Infrastructure takes the build artefacts from the backend and UI projects and deploys them to
AWS using terraform.
If you are deploying on AWS for the first time run Module ecr first and push an image into it which can
be used by the Restler lambda.

## Module Discrption

### Main
Most resources, can be deployed onto localstack

### Ecr
Just the ECR needs to be deployed first on a fresh aws account

### AWS_only
Resources that can not be tested on local stack

## Install

1. Install terraform


2. Initialise terraform

```
terraform init \
  -backend-config="bucket=go-lambda-skeleton-state" \
  -backend-config="key=main" \
  -backend-config="region=eu-west-1" \
  -backend-config="encrypt=true" \
  -backend-config="dynamodb_table=go-lambda-skeleton-state-lock";
```

## Development

Create a `build` directory with the artefacts from the UI and a `bin` directory with the artefacts from the backend.
For convenience create a symlink.
