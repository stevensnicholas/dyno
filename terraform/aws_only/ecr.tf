<<<<<<< HEAD
resource "aws_kms_key" "ecr_kms" {
  enable_key_rotation = true
}


resource "aws_ecr_repository" "image_repository" {
  name                 = "dyno_image_repository"
  image_tag_mutability = "IMMUTABLE"
=======
resource "aws_ecr_repository" "image_repository" {
  name                 = "dyno_image_repository"
  image_tag_mutability = "MUTABLE"
>>>>>>> Add ecr so we can push docker images for lambda

  image_scanning_configuration {
    scan_on_push = true
  }
<<<<<<< HEAD

  encryption_configuration {
    encryption_type = "KMS"
    kms_key         = aws_kms_key.ecr_kms.key_id
  }
}
=======
}
>>>>>>> Add ecr so we can push docker images for lambda
