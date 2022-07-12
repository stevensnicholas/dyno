locals {
  s3_origin_id = "${var.deployment_id}-S3-origin-react-app"
  mime_type_mappings = {
    html = "text/html",
    js   = "text/javascript",
    css  = "text/css"
  }
  frontend_build_path = "${path.module}/../../frontend/build"
}

resource "aws_s3_bucket" "static_react_bucket" {
  bucket = "${var.deployment_id}-go-lambda-skeleton-frontend"
}

resource "aws_s3_bucket_versioning" "static_react_bucket" {
  bucket = aws_s3_bucket.static_react_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_acl" "static_react_bucket" {
  bucket = aws_s3_bucket.static_react_bucket.id
  acl    = "private"
}

resource "aws_s3_bucket_public_access_block" "static_react_bucket" {
  bucket = aws_s3_bucket.static_react_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}


