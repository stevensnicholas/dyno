resource "aws_cloudwatch_log_group" "lambda_restler" {
  name              = "/aws/lambda/${var.aws_lambda_function.function_name}"
  retention_in_days = 30
}

resource "aws_lambda_function" "lambda_restler" {
  function_name = "${var.deployment_id}-restler-fuzzer"
  image_uri     = "${aws_ecr_repository.image_repository.repository_url}:latest"

  runtime = "python3.9"
  handler = "app.handler"

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.lambda_exec_test.arn
}