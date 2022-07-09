data "archive_file" "lambda" {
  type        = "zip"
  source_file = "${path.module}/../bin/main"
  output_path = "${path.module}/files/backend.zip"
}

resource "aws_iam_role" "lambda_exec" {
  name = "${var.deployment_id}-lambda-execution-comp9447"

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

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "lambda" {
  function_name    = "${var.deployment_id}-server"
  filename         = "files/backend.zip"
  source_code_hash = data.archive_file.lambda.output_base64sha256

  runtime = "go1.x"
  handler = "main"

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_cloudwatch_log_group" "lambda" {
  name              = "/aws/lambda/${aws_lambda_function.lambda.function_name}"
  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "lambda" {
  count = "${var.is_local ? 0 : 1}"
  api_id = aws_apigatewayv2_api.gateway[0].id

  integration_uri    = aws_lambda_function.lambda.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "lambda" {
  count = "${var.is_local ? 0 : 1}"
  api_id = aws_apigatewayv2_api.gateway[0].id

  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.lambda[0].id}"
}

resource "aws_lambda_permission" "api_gw" {
  count = "${var.is_local ? 0 : 1}"
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.gateway[0].execution_arn}/*/*"
}
