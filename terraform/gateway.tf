resource "aws_apigatewayv2_api" "gateway" {
  name          = "${var.deployment_id}_gateway"
  protocol_type = "HTTP"
  cors_configuration {
    allow_origins = ["https://${aws_cloudfront_distribution.frontend.domain_name}"]
    allow_methods = ["*"]
    allow_headers = ["*"]
  }
}

resource "aws_cloudwatch_log_group" "api_gw" {
  name              = "/aws/api_gw/${aws_apigatewayv2_api.gateway.name}"
  retention_in_days = 30
}

resource "aws_apigatewayv2_stage" "lambda" {
  api_id = aws_apigatewayv2_api.gateway.id

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
