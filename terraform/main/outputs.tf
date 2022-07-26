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