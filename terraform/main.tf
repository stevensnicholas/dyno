terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "<= 4.23.0"
    }
  }
}

provider "aws" {
  region = "ap-southeast-2"
  default_tags {
    tags = {
      Deployment = var.deployment_id
    }
  }
}

module "main" {
  source        = "./main"
  deployment_id = var.deployment_id
}

module "ecr" {
  source        = "./ecr"
  deployment_id = var.deployment_id
}

module "aws_only" {
  source                 = "./aws_only"
  deployment_id          = var.deployment_id
  aws_lambda_function    = module.main.aws_lambda_function
  static_react_bucket    = module.main.static_react_bucket
  restler_image_tag      = var.restler_image_tag
  lambda_restler_iam_arn = module.main.lambda_restler_iam_arn
}

data "aws_canonical_user_id" "current" {}
