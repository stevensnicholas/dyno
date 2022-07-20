resource "aws_kms_key" "ecr_kms" {
  enable_key_rotation = true
}

resource "aws_kms_alias" "ecr_kms_alias" {
  name          = "alias/ecr_kms_alias"
  target_key_id = aws_kms_key.ecr_kms.key_id
}

resource "aws_ecr_repository" "image_repository" {
  name                 = "dyno_image_repository"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }

  encryption_configuration {
    encryption_type = "KMS"
    kms_key         = aws_kms_alias.ecr_kms_alias.target_key_arn
  }
}
