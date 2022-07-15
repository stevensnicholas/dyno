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
>>>>>>> 427130f36b90e70fc6dfdb0478bf88b8f35952bb

  image_scanning_configuration {
    scan_on_push = true
  }
<<<<<<< HEAD

  encryption_configuration {
    encryption_type = "KMS"
    kms_key         = aws_kms_key.ecr_kms.key_id
  }
=======
>>>>>>> 427130f36b90e70fc6dfdb0478bf88b8f35952bb
}
