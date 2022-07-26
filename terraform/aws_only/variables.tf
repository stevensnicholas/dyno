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