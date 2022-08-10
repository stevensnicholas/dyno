variable "deployment_id" {
  type        = string
  description = "main-api"
}

variable "restler_image_tag" {
  type        = string
  description = <<EOT
  restler image tag to pull, if this is the first run then you need to supply an image from another ecr
  For example to get the current latest pushed image on main (production) run the following comand
  echo 117712065617.dkr.ecr.ap-southeast-2.amazonaws.com/main_dyno_image_repository:`aws ecr describe-images --repository-name main_dyno_image_repository --query 'sort_by(imageDetails,& imagePushedAt)[-1].imageTags[0]' | tr -d '"'`
  EOT
}

variable "restler_lambda_timeout" {
  type        = number
  description = "restler lambda timeout - also used by SQS"
  default     = 300
}

variable "client_id"{
  type        = string
  description = "GitHub OAuth client ID"
}

variable "client_secret"{
  type        = string
  description = "GitHub OAuth client secret"
}