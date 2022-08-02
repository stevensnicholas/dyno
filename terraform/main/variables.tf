variable "deployment_id" {
  type        = string
  description = "main-api"
}

variable "restler_lambda_timeout" {
  type        = number
  description = "restler lambda timeout - also used by SQS"
  default     = 60
}

variable "sns_fuzz_results" {
  type = string 
  description = "sns topic for issues and dynamodb messaging"
}