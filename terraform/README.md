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

1. Install terraform and jq 

2. If and only if this is the first time creating your resouces or after a full destyoy 

First set a name for your infrastructure, this will be prefixed to any resources yoru create.
Your name lower case is a safe bet
```
read -p "Enter A Prefix To Track Your Infrastructure: " infrastructure_prefix
```
Then provide a existing ecr image to use.
```
temp_restler_img=117712065617.dkr.ecr.ap-southeast-2.amazonaws.com/main_dyno_image_repository:`aws ecr describe-images --repository-name main_dyno_image_repository --query 'sort_by(imageDetails,& imagePushedAt)[-1].imageTags[0]' | tr -d '"'`
```

Then the follwoing lines, 
1. create the ecr
2. move the resouces in the root terrform
3. Create the remaining resouces

```
cd terraform/ecr
terraform init
terraform apply -var="deployment_id=$infrastructure_prefix" -auto-approve
ecr_kms_named=$(terraform show -json | jq '.values.root_module.resources[] | select(.address=="aws_kms_key.ecr_kms_named") | .values.id' | tr -d '"')
ecr_kms_alias_named=$(terraform show -json | jq '.values.root_module.resources[] | select(.address=="aws_kms_alias.ecr_kms_alias_named") | .values.id' | tr -d '"')
image_repository_named=$(terraform show -json | jq '.values.root_module.resources[] | select(.address=="aws_ecr_repository.image_repository_named") | .values.id' | tr -d '"')
cd ../
terraform init
terraform import -var "deployment_id=$infrastructure_prefix" -var "restler_image_tag=$temp_restler_img" module.ecr.aws_kms_key.ecr_kms_named "$ecr_kms_named" 
terraform import -var "deployment_id=$infrastructure_prefix" -var "restler_image_tag=$temp_restler_img" module.ecr.aws_kms_alias.ecr_kms_alias_named "$ecr_kms_alias_named"
terraform import -var "deployment_id=$infrastructure_prefix" -var "restler_image_tag=$temp_restler_img" module.ecr.aws_ecr_repository.image_repository_named "$image_repository_named"
terraform apply -var "deployment_id=$infrastructure_prefix" -var "restler_image_tag=$temp_restler_img" -auto-approve
```
## Development

Create a `build` directory with the artefacts from the UI and a `bin` directory with the artefacts from the backend.
For convenience create a symlink.
