<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Add Kms ley
=======
>>>>>>> Add ecr so we can push docker images for lambda
resource "aws_kms_key" "ecr_kms" {
  enable_key_rotation = true
}


<<<<<<< HEAD
resource "aws_ecr_repository" "image_repository" {
  name                 = "dyno_image_repository"
  image_tag_mutability = "IMMUTABLE"
=======
<<<<<<< HEAD
=======
>>>>>>> Add Kms ley
=======
>>>>>>> Add Kms ley
resource "aws_ecr_repository" "image_repository" {
  name                 = "dyno_image_repository"
  image_tag_mutability = "MUTABLE"
<<<<<<< HEAD
>>>>>>> Add ecr so we can push docker images for lambda
=======
  image_tag_mutability = "IMMUTABLE"
>>>>>>> Fix secreity issue - make immutable
<<<<<<< HEAD
=======
resource "aws_ecr_repository" "image_repository" {
  name                 = "dyno_image_repository"
  image_tag_mutability = "MUTABLE"
>>>>>>> Add ecr so we can push docker images for lambda
>>>>>>> Add ecr so we can push docker images for lambda
=======
>>>>>>> Fix secreity issue - make immutable
=======
>>>>>>> Add Kms ley

  image_scanning_configuration {
    scan_on_push = true
  }
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Add Kms ley
=======
>>>>>>> Add ecr so we can push docker images for lambda
=======
>>>>>>> Fix secreity issue - make immutable

  encryption_configuration {
    encryption_type = "KMS"
    kms_key         = aws_kms_key.ecr_kms.key_id
  }
<<<<<<< HEAD
}
=======
}
>>>>>>> Add ecr so we can push docker images for lambda
=======
=======
>>>>>>> Add Kms ley
}
<<<<<<< HEAD
>>>>>>> Fix secreity issue - make immutable
=======
=======
}
>>>>>>> Add ecr so we can push docker images for lambda
<<<<<<< HEAD
>>>>>>> Add ecr so we can push docker images for lambda
=======
=======
}
>>>>>>> Fix secreity issue - make immutable
>>>>>>> Fix secreity issue - make immutable
=======
}
>>>>>>> Add Kms ley
