# Github Issues Architecture 
# Includes Github Issues Queue -> Github Issues Lambda -> Github Issues API
data "archive_file" "lambda" {
  type        = "zip"
  source_file = "${path.module}/../bin/main"
  output_path = "${path.module}/files/issues.zip"
}

# Github Issues Queue
resource "aws_sqs_queue" "github_issues_queue" {
  name                      = "github_issues_queue.fifo"
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
  fifo_queue                  = true
  content_based_deduplication = true

  tags = {
    name = "github_issues_queue"
  }
}

resource "aws_sqs_queue_policy" "github_issues_queue_policy" {
  queue_url = aws_sqs_queue.github_issues_queue.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Sid = ""
      Action = "sqs:SendMessage"
      Effect = "Allow"
      Resource = "${aws_sqs_queue.github_issues_queue.arn}"
      Principal = {
        Service = "lambda.amnazonaws.com"
      }
      
    }]
  })
}

# Github Issues Lambda
resource "aws_iam_role" "iam_github_issues_lambda" {
  name = "iam_github_issues_lambda"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Sid = ""
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "lambda.amnazonaws.com"
      }
    }]
  })
}

resource "aws_lambda_function" "github_issues_lambda" {
  filename = "files/issues.zip"
  function_name = "${var.deployment_id}-github-issues"
  role = aws_iam_role.iam_github_issues_lambda.arn
  handler = 
  source_code_hash = "${base64sha256(file("issues.zip"))}"
  runtime = "go1.x"

  tracing_config {
    mode = "Active"
  }
}

# Event source from SQS
resource "aws_lambda_event_source_mapping" "event_source_mapping" {
  event_source_arn = "${var.terraform_queue_arn}"
  enabled          = true
  function_name    = "${aws_lambda_function.test_lambda.arn}"
  batch_size       = 1
}
