# Infrastructure

This module contains the core components and can be run in isolation  with free locastack resouces.


## Install

1. Install terraform

2. Install tflocal script 
```
pip install terraform-local
```

3. Make sure the docker compose is up running the localstack docker

4. Int and deploy
```
tflocal init
tflocal apply
```
5. Check that it is running by listing local lambda functions (requires aws cli)
```
aws lambda --endpoint http://localhost:4566 list-functions
```
## Development

Create a `build` directory with the artefacts from the UI and a `bin` directory with the artefacts from the backend.
For convenience create a symlink.
