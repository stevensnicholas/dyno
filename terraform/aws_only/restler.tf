resource "aws_lambda_function" "lambda_restler" {
  function_name = "${var.deployment_id}-restler-fuzzer"
  image_uri     = var.restler_image_tag
  package_type  = "Image"
  timeout       = 60

  tracing_config {
    mode = "Active"
  }

  role = var.lambda_restler_iam_arn
}

resource "aws_lambda_event_source_mapping" "openapi_bucket_to_restler" {
  event_source_arn = var.open_api_sqs_arn
  enabled          = true
  function_name    = aws_lambda_function.lambda_restler.arn
  batch_size       = 1
}