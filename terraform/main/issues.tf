# Issues Architecture 
# Includes Restler SNS ->  Issues Queue ->  Issues Lambda ->  Issues API
resource "aws_kms_key" "fuzz_results_key" {
  description         = "fuzz-results-topic-key"
  policy              = <<POLICY
  {
    "Version": "2012-10-17",
    "Id": "key-issues",
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
          "kms:Encrypt",
          "kms:Decrypt"
        ],
        "Resource": "*"
      }
    ]
  }
  POLICY
  enable_key_rotation = true
}

resource "aws_kms_alias" "fuzz_results_key_alias" {
  name          = "alias/fuzz-results-key"
  target_key_id = aws_kms_key.fuzz_results_key.key_id
}

resource "aws_sns_topic" "sns_fuzz_results" {
  name = "${var.deployment_id}-sns-fuzz-results"
  # kms_master_key_id = aws_kms_alias.fuzz_results_key_alias.id
  #tfsec:ignore
  policy = <<POLICY
    {
      "Version":"2012-10-17",
      "Statement":[
        {
          "Effect": "Allow",
          "Principal": {"Service":"s3.amazonaws.com"},
          "Action": "SNS:Publish",
          "Resource":  "arn:aws:sns:*:*:${var.deployment_id}-sns-fuzz-results",
          "Condition":{
              "ArnLike":{"aws:SourceArn":"${aws_s3_bucket.openapi_files_bucket.arn}"}
          }
        }
      ]
    }
  POLICY
}

resource "aws_sqs_queue" "issues_queue" {
  name                       = "${var.deployment_id}-issues-queue"
  visibility_timeout_seconds = 300
  # kms_master_key_id = aws_kms_alias.fuzz_results_key_alias.id
  #tfsec:ignore

}

resource "aws_sqs_queue_policy" "issues_queue_sns_lambda_policy" {
  queue_url = aws_sqs_queue.issues_queue.id
  policy    = <<POLICY
  {
    "Version": "2012-10-17",
    "Id": "sqspolicy",
    "Statement": [
      {
        "Sid": "sqs-sns-policy",
        "Effect": "Allow",
        "Principal": "*",
        "Action": "sqs:SendMessage",
        "Resource": "${aws_sqs_queue.issues_queue.arn}",
        "Condition": {
          "ArnEquals": {
            "aws:SourceArn": "${aws_sns_topic.sns_fuzz_results.arn}"
          }
        }
      }, 
      {
        "Sid": "sqs-lambda-issues-policy",
        "Effect": "Allow",
        "Action": "sqs:SendMessage",
        "Resource": "${aws_sqs_queue.issues_queue.arn}", 
        "Principal": {
          "Service": "lambda.amazonaws.com"
        }
      }
    ]
  }
  POLICY
}

resource "aws_sns_topic_subscription" "issues_sns_sqs" {
  topic_arn = aws_sns_topic.sns_fuzz_results.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.issues_queue.arn
  #tfsec:ignore
}

resource "aws_iam_role" "issues_lambda_role" {
  name               = "${var.deployment_id}-issues-lambda-role"
  assume_role_policy = <<POLICY
  {
    "Version": "2012-10-17",
    "Statement": [
      {
          "Action": "sts:AssumeRole",
          "Effect": "Allow",
          "Principal": {
              "Service": "lambda.amazonaws.com"
          }
      }
    ]
  }
  POLICY
}

resource "aws_iam_policy" "issues_lambda_policy" {
  name        = "${var.deployment_id}-issues-lambda-policy"
  description = "${var.deployment_id}-issues-lambda-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "sqs:ReceiveMessage",
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes",
          "sqs:ChangeMessageVisibility"
        ]
        Effect   = "Allow"
        Resource = "${aws_sqs_queue.issues_queue.arn}"
      },
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:*:*:*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "issues_lambda_logs_sqs_policy_attachment" {
  role       = aws_iam_role.issues_lambda_role.id
  policy_arn = aws_iam_policy.issues_lambda_policy.arn
}

data "archive_file" "issues_lambda" {
  type        = "zip"
  source_file = "${path.module}/../../bin/issues/main"
  output_path = "${path.module}/files/issues.zip"
}

resource "aws_lambda_function" "issues_lambda" {
  function_name    = "${var.deployment_id}-issues-lambda"
  filename         = "${path.module}/files/issues.zip"
  source_code_hash = data.archive_file.issues_lambda.output_base64sha256

  runtime = "go1.x"
  handler = "main"

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.issues_lambda_role.arn
}

resource "aws_lambda_event_source_mapping" "issues_sqs_lambda_event_source_mapping" {
  batch_size       = 1
  event_source_arn = aws_sqs_queue.issues_queue.arn
  enabled          = true
  function_name    = aws_lambda_function.issues_lambda.arn
}

resource "aws_cloudwatch_log_group" "issues_lambda_loggroup" {
  name              = "/aws/lambda/${aws_lambda_function.issues_lambda.function_name}"
  retention_in_days = 14
}