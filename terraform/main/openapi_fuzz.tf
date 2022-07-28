resource "aws_s3_bucket" "openapi_files_bucket" {
  depends_on = [
    aws_sqs_queue_policy.openapi_s3_notify_sqs_policy
  ]
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
  policy              = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "key-allow-s3",
  "Statement": [
    {
      "Sid": "Enable IAM User Permissions",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
      },
      "Action": "kms:*",
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "s3.amazonaws.com"
      },
      "Action": [
        "kms:GenerateDataKey",
        "kms:Decrypt"
      ],
      "Resource": "*"
    }
  ]
}
POLICY
}

resource "aws_kms_alias" "openapi_fuzz_alias" {
  name          = "alias/${var.deployment_id}_openapi_fuzz_alias"
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
  "Id": "allow s3",
  "Statement": [
    {
      "Sid": "Allow sqs",
      "Effect": "Allow",
      "Principal": {
        "Service": "s3.amazonaws.com"
      },
      "Action": "SQS:SendMessage",
      "Resource": "${aws_sqs_queue.openapi_sqs_queue.arn}",
      "Condition": {
        "ArnLike": {
          "aws:SourceArn": "arn:aws:s3:::${var.deployment_id}-client-openapi-files"
        }
      }
    }
  ]
}
POLICY
}

resource "aws_s3_bucket_notification" "openapi_notify_sqs" {
  bucket = aws_s3_bucket.openapi_files_bucket.id
  depends_on = [
    aws_sqs_queue_policy.openapi_s3_notify_sqs_policy,
    aws_s3_bucket.openapi_files_bucket
  ]

  queue {
    queue_arn = aws_sqs_queue.openapi_sqs_queue.arn
    events    = ["s3:ObjectCreated:*"]
  }
}