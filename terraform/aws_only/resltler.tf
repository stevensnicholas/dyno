resource "aws_cloudwatch_log_group" "lambda_restler" {
  name              = "/aws/lambda/${var.aws_lambda_function.function_name}"
  retention_in_days = 30
}

resource "aws_iam_role" "lambda_restler" {
  name = "${var.deployment_id}-lambda-execution-restler"

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

resource "aws_iam_role_policy_attachment" "lambda_restler" {
  role       = aws_iam_role.lambda_restler.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "lambda_restler" {
  function_name = "${var.deployment_id}-restler-fuzzer"
  image_uri     = "${aws_ecr_repository.image_repository.repository_url}:latest"
  handler       = "app.handler"
  timeout       = 60

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.lambda_restler.arn
}