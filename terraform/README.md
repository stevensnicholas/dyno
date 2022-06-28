# Infrastructure

The Infrastructure takes the build artefacts from the backend and UI projects and deploys them to
AWS using terraform.

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
