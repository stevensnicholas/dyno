output "aws_lambda_function" {
  value = aws_lambda_function.lambda
}

output "static_react_bucket" {
  value = aws_s3_bucket.static_react_bucket

}

output "lambda_restler_iam_arn" {
  value = aws_iam_role.lambda_restler.arn

}