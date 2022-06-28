# Golang Lambda React Template

The Golang Lambda React Template is a web application which deploys a golang API to AWS lambda
and a React frontend to AWS CloudFront / S3.

## Install

Dependencies:
- golang compiler
- node
- npm
- docker
- terraform

Install code dependencies
```
go mod download
cd frontend
npm ci
```

## Build

Compile the backend into the `bin` directory
```
make
```

Compile the frontend into the `frontend/build` directory
```
cd frontend
npm run build
```

## Run

Usage:
```
./bin/main --help
Usage of ./bin/main:
  -gendocs
    	Generate API Docs
  -hostdocs
    	Host API Docs
  -http
    	Run an HTTP server instead of in AWS Lambda
  -loglevel string
    	Log level (default "info")
```

Run backend http server
```
./bin/main -http
```

Run frontend development server
```
cd frontend
npm start
```

## Deploy

Deploys the frontend and backend to AWS cloudfront and lambda
```
cd terraform
terraform init
terraform apply
```

Github pipelines will deploy when the master and production branches are updated

## Test

```
make test # without coverage
make cov  # with coverage
```

## Lint

Run the backend linter\
Runs the golangci-lint docker image with the following tools:
- gofmt
- stylecheck
- gosec
```
make lint
```

Run the frontend linter
```
cd frontend
npm run lint
npm run fix # auto fixes problems
```
