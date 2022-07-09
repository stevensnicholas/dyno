output "api_endpoint" {
  value = var.is_local ? null : aws_apigatewayv2_stage.lambda[0].invoke_url
}

output "api_test_endpoint" {
  value = var.is_local ? null : aws_apigatewayv2_stage.lambda_test[0].invoke_url
}

output "cf_endpoint" {
  value = var.is_local ? null : aws_cloudfront_distribution.frontend[0].domain_name
}
