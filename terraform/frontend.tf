locals {
  s3_origin_id = "${var.deployment_id}-S3-origin-react-app"
  mime_type_mappings = {
    html = "text/html",
    js   = "text/javascript",
    css  = "text/css"
  }
  frontend_build_path = "${path.module}/../frontend/build"
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
      type = "AWS"
      identifiers = [
        aws_cloudfront_origin_access_identity.oai.iam_arn,
      ]
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

resource "aws_s3_object" "frontend_settings_object" {
  key = "settings.json"
  content = jsonencode({
    backend = aws_apigatewayv2_stage.lambda.invoke_url
  })
  bucket       = aws_s3_bucket.static_react_bucket.bucket
  content_type = "application/json"
}

resource "aws_cloudfront_origin_access_identity" "oai" {
  comment = "my-react-app OAI"
}

resource "aws_cloudfront_distribution" "frontend" {
  origin {
    domain_name = aws_s3_bucket.static_react_bucket.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.oai.cloudfront_access_identity_path
    }
  }

  enabled         = true
  is_ipv6_enabled = true

  default_root_object = "index.html"

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.s3_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
    compress               = true
  }

  ordered_cache_behavior {
    path_pattern     = "/index.html"
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = local.s3_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    min_ttl                = 0
    default_ttl            = 0
    max_ttl                = 0
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }

  price_class = "PriceClass_100"

  viewer_certificate {
    cloudfront_default_certificate = true
    minimum_protocol_version       = "TLSv1.2_2021"
  }

  retain_on_delete = true

  custom_error_response {
    error_caching_min_ttl = 300
    error_code            = 403
    response_code         = 200
    response_page_path    = "/index.html"
  }

  custom_error_response {
    error_caching_min_ttl = 300
    error_code            = 404
    response_code         = 200
    response_page_path    = "/index.html"
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  logging_config {
    include_cookies = false
    bucket          = aws_s3_bucket.cloudfront_logs.bucket_domain_name
  }

  wait_for_deployment = false
}
