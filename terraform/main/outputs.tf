output "aws_lambda_function" {
  value = aws_lambda_function.lambda
}

output "static_react_bucket" {
  value = aws_s3_bucket.static_react_bucket

}

output "account_id" {
  value = data.aws_caller_identity.current.account_id
}

output "lambda_restler_iam_arn" {
  value = aws_iam_role.lambda_restler.arn

}

output "open_api_sqs_arn" {
  value = aws_sqs_queue.openapi_sqs_queue.arn

}

output "fuzz_results_bucket" {
  value = aws_s3_bucket.fuzz_results_bucket.id
}

output "open_api_sqs_url" {
  value = aws_sqs_queue.openapi_sqs_queue.id

}

output "open_api_s3_name" {
  value = aws_s3_bucket.openapi_files_bucket.id

}