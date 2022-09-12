terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "<= 4.31.0"
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

data "aws_canonical_user_id" "current" {}

data "aws_caller_identity" "current" {}

data "aws_region" "current" {}