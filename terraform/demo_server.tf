# Gatwway

resource "aws_apigatewayv2_api" "gateway_testing" {
  name          = "${var.deployment_id}_gateway_testing_api"
  protocol_type = "HTTP"

  cors_configuration {
    allow_origins = ["https://${aws_cloudfront_distribution.frontend.domain_name}"]
    allow_methods = ["*"]
    allow_headers = ["*"]
  }
}

resource "aws_cloudwatch_log_group" "api_gw_testing_api" {
  name              = "/aws/api_gw/${aws_apigatewayv2_api.gateway_testing.name}"
  retention_in_days = 30
}

resource "aws_apigatewayv2_stage" "lambda_testing_api" {
  api_id = aws_apigatewayv2_api.gateway_testing.id

  name        = "api"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gw.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
      }
    )
  }
}

# Backend
data "archive_file" "lambda_test_api" {
  type        = "zip"
  source_dir  = "${path.module}/../demo_server/demo_server/"
  output_path = "${path.module}/files/lambda_test_code.zip"
}
resource "aws_iam_role" "lambda_exec_test" {
  name = "${var.deployment_id}-lambda-execution-comp9447-test-api"

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

resource "aws_iam_role_policy_attachment" "lambda_policy_testing" {
  role       = aws_iam_role.lambda_exec_test.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "lambda_testing" {
  function_name    = "${var.deployment_id}-test-api"
  filename         = "files/lambda_test_code.zip"
  source_code_hash = data.archive_file.lambda_test_api.output_base64sha256

  runtime = "python3.9"
  handler = "app.handler"

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.lambda_exec_test.arn
}

resource "aws_cloudwatch_log_group" "lambda_testing" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_testing.function_name}"
  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "lambda_testing" {
  api_id = aws_apigatewayv2_api.gateway.id

  integration_uri    = aws_lambda_function.lambda_testing.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "lambda_testing" {
  api_id = aws_apigatewayv2_api.gateway.id

  route_key = "/test"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_testing.id}"
}

resource "aws_lambda_permission" "api_gw_testing" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_testing.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.gateway.execution_arn}/*/*"
}