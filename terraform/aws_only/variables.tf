variable "deployment_id" {
  type        = string
  description = "main-api"
}

variable "aws_lambda_function" {
  type        = any
  description = "lambda made in main"

}

variable "static_react_bucket" {
  type        = any
  description = "lambda made in main"

}

variable "restler_image_tag" {
  type        = string
  description = "restler image tag to pull"
}

variable "lambda_restler_iam_arn" {
  type        = string
  description = "lamdda restler iam arn"

}

variable "open_api_sqs_arn" {
  type        = string
  description = "open api bucket name used for lambda"

}

variable "restler_lambda_timeout" {
  type        = number
  description = "Time out for lambda - Related to SQS Queue visibility"

}

variable "fuzz_results_bucket" {
  type        = string
  description = "fuzz output bucket"

}

variable "open_api_sqs_url" {
  type        = string
  description = "open api sqs queue url lambda"

}

variable "issues_sns_topic_arn" {
  type        = string
  description = "issues sns arn"

}

