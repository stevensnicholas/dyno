resource "aws_cloudwatch_log_group" "lambda_restler" {
  name              = "/aws/lambda/${var.deployment_id}-restler-fuzzer-lambda"
  retention_in_days = 30
}

resource "aws_iam_role" "lambda_restler" {
  name = "${var.deployment_id}-lambda-execution-restler-lambda"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

resource "aws_iam_policy" "lambda_policy" {
  name        = "${var.deployment_id}-lambda-custom-policy-restler-lambda"
  description = "restler lambda policy"

  policy = <<EOT
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:ListAllMyBuckets"
      ],
      "Effect": "Allow",
      "Resource": "*"
    },
    {
      "Action": [
        "s3:PutObject",
        "s3:PutObjectAcl",
        "s3:GetObject",
        "s3:GetObjectAcl"
      ],
      "Effect": "Allow",
      "Resource": "${aws_s3_bucket.openapi_files_bucket.arn}"
    },
    {
        "Effect": "Allow",
        "Action": [
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:PutLogEvents"
        ],
        "Resource": "*"
    },
    {
        "Effect": "Allow",
        "Action": [
            "sqs:ReceiveMessage",
            "sqs:DeleteMessage",
            "sqs:GetQueueAttributes"
        ],
        "Resource": "${aws_sqs_queue.openapi_sqs_queue.arn}"
    }
  ]
}
EOT
}
resource "aws_iam_role_policy_attachment" "lambda_restler" {
  policy_arn = aws_iam_policy.lambda_policy.arn
  role       = aws_iam_role.lambda_restler.name
}