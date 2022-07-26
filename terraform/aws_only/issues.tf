# Github Issues Architecture 
# Includes Restler Github Issues Queue -> Github Issues Lambda -> Github Issues API


data "archive_file" "github_issues_lambda" {
  type        = "zip"
  source_file = "${path.module}/../../bin/issues/main"
  output_path = "${path.module}/files/issues.zip"
}

# Github Issues Queue
resource "aws_sqs_queue" "github_issues_queue" {
  name                      = "${var.deployment_id}-github_issues_queue"
  delay_seconds             = 90
  message_retention_seconds = 86400
  max_message_size          = 2048
  receive_wait_time_seconds = 10
  sqs_managed_sse_enabled   = true
  policy                    = <<POLICY
  {
    "Version": "2012-10-17",
    "Id": "${var.deployment_id}.s3-interaction-sqs-github-issues",
    "Statement": [
      {
        "Sid": "sqs-github-issues-statement-id",
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
  tags = {
    name = "github_issues_queue"
  }
}

# SQS Policy for Lambda Function 
resource "aws_sqs_queue_policy" "github_issues_queue_policy" {
  queue_url = aws_sqs_queue.github_issues_queue.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Sid      = ""
      Action   = ["sqs:SendMessage", "sqs:ReceiveMessage", "sqs:DeleteMessage", "sqs:GetQueueAttributes"]
      Effect   = "Allow"
      Resource = "${aws_sqs_queue.github_issues_queue.arn}"
      Principal = {
        Service = "lambda.amazonaws.com"
      }

    }]
  })
}

# # Github Issues Lambda
# Lambda function policy
resource "aws_iam_policy" "github_issues_lambda_policy" {
  name        = "${var.deployment_id}-github-issues-lambda-policy"
  description = "${var.deployment_id}-github-issues-lambda-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "sqs:ReceiveMessage",
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes"
        ]
        Effect   = "Allow"
        Resource = "${aws_sqs_queue.github_issues_queue.arn}"
      },
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

# Lambda function role
resource "aws_iam_role" "iam_github_issues_lambda" {
  name = "${var.deployment_id}-github-issues-lambda-role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
        Effect = "Allow"
      }
    ]
  })
}

# Role to Policy attachment
resource "aws_iam_role_policy_attachment" "terraform_lambda_iam_policy_basic_execution" {
  role       = aws_iam_role.iam_github_issues_lambda.id
  policy_arn = aws_iam_policy.github_issues_lambda_policy.arn
}

# Lambda function declaration
resource "aws_lambda_function" "github_issues_lambda" {
  function_name    = "${var.deployment_id}-github-issues-lambda"
  filename         = "files/issues.zip"
  source_code_hash = data.archive_file.github_issues_lambda.output_base64sha256

  runtime = "go1.x"
  handler = "main"

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.iam_github_issues_lambda.arn
}

# Trigger 
resource "aws_lambda_event_source_mapping" "event_source_mapping" {
  batch_size       = 1
  event_source_arn = aws_sqs_queue.github_issues_queue.arn
  enabled          = true
  function_name    = aws_lambda_function.github_issues_lambda.arn
}

# CloudWatch Log Group for the Lambda function
resource "aws_cloudwatch_log_group" "lambda_loggroup" {
  name              = "/aws/lambda/${aws_lambda_function.github_issues_lambda.function_name}"
  retention_in_days = 14
}