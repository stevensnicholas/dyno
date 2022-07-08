# Lambda Function to send a fuzz request to queue

resource "aws_iam_role" "iam_lambda_fuzz_request" {
  name = "${var.deployment_id}-lambda-fuzz-request-test" #TODO remove test when ready to deploy

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

resource "aws_iam_role_policy_attachment" "lambda_fuzz_request_policy" {
  role       = aws_iam_role.iam_lambda_fuzz_request.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_cloudwatch_log_group" "lambda_fuzz_request" {
  name              = "/aws/lambda/${aws_lambda_function.lambda_fuzz_request.function_name}"
  retention_in_days = 30
}

resource "aws_lambda_function" "lambda_fuzz_request" {
  function_name    = "${var.deployment_id}-server-test" #TODO remove test when ready to deploy
  filename         = "files/backend.zip"
  source_code_hash = data.archive_file.lambda.output_base64sha256

  runtime = "go1.x"
  handler = "main"

  tracing_config {
    mode = "Active"
  }

  role = aws_iam_role.lambda_exec.arn
}

resource "aws_apigatewayv2_integration" "lambda_fuzz_request" {
  api_id = aws_apigatewayv2_api.gateway.id

  integration_uri    = aws_lambda_function.lambda_fuzz_request.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "lambda_fuzz_request" {
  api_id = aws_apigatewayv2_api.gateway.id

  route_key = "$default"
  target    = "integrations/${aws_apigatewayv2_integration.lambda_fuzz_request.id}"
}

# SQS Terraform Setup 
# TODO Get the url of the SQS to allow for your lambda to write to it 
resource "aws_sqs_queue" "fuzz_queue" {
  name                      = "${var.deployment_id}-fuzz-request-test.fifo" #TODO remove test when ready to deploy
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
  fifo_queue                  = true
  content_based_deduplication = true

  tags = {
    Environment = "production"
  }
}
