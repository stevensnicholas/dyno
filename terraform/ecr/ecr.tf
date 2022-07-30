resource "aws_kms_key" "ecr_kms_named" {
  enable_key_rotation = true
}

resource "aws_kms_alias" "ecr_kms_alias_named" {
  name          = "alias/${var.deployment_id}_ecr_kms_alias"
  target_key_id = aws_kms_key.ecr_kms_named.key_id
}

resource "aws_ecr_repository" "image_repository_named" {
  name                 = "${var.deployment_id}_dyno_image_repository"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  encryption_configuration {
    encryption_type = "KMS"
    kms_key         = aws_kms_alias.ecr_kms_alias_named.target_key_arn
  }
}
