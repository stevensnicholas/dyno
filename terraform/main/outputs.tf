output "aws_lambda_function" {
  value = aws_lambda_function.lambda
}

output "static_react_bucket" {
  value = aws_s3_bucket.static_react_bucket

}