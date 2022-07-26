resource "aws_s3_bucket" "openapi_files_bucket" {
  bucket = "${var.deployment_id}-client-openapi-files"
  versioning {
    enabled = true
  }
}

resource "aws_s3_bucket_public_access_block" "openapi_files_bucket_access" {
  bucket = aws_s3_bucket.openapi_files_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_kms_key" "openapi_fuzz" {
  enable_key_rotation = true
}

resource "aws_kms_alias" "openapi_fuzz_alias" {
  name          = "${var.deployment_id}_alias/openapi_fuzz_alias"
  target_key_id = aws_kms_key.openapi_fuzz.key_id
}

resource "aws_sqs_queue" "openapi_sqs_queue" {
  name                      = "${var.deployment_id}-openapifiles-queue"
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
  kms_master_key_id         = aws_kms_alias.openapi_fuzz_alias.target_key_arn
  tags = {
    Environment = "production"
  }
}

resource "aws_sqs_queue_policy" "openapi_s3_notify_sqs_policy" {
  queue_url = aws_sqs_queue.openapi_sqs_queue.id

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "example-ID",
  "Statement": [
    {
      "Sid": "example-statement-ID",
      "Effect": "Allow",
      "Principal": {
        "Service": "s3.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "${aws_sqs_queue.openapi_sqs_queue.arn}",
      "Condition": {
        "StringEquals": {
          "aws:SourceAccount": "${var.account_id}"
        },
        "ArnLike": {
          "aws:SourceArn": "${aws_s3_bucket.openapi_files_bucket.id}"
        }
      }
    }
  ]
}
POLICY
}

resource "aws_s3_bucket_notification" "openapi_notify_sqs" {
  bucket = aws_s3_bucket.openapi_files_bucket.id

  queue {
    queue_arn = aws_sqs_queue.openapi_sqs_queue.arn
    events    = ["s3:ObjectCreated:*"]
  }
}
