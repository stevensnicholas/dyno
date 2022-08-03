terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "<= 4.25.0"
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

resource "aws_kms_key" "ecr_kms" {
  enable_key_rotation = true
}

resource "aws_kms_alias" "ecr_kms_alias" {
  name          = "alias/${var.deployment_id}_ecr_kms_alias_demoserver"
  target_key_id = aws_kms_key.ecr_kms.key_id
}

resource "aws_ecr_repository" "image_repository" {
  name                 = "${var.deployment_id}_demoserver_image_repository"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  encryption_configuration {
    encryption_type = "KMS"
    kms_key         = aws_kms_alias.ecr_kms_alias.target_key_arn
  }
}

resource "null_resource" "tag_and_push" {
  provisioner "local-exec" {
    command = <<EOF
    docker tag ${var.demo_server_tag} ${aws_ecr_repository.image_repository.repository_url}:latest
    docker push ${aws_ecr_repository.image_repository.repository_url}:latest
    EOF
  }
}

resource "aws_s3_bucket" "demo_server" {
  bucket_prefix = "${var.deployment_id}-demo-server-application"
}

resource "aws_s3_bucket_versioning" "demo_server" {
  bucket = aws_s3_bucket.demo_server.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_public_access_block" "demo_server" {
  bucket = aws_s3_bucket.demo_server.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}


resource "aws_s3_object" "demo_server" {
  depends_on = [
    null_resource.tag_and_push
  ]
  bucket  = aws_s3_bucket.demo_server.id
  key     = "beanstalk/docker-compose.yml"
  content = yamlencode({ "version" : "3.8", "services" : { "demo_server" : { "image" : "${aws_ecr_repository.image_repository.repository_url}:latest", "ports" : ["80:8888"] } } })
}



resource "aws_elastic_beanstalk_application" "demo_server" {
  name        = "${var.deployment_id}-demo-server-application"
  description = "Demo server to test RESTLer aginst"
}

resource "aws_iam_instance_profile" "demo_server" {
  name = "${var.deployment_id}-demo-server-application"
  role = aws_iam_role.demo_server.name
}

resource "aws_iam_role" "demo_server" {
  name = "${var.deployment_id}-demo-server-in-beanstalk"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ec2.amazonaws.com"
      }
    }]
  })
}

locals {
  managed_policies = ["arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly", "arn:aws:iam::aws:policy/AWSElasticBeanstalkWebTier", "arn:aws:iam::aws:policy/AWSElasticBeanstalkMulticontainerDocker", "arn:aws:iam::aws:policy/AWSElasticBeanstalkWorkerTier"]
}

resource "aws_iam_role_policy_attachment" "demo_server" {
  count      = length(local.managed_policies)
  policy_arn = local.managed_policies[count.index]
  role       = aws_iam_role.demo_server.name
}

resource "aws_elastic_beanstalk_environment" "demo_server" {
  name                = "${var.deployment_id}-demo-server-environment"
  application         = aws_elastic_beanstalk_application.demo_server.name
  solution_stack_name = "64bit Amazon Linux 2 v3.4.17 running Docker"
  version_label       = aws_elastic_beanstalk_application_version.demo_server.name
  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "IamInstanceProfile"
    value     = aws_iam_instance_profile.demo_server.name
  }
}

resource "aws_elastic_beanstalk_application_version" "demo_server" {
  name        = "${var.deployment_id}-demo-server-application-version"
  application = aws_elastic_beanstalk_application.demo_server.name
  description = "Demo server to test RESTLer aginst"
  bucket      = aws_s3_bucket.demo_server.id
  key         = aws_s3_object.demo_server.id
}