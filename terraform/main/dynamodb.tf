resource "aws_sqs_queue" "dynamodb_queue" {
  name                       = "${var.deployment_id}-dynamodb-queue"
  visibility_timeout_seconds = 300
}

resource "aws_sqs_queue_policy" "dynamodb_queue_sns_lambda_policy" {
  queue_url = aws_sqs_queue.dynamodb_queue.id
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
        "Resource": "${aws_sqs_queue.dynamodb_queue.arn}",
        "Condition": {
          "ArnEquals": {
            "aws:SourceArn": "${aws_sns_topic.sns_fuzz_results.arn}"
          }
        }
      }, 
      {
        "Sid": "sqs-lambda-dynamodb-policy",
        "Effect": "Allow",
        "Action": "sqs:SendMessage",
        "Resource": "${aws_sqs_queue.dynamodb_queue.arn}", 
        "Principal": {
          "Service": "lambda.amazonaws.com"
        }
      }
    ]
  }
  POLICY
}

resource "aws_sns_topic_subscription" "dynamodb_sns_sqs" {
  topic_arn = aws_sns_topic.sns_fuzz_results.arn
  protocol  = "sqs"
  endpoint  = aws_sqs_queue.dynamodb_queue.arn
}

resource "aws_iam_role" "dynamodb_lambda_role" {
  name               = "${var.deployment_id}-dynamodb-lambda-role"
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

resource "aws_iam_policy" "dynamodb_lambda_policy" {
  name        = "${var.deployment_id}-dynamodb-lambda-policy"
  description = "${var.deployment_id}-dynamodb-lambda-policy"

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
        Resource = "${aws_sqs_queue.dynamodb_queue.arn}"
      },
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:*:*:*"
      },
      {
        Action = [
          "s3:PutObject",
          "s3:PutObjectAcl",
          "s3:GetObject",
          "s3:GetObjectAcl",
          "s3:ListAllMyBuckets"
        ]
        Effect   = "Allow"
        Resource = "${aws_s3_bucket.fuzz_results_bucket.arn}"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "dynamodb_lambda_logs_sqs_policy_attachment" {
  role       = aws_iam_role.dynamodb_lambda_role.id
  policy_arn = aws_iam_policy.dynamodb_lambda_policy.arn
}

data "archive_file" "dynamodb_lambda" {
  type        = "zip"
  source_file = "${path.module}/../../bin/dynamodb/main"
  output_path = "${path.module}/files/dynamodb.zip"
}

resource "aws_lambda_function" "dynamodb_lambda" {
  function_name    = "${var.deployment_id}-dynamodb-lambda"
  filename         = "${path.module}/files/dynamodb.zip"
  source_code_hash = data.archive_file.dynamodb_lambda.output_base64sha256

  runtime = "go1.x"
  handler = "main"

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.dynamodb_lambda_role.arn
}

resource "aws_lambda_event_source_mapping" "dynamodb_sqs_lambda_event_source_mapping" {
  batch_size       = 1
  event_source_arn = aws_sqs_queue.dynamodb_queue.arn
  enabled          = true
  function_name    = aws_lambda_function.dynamodb_lambda.arn
}

resource "aws_cloudwatch_log_group" "dynamodb_lambda_loggroup" {
  name              = "/aws/lambda/${aws_lambda_function.dynamodb_lambda.function_name}"
  retention_in_days = 14
}