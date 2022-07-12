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

data "aws_iam_policy_document" "static_react_bucket" {
  statement {
    sid = "AllowCloudFront"

    actions = [
      "s3:GetObject"
    ]

    resources = [
      "${aws_s3_bucket.static_react_bucket.arn}/*"
    ]

    principals {
      type        = "AWS"
      identifiers = ["*", ]
    }
  }
}

resource "aws_s3_bucket_policy" "static_react_bucket" {
  bucket = aws_s3_bucket.static_react_bucket.id
  policy = data.aws_iam_policy_document.static_react_bucket.json
}

resource "aws_s3_object" "frontend_object" {
  for_each = fileset(local.frontend_build_path, "**")

  key    = each.value
  source = "${local.frontend_build_path}/${each.value}"
  bucket = aws_s3_bucket.static_react_bucket.bucket

  etag         = filemd5("${local.frontend_build_path}/${each.value}")
  content_type = lookup(local.mime_type_mappings, concat(regexall("\\.([^\\.]*)$", each.value), [[""]])[0][0], "application/octet-stream")
}


