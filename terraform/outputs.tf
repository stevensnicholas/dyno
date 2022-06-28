output "api_endpoint" {
  value = aws_apigatewayv2_stage.lambda.invoke_url
}

output "cf_endpoint" {
  value = aws_cloudfront_distribution.frontend.domain_name
}
