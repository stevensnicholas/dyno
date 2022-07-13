resource "aws_ecr_repository" "image_repository" {
  name                 = "dyno_image_repository"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
